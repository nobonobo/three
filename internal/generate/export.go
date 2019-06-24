package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode"

	"go/format"
)

var (
	// Files ...
	Files []*File
	// Links ...
	Links = map[int]*DefineAndPos{}
)

// DefineAndPos ...
type DefineAndPos struct {
	//Define *Define
	Pos       string
	Parent    string
	Name      string
	Interface bool
	EnumType  bool
}

// File ...
type File struct {
	SrcPath    string
	Path       string
	Module     bool
	Imports    map[string]struct{} `json:",omitempty"`
	Enums      []*Enum             `json:",omitempty"`
	Vars       []*Prop             `json:",omitempty"`
	Funcs      []*Signature        `json:",omitempty"`
	Interfaces []*Klass            `json:",omitempty"`
	Klasses    []*Klass            `json:",omitempty"`
}

// Enum ...
type Enum struct {
	Name    string
	Members []string
}

// Prop ...
type Prop struct {
	Name string
	Type *Type
}

// Klass ...
type Klass struct {
	Name        string
	Constructor []*Signature `json:",omitempty"`
	Enums       []*Enum      `json:",omitempty"`
	Properties  []*Prop      `json:",omitempty"`
	Klasses     []*Klass     `json:",omitempty"`
	Methods     []*Signature `json:",omitempty"`
}

func resolveUnion(t *Type) *Type {
	res := t.Types[0]
	/*
		for _, t := range t.Types[1:] {
			log.Println(t)
		}
	*/
	return res
}

var nameStack []string

func nsPush(s string) {
	nameStack = append(nameStack, s)
}

func nsPop() {
	nameStack = nameStack[:len(nameStack)-1]
}

func nsPath() string {
	return strings.Join(nameStack, ".")
}

func importAppendJS(file *File) {
	file.Imports["syscall/js"] = struct{}{}
}

func importAppend(file *File, t *Type) {
	def, ok := Links[t.ID]
	dir := filepath.Clean(path.Dir(file.Path))
	if dir == "." {
		dir = ""
	}
	if ok && dir != def.Pos {
		file.Imports[path.Join(outputPrefix, def.Pos)] = struct{}{}
	}
}

func getFuncType(file *File, sig *Signature) string {
	params := []string{}
	for _, param := range sig.Params {
		params = append(params, fmt.Sprintf("%s %s", param.Name, getType(file, param.Type)))
	}
	return fmt.Sprintf("func (%s) %s", strings.Join(params, ", "), getType(file, sig.Type))
}

func getIndexType(file *File, sig *Signature) string {
	if len(sig.Params) > 0 {
		return fmt.Sprintf("map[%s]%s", getType(file, sig.Params[0].Type), getType(file, sig.Type))
	}
	return ""
}

func getReflection(file *File, t *Type) string {
	if len(t.Decolaration.IndexSignature) > 0 {
		return getIndexType(file, t.Decolaration.IndexSignature[0])
	}
	if len(t.Decolaration.Signatures) > 0 {
		return getFuncType(file, t.Decolaration.Signatures[0])
	}
	if t.Name == "" {
		importAppendJS(file)
		return "js.Value"
	}
	return ""
}

func getArrayType(file *File, t *Type) string {
	if t.ElementType != nil {
		return "[]" + getType(file, t.ElementType)
	}
	switch t.Name {
	case "Int8Array":
		return "[]int8"
	case "Int16Array":
		return "[]int16"
	case "Int32Array":
		return "[]int32"
	case "Uint8Array":
		return "[]uint8"
	case "Uint16Array":
		return "[]uint16"
	case "Uint32Array":
		return "[]uint32"
	case "Float32Array":
		return "[]float32"
	case "Float64Array":
		return "[]float64"
	}
	return "unkown"
}

func getType(file *File, t *Type) string {
	switch t.Type {
	case "intrinsic":
		switch t.Name {
		default:
			return t.Name
		case "boolean", "true", "false":
			return "bool"
		case "number":
			ns := nsPath()
			fmt.Fprintln(numlog, ns)
			if FloatList[ns] {
				return "float64"
			}
			return "int"
		case "any":
			importAppendJS(file)
			return "js.Value"
		}
	case "array":
		return getArrayType(file, t)
	case "reference":
		if _, ok := BlackList[t.Name]; ok {
			importAppendJS(file)
			return "js.Value"
		}
		switch t.Name {
		case "Function":
			importAppendJS(file)
			return "js.Value"
		}
		importAppend(file, t)
		if v, ok := Links[t.ID]; ok {
			dir := path.Dir(file.Path)
			if dir == "." {
				dir = ""
			}
			if v.Parent == "" && dir == "" {
				return t.Name
			}
			if dir == v.Pos {
				return t.Name
			}
			return fmt.Sprintf("%s.%s", v.Parent, t.Name)
		}
		return t.Name
	case "union":
		return getType(file, resolveUnion(t))
	case "reflection":
		cb := getReflection(file, t)
		if cb != "" {
			return cb
		}
	case "stringLiteral":
		return "string"
	}
	log.Printf("unknown: %d %q %s", t.ID, t.Name, t.Type)
	return ""
}

func isReference(t *Type) bool {
	if v, ok := Links[t.ID]; ok {
		return v.Interface || v.EnumType
	}
	switch t.Type {
	default:
		return true
	case "intrinsic":
		return t.Name != "this"
	case "union":
		if _, ok := BlackList[resolveUnion(t).Name]; ok {
			return true
		}
		return resolveUnion(t).Type != "reference"
	case "reference":
	}
	return false
}

func getParamType(file *File, t *Type) string {
	tp := getType(file, t)
	switch tp {
	case "void":
		return ""
	}
	if _, ok := BlackListArrayType[tp]; ok {
		importAppendJS(file)
		return "js.Value"
	}
	switch t.Type {
	case "union":
		if tp == "js.Value" {
			importAppendJS(file)
			break
		}
		switch resolveUnion(t).Type {
		case "stringLiteral", "intrinsic":
			return tp
		}
		return "*" + tp
	case "reference":
		if tp == "js.Value" {
			importAppendJS(file)
			break
		}
		if isReference(t) {
			return tp
		}
		return "*" + tp
	case "reflection":
		importAppendJS(file)
		return "js.Value"
	}
	return tp
}

func getKlassType(file *File, klass *Klass, t *Type) string {
	tp := getType(file, t)
	switch tp {
	case "this":
		return klass.Name
	}
	return tp
}

func getKlassParamType(file *File, klass *Klass, t *Type) string {
	tp := getParamType(file, t)
	switch tp {
	case "this":
		return "*" + klass.Name
	}
	return tp
}

func getName(name string) string {
	switch name {
	case "func":
		return "fn"
	case "type":
		return "typ"
	case "range":
		return "rng"
	}
	return name
}

func getCName(name string) string {
	cn := []rune{}
	for _, r := range name {
		if unicode.IsUpper(r) {
			cn = append(cn, r)
		}
	}
	if len(cn) == 1 {
		cn = append(cn, cn[0])
	}
	return strings.ToLower(string(cn))
}

func scan(file *File) {
	if len(file.Enums) > 0 {
		importAppendJS(file)
	}
	for _, v := range file.Vars {
		if strings.HasPrefix(getParamType(file, v.Type), "js.") {
			importAppendJS(file)
		}
	}
	for _, sig := range file.Funcs {
		if strings.HasPrefix(getParamType(file, sig.Type), "js.") {
			importAppendJS(file)
		}
		for _, param := range sig.Params {
			if strings.HasPrefix(getParamType(file, param.Type), "js.") {
				importAppendJS(file)
			}
		}
	}
	for _, intf := range file.Interfaces {
		for _, v := range intf.Properties {
			if strings.HasPrefix(getParamType(file, v.Type), "js.") {
				importAppendJS(file)
			}
		}
		for _, sig := range intf.Methods {
			if strings.HasPrefix(getParamType(file, sig.Type), "js.") {
				importAppendJS(file)
			}
			for _, param := range sig.Params {
				if strings.HasPrefix(getParamType(file, param.Type), "js.") {
					importAppendJS(file)
				}
			}
		}
	}
	for _, klass := range file.Klasses {
		importAppendJS(file)
		for _, v := range klass.Properties {
			if strings.HasPrefix(getParamType(file, v.Type), "js.") {
				importAppendJS(file)
			}
		}
		for _, sig := range klass.Methods {
			if strings.HasPrefix(getParamType(file, sig.Type), "js.") {
				importAppendJS(file)
			}
			for _, param := range sig.Params {
				if strings.HasPrefix(getParamType(file, param.Type), "js.") {
					importAppendJS(file)
				}
			}
		}
	}
}

func render(pkg string, file *File) ([]byte, error) {
	fname := strings.TrimSuffix(file.Path, ".d.ts")
	nsPush(fname)
	defer nsPop()
	buff := bytes.NewBuffer(nil)
	var w io.Writer
	//w = io.MultiWriter(os.Stdout, buff)
	w = buff
	fmt.Fprintf(w, "// Code generated from %q; DO NOT EDIT.\n\n", file.SrcPath)
	fmt.Fprintln(w, "// +build go1.12 js\n")
	fmt.Fprintf(w, "package %s\n\n", pkg)
	if len(file.Imports) > 0 {
		fmt.Fprintln(w, "import (")
		for pos := range file.Imports {
			fmt.Fprintf(w, "\t%q\n", pos)
		}
		fmt.Fprintln(w, ")")
	}
	for _, enum := range file.Enums {
		nsPush(enum.Name)
		fmt.Fprintf(w, "type %s js.Value\n", enum.Name)
		if len(enum.Members) > 0 {
			fmt.Fprintf(w, "var (\n")
			for _, v := range enum.Members {
				fmt.Fprintf(w, "%s %s = %s(get(%q))\n",
					strings.Title(v), enum.Name, enum.Name, v)
			}
			fmt.Fprintf(w, ")\n")
		}
		nsPop()
	}
	if len(file.Vars) > 0 {
		if _, ok := ConstantFiles[file.SrcPath]; ok {
			fmt.Fprintf(w, "var (\n")
			for _, v := range file.Vars {
				nsPush(strings.Title(v.Name))
				primType := getType(file, v.Type)
				paramType := getParamType(file, v.Type)
				switch paramType {
				case "bool":
					fmt.Fprintf(w, "%s bool = get(%q).Bool()\n",
						strings.Title(v.Name), v.Name,
					)
				case "int":
					fmt.Fprintf(w, "%s int = get(%q).Int()\n",
						strings.Title(v.Name), v.Name,
					)
				case "float64":
					fmt.Fprintf(w, "%s float64 = get(%q).Float()\n",
						strings.Title(v.Name), v.Name,
					)
				case "string":
					fmt.Fprintf(w, "%s string = get(%q).String()\n",
						strings.Title(v.Name), v.Name,
					)
				case "js.Value":
					fmt.Fprintf(w, "%s js.Value = get(%q)\n",
						strings.Title(v.Name), v.Name,
					)
				default:
					if isReference(v.Type) {
						fmt.Fprintf(w, "%s %s = %s(get(%q))\n",
							strings.Title(v.Name), paramType, paramType, v.Name,
						)
					} else {
						fmt.Fprintf(w, "%s = &%s{Value:get(%q)}\n",
							strings.Title(v.Name), primType, v.Name,
						)
					}
				}
				nsPop()
			}
			fmt.Fprintf(w, ")\n")
		} else {
			for _, prop := range file.Vars {
				nsPush(strings.Title(prop.Name))
				primType := getType(file, prop.Type)
				paramType := getParamType(file, prop.Type)
				fmt.Fprintf(w, "func %s() %s {\n",
					strings.Title(prop.Name), paramType,
				)
				switch paramType {
				case "bool":
					fmt.Fprintf(w, "\treturn get(%q).Bool()\n", prop.Name)
				case "int":
					fmt.Fprintf(w, "\treturn get(%q).Int()\n", prop.Name)
				case "float64":
					fmt.Fprintf(w, "\treturn get(%q).Float()\n", prop.Name)
				case "string":
					fmt.Fprintf(w, "\treturn get(%q).String()\n", prop.Name)
				case "js.Value":
					fmt.Fprintf(w, "\treturn get(%q)\n", prop.Name)
				default:
					if isReference(prop.Type) {
						fmt.Fprintf(w, "\treturn %s(get(%q))\n",
							primType, prop.Name,
						)
					} else {
						fmt.Fprintf(w, "\treturn &%s{Value:get(%q)}\n",
							primType, prop.Name,
						)
					}
				}
				fmt.Fprintln(w, "}")
				nsPop()
				nsPush("Set" + strings.Title(prop.Name))
				paramType = getParamType(file, prop.Type)
				fmt.Fprintf(w, "func Set%s(v %s) {\n",
					strings.Title(prop.Name), paramType,
				)
				fmt.Fprintf(w, "\tset(%q, v)\n", prop.Name)
				fmt.Fprintln(w, "}")
				nsPop()
			}
		}
	}
	for _, fn := range file.Funcs {
		nsPush(strings.Title(fn.Name))
		primType := getType(file, fn.Type)
		paramType := getParamType(file, fn.Type)
		if paramType == "" {
			fmt.Fprintf(w, "func %s(",
				strings.Title(fn.Name),
			)
			params := []string{""}
			paramDefs := []string{}
			for _, param := range fn.Params {
				nsPush(param.Name)
				paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
					getName(param.Name), getParamType(file, param.Type),
				))
				params = append(params, getName(param.Name))
				nsPop()
			}
			fmt.Fprint(w, strings.Join(paramDefs, ", "))
			fmt.Fprintf(w, ") {\n")
			fmt.Fprintf(w, "\tglobal().Call(%q%s)\n", fn.Name, strings.Join(params, ", "))
			fmt.Fprintln(w, "}")
			nsPop()
			continue
		}
		fmt.Fprintf(w, "func %s(",
			strings.Title(fn.Name),
		)
		params := []string{""}
		paramDefs := []string{}
		for _, param := range fn.Params {
			nsPush(param.Name)
			paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
				getName(param.Name), getParamType(file, param.Type),
			))
			params = append(params, getName(param.Name))
			nsPop()
		}
		fmt.Fprint(w, strings.Join(paramDefs, ", "))
		fmt.Fprintf(w, ") %s {\n", paramType)
		switch paramType {
		case "bool":
			fmt.Fprintf(w, "\treturn global().Call(%q%s).Bool()\n", fn.Name, strings.Join(params, ", "))
		case "int":
			fmt.Fprintf(w, "\treturn global().Call(%q%s).Int()\n", fn.Name, strings.Join(params, ", "))
		case "float64":
			fmt.Fprintf(w, "\treturn global().Call(%q%s).Float()\n", fn.Name, strings.Join(params, ", "))
		case "string":
			fmt.Fprintf(w, "\treturn global().Call(%q%s).String()\n", fn.Name, strings.Join(params, ", "))
		case "js.Value":
			fmt.Fprintf(w, "\treturn global().Call(%q%s)\n", fn.Name, strings.Join(params, ", "))
		default:
			if isReference(fn.Type) {
				fmt.Fprintf(w, "\treturn %s(global().Call(%q%s))\n",
					primType, fn.Name, strings.Join(params, ", "),
				)
			} else {
				fmt.Fprintf(w, "\treturn &%s{Value:global().Call(%q%s)}\n",
					primType, fn.Name, strings.Join(params, ", "),
				)
			}
		}
		fmt.Fprintln(w, "}")
		nsPop()
	}
	for _, intf := range file.Interfaces {
		nsPush(intf.Name)
		fmt.Fprintf(w, "type %s interface{\n", strings.Title(intf.Name))
		for _, fn := range intf.Methods {
			nsPush(strings.Title(fn.Name))
			//primType := getKlassType(file, intf, fn.Type)
			paramType := getKlassParamType(file, intf, fn.Type)
			fmt.Fprintf(w, "\t%s(", strings.Title(fn.Name))
			if len(fn.Params) > 0 {
				params := []string{}
				paramDefs := []string{}
				for _, param := range fn.Params {
					nsPush(param.Name)
					paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
						getName(param.Name), getKlassParamType(file, intf, param.Type),
					))
					params = append(params, getName(param.Name))
					nsPop()
				}
				fmt.Fprint(w, strings.Join(paramDefs, ", "))
			}
			fmt.Fprintf(w, ") %s\n", paramType)
			nsPop()
		}
		fmt.Fprintln(w, "}")
		nsPop()
	}
	for _, klass := range file.Klasses {
		nsPush(strings.Title(klass.Name))
		fmt.Fprintf(w, "type %s struct{\n", strings.Title(klass.Name))
		fmt.Fprintln(w, "\t js.Value")
		fmt.Fprintln(w, "}")

		// for Constructor
		multi := len(klass.Constructor) > 1
		for i, fn := range klass.Constructor {
			suffix := ""
			if multi && i > 0 {
				suffix = fmt.Sprintf("%d", i+1)
			}
			fname := strings.Title(strings.ReplaceAll(fn.Name+suffix, " ", ""))
			nsPush(fname)
			primType := getKlassType(file, klass, fn.Type)
			paramType := getKlassParamType(file, klass, fn.Type)
			fmt.Fprintf(w, "func %s(", fname)
			params := []string{}
			paramDefs := []string{}
			for _, param := range fn.Params {
				nsPush(param.Name)
				paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
					getName(param.Name), getKlassParamType(file, klass, param.Type),
				))
				params = append(params, getName(param.Name))
				nsPop()
			}
			fmt.Fprint(w, strings.Join(paramDefs, ", "))
			fmt.Fprintf(w, ") %s {\n", paramType)
			fmt.Fprintf(w, "\treturn &%s{Value:get(%q).New(%s)}\n",
				primType, klass.Name, strings.Join(params, ", "),
			)
			fmt.Fprintln(w, "}")
			nsPop()
		}

		methods := map[string]int{"Value": 1, "SetValue": 1, "Get": 1, "Set": 1}
		name := getCName(klass.Name)
		for _, prop := range klass.Properties {
			methods[strings.Title(prop.Name)]++
			index := methods[strings.Title(prop.Name)]
			getSuffix := ""
			if index > 1 {
				getSuffix = fmt.Sprintf("%d", index)
			}
			methods["Set"+strings.Title(prop.Name)]++
			index = methods["Set"+strings.Title(prop.Name)]
			setSuffix := ""
			if index > 1 {
				setSuffix = fmt.Sprintf("%d", index)
			}
			nsPush(strings.Title(prop.Name) + getSuffix)
			primType := getType(file, prop.Type)
			paramType := getParamType(file, prop.Type)
			fmt.Fprintf(w, "func (%s *%s) %s() %s {\n",
				name, strings.Title(klass.Name),
				strings.Title(prop.Name)+getSuffix, paramType,
			)
			switch paramType {
			case "bool":
				fmt.Fprintf(w, "\treturn %s.Get(%q).Bool()\n", name, prop.Name)
			case "int":
				fmt.Fprintf(w, "\treturn %s.Get(%q).Int()\n", name, prop.Name)
			case "float64":
				fmt.Fprintf(w, "\treturn %s.Get(%q).Float()\n", name, prop.Name)
			case "string":
				fmt.Fprintf(w, "\treturn %s.Get(%q).String()\n", name, prop.Name)
			case "js.Value":
				fmt.Fprintf(w, "\treturn %s.Get(%q)\n", name, prop.Name)
			default:
				if klass.Name == "LineBasicMaterial" && prop.Name == "blendSrc" {
					log.Println("REF:", prop.Type.Name, primType, paramType, Links[prop.Type.ID], prop.Type.Type, isReference(prop.Type), getType(file, prop.Type))
				}
				if isReference(prop.Type) {
					fmt.Fprintf(w, "\treturn %s(%s.Get(%q))\n",
						primType, name, prop.Name,
					)
				} else {
					fmt.Fprintf(w, "\treturn &%s{Value:%s.Get(%q)}\n",
						primType, name, prop.Name,
					)
				}
			}
			fmt.Fprintln(w, "}")
			nsPop()
			nsPush("Set" + strings.Title(prop.Name))
			paramType = getParamType(file, prop.Type)
			fmt.Fprintf(w, "func (%s *%s) Set%s(v %s) {\n",
				name, strings.Title(klass.Name),
				strings.Title(prop.Name)+setSuffix, paramType,
			)
			fmt.Fprintf(w, "\t%s.Set(%q, v)\n", name, prop.Name)
			fmt.Fprintln(w, "}")
			nsPop()
		}
		for _, fn := range klass.Methods {
			methods[strings.Title(fn.Name)]++
			index := methods[strings.Title(fn.Name)]
			suffix := ""
			if index > 1 {
				suffix = fmt.Sprintf("%d", index)
			}
			nsPush(strings.Title(fn.Name) + suffix)
			primType := getKlassType(file, klass, fn.Type)
			paramType := getKlassParamType(file, klass, fn.Type)
			if paramType == "" {
				fmt.Fprintf(w, "func (%s *%s) %s(",
					name, strings.Title(klass.Name),
					strings.Title(fn.Name)+suffix,
				)
				params := []string{""}
				paramDefs := []string{}
				for _, param := range fn.Params {
					nsPush(param.Name)
					paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
						getName(param.Name), getKlassParamType(file, klass, param.Type),
					))
					params = append(params, getName(param.Name))
					nsPop()
				}
				fmt.Fprint(w, strings.Join(paramDefs, ", "))
				fmt.Fprintf(w, ") {\n")
				fmt.Fprintf(w, "\t%s.Call(%q%s)\n", name, fn.Name, strings.Join(params, ", "))
				fmt.Fprintln(w, "}")
				nsPop()
				continue
			}
			fmt.Fprintf(w, "func (%s *%s) %s(",
				name, strings.Title(klass.Name),
				strings.Title(fn.Name)+suffix,
			)
			params := []string{""}
			paramDefs := []string{}
			for _, param := range fn.Params {
				nsPush(param.Name)
				paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
					getName(param.Name), getKlassParamType(file, klass, param.Type),
				))
				params = append(params, getName(param.Name))
				nsPop()
			}
			fmt.Fprint(w, strings.Join(paramDefs, ", "))
			fmt.Fprintf(w, ") %s {\n", paramType)
			switch paramType {
			case "bool":
				fmt.Fprintf(w, "\treturn %s.Call(%q%s).Bool()\n", name, fn.Name, strings.Join(params, ", "))
			case "int":
				fmt.Fprintf(w, "\treturn %s.Call(%q%s).Int()\n", name, fn.Name, strings.Join(params, ", "))
			case "float64":
				fmt.Fprintf(w, "\treturn %s.Call(%q%s).Float()\n", name, fn.Name, strings.Join(params, ", "))
			case "string":
				fmt.Fprintf(w, "\treturn %s.Call(%q%s).String()\n", name, fn.Name, strings.Join(params, ", "))
			case "js.Value":
				fmt.Fprintf(w, "\treturn %s.Call(%q%s)\n", name, fn.Name, strings.Join(params, ", "))
			default:
				if isReference(fn.Type) {
					fmt.Fprintf(w, "\treturn %s(%s.Call(%q%s))\n",
						primType, name, fn.Name, strings.Join(params, ", "),
					)

				} else {
					fmt.Fprintf(w, "\treturn &%s{Value:%s.Call(%q%s)}\n",
						primType, name, fn.Name, strings.Join(params, ", "),
					)
				}
			}
			fmt.Fprintln(w, "}")
			nsPop()
		}
		nsPop()
	}
	return buff.Bytes(), nil
}

func output(fpath string, content []byte) error {
	fp, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer fp.Close()
	out, err := format.Source(content)
	if err != nil {
		return err
	}
	if _, err := fp.Write(out); err != nil {
		return err
	}
	return nil
}

const template = `// Code generated from %q; DO NOT EDIT.

package %s

import (
	"syscall/js"
)

func global() js.Value {
	return js.Global().Get("THREE")
}

func get(key string) js.Value {
	return global().Get(key)
}

func set(key string, v interface{}) {
	global().Set(key, v)
}
`

func writeHeader(pkg, root string, file *File) error {
	dir := filepath.Dir(file.Path)
	fpath := filepath.Join(root, dir, "doc.go")
	sub := []string{}
	for _, p := range strings.Split(dir, string(filepath.Separator)) {
		if p == "." {
			continue
		}
		sub = append(sub, fmt.Sprintf(".Get(%q)", p))
	}
	fp, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err := fmt.Fprintf(fp, template,
		file.SrcPath, pkg); err != nil {
		return err
	}
	return nil
}

func writeTestFile(pkg, root string, file *File) error {
	dir := filepath.Dir(file.Path)
	fpath := filepath.Join(root, dir, pkg+"_test.go")
	if _, err := os.Stat(fpath); os.IsExist(err) {
		return nil
	}
	fp, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer fp.Close()
	fmt.Fprintf(fp, "package %s\n", pkg)
	return nil
}

// Export ...
func Export(root string) error {
	for _, file := range Files {
		if _, ok := ExcludeFiles[file.SrcPath]; ok {
			continue
		}
		if len(file.Enums) == 0 && len(file.Funcs) == 0 && len(file.Interfaces) == 0 &&
			len(file.Klasses) == 0 && len(file.Vars) == 0 {
			continue
		}
		packageName := filepath.Base(filepath.Dir(file.Path))
		fname := strings.TrimSuffix(file.Path, ".d.ts") + ".go"
		fpath := filepath.Join(root, fname)
		if file.Module {
			log.Println("export module:", fname, "from", file.SrcPath)
		} else {
			log.Println("export file:", fname, "from", file.SrcPath)
		}
		if packageName == "." {
			packageName = path.Base(outputPrefix)
		}
		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return fmt.Errorf("export failed: %s", err)
		}
		if err := writeHeader(packageName, root, file); err != nil {
			return fmt.Errorf("write header failed: %s", err)
		}
		if err := writeTestFile(packageName, root, file); err != nil {
			return fmt.Errorf("write test file failed: %s", err)
		}
		scan(file)
		b, err := render(packageName, file)
		if err != nil {
			return fmt.Errorf("render failed: %s", err)
		}
		if err := output(fpath, b); err != nil {
			return fmt.Errorf("output failed: %s", err)
		}
	}
	return nil
}
