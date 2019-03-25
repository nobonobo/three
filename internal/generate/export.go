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
	// CoreDependency ...
	CoreDependency = map[string]int{
		"math":               0,
		"math/ColorKeywords": 1,
		"core":               2,
		"objects":            3,
		"materials":          4,
		"camears":            5,
		"animation":          6,
		"renderers":          7,
		"renderers/webgl":    8,
	}
	// ConstantFiles definition files
	ConstantFiles = map[string]struct{}{
		"constants.d.ts":                        struct{}{},
		"math/ColorKeywords/colorkeywords.d.ts": struct{}{},
	}
	// ExcludeFiles exclude files
	ExcludeFiles = map[string]struct{}{
		"extras/core/Curve.d.ts":     struct{}{}, // generics included
		"extras/core/CurvePath.d.ts": struct{}{}, // generics included
		"loaders/ObjectLoader.d.ts":  struct{}{}, // generics included
	}
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
}

// File ...
type File struct {
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

func importAppendJS(file *File) {
	file.Imports["github.com/gopherjs/gopherwasm/js"] = struct{}{}
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
		case "boolean":
			return "bool"
		case "number":
			return "int"
		case "any":
			importAppendJS(file)
			return "js.Value"
		}
	case "array":
		return getArrayType(file, t)
	case "reference":
		switch t.Name {
		case "Function":
			importAppendJS(file)
			return "js.Value"
		}
		if v, ok := Links[t.ID]; ok {
			if v.Parent == "" {
				return t.Name
			}
			if path.Dir(file.Path) == v.Pos {
				return t.Name
			}
			importAppend(file, t)
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
		return v.Interface
	}
	switch t.Type {
	default:
		return true
	case "reference":
	}
	return false
}

func getParamType(file *File, t *Type) string {
	tp := getType(file, t)
	if tp == "void" {
		return ""
	}
	switch t.Type {
	case "reference":
		if tp == "js.Value" {
			break
		}
		if isReference(t) {
			return tp
		}
		return "*" + tp
	case "reflection":
		return "js.Value"
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

func scan(file *File) {
	for _, v := range file.Vars {
		getType(file, v.Type)
	}
	for _, sig := range file.Funcs {
		getType(file, sig.Type)
		for _, param := range sig.Params {
			getType(file, param.Type)
		}
	}
	for _, intf := range file.Interfaces {
		for _, v := range intf.Properties {
			getType(file, v.Type)
		}
		for _, sig := range intf.Methods {
			getType(file, sig.Type)
			for _, param := range sig.Params {
				getType(file, param.Type)
			}
		}
	}
	for _, klass := range file.Klasses {
		for _, v := range klass.Properties {
			getType(file, v.Type)
		}
		for _, sig := range klass.Methods {
			getType(file, sig.Type)
			for _, param := range sig.Params {
				getType(file, param.Type)
			}
		}
	}
}

func render(pkg string, file *File) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	var w io.Writer
	//w = io.MultiWriter(os.Stdout, buff)
	w = buff
	fmt.Fprintf(w, "package %s\n\n", pkg)
	if len(file.Imports) > 0 {
		fmt.Fprintln(w, "import (")
		for pos := range file.Imports {
			fmt.Fprintf(w, "\t%q\n", pos)
		}
		fmt.Fprintln(w, ")")
	}
	for _, enum := range file.Enums {
		fmt.Fprintf(w, "type %s js.Value\n", enum.Name)
		if len(enum.Members) > 0 {
			fmt.Fprintf(w, "var (\n")
			for _, v := range enum.Members {
				fmt.Fprintf(w, "%s %s = get(%q)\n", strings.Title(v), enum.Name, v)
			}
			fmt.Fprintf(w, ")\n")
		}
	}
	if len(file.Vars) > 0 {
		if _, ok := ConstantFiles[file.Path]; ok {
			fmt.Fprintf(w, "var (\n")
			for _, v := range file.Vars {
				tp := getParamType(file, v.Type)
				switch tp {
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
							strings.Title(v.Name), tp, tp, v.Name,
						)
					} else {
						fmt.Fprintf(w, "%s = &%s{Value:get(%q)}\n",
							strings.Title(v.Name), tp, v.Name,
						)
					}
				}
			}
			fmt.Fprintf(w, ")\n")
		} else {
			for _, prop := range file.Vars {
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
				fmt.Fprintf(w, "func Set%s(v %s) {\n",
					strings.Title(prop.Name), paramType,
				)
				fmt.Fprintf(w, "\tset(%q, v)\n", prop.Name)
				fmt.Fprintln(w, "}")
			}

		}
	}
	for _, fn := range file.Funcs {
		primType := getType(file, fn.Type)
		paramType := getParamType(file, fn.Type)
		if paramType == "" {
			fmt.Fprintf(w, "func %s(",
				strings.Title(fn.Name),
			)
			params := []string{""}
			paramDefs := []string{}
			for _, param := range fn.Params {
				paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
					getName(param.Name), getParamType(file, param.Type),
				))
				params = append(params, getName(param.Name))
			}
			fmt.Fprint(w, strings.Join(paramDefs, ", "))
			fmt.Fprintf(w, ") {\n")
			fmt.Fprintf(w, "\t_Global.Call(%q%s)\n", fn.Name, strings.Join(params, ", "))
			fmt.Fprintln(w, "}")
			continue
		}
		fmt.Fprintf(w, "func %s(",
			strings.Title(fn.Name),
		)
		params := []string{""}
		paramDefs := []string{}
		for _, param := range fn.Params {
			paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
				getName(param.Name), getParamType(file, param.Type),
			))
			params = append(params, getName(param.Name))
		}
		fmt.Fprint(w, strings.Join(paramDefs, ", "))
		fmt.Fprintf(w, ") %s {\n", paramType)
		switch paramType {
		case "bool":
			fmt.Fprintf(w, "\treturn _Global.Call(%q%s).Bool()\n", fn.Name, strings.Join(params, ", "))
		case "int":
			fmt.Fprintf(w, "\treturn _Global.Call(%q%s).Int()\n", fn.Name, strings.Join(params, ", "))
		case "float64":
			fmt.Fprintf(w, "\treturn _Global.Call(%q%s).Float()\n", fn.Name, strings.Join(params, ", "))
		case "string":
			fmt.Fprintf(w, "\treturn _Global.Call(%q%s).String()\n", fn.Name, strings.Join(params, ", "))
		case "js.Value":
			fmt.Fprintf(w, "\treturn _Global.Call(%q%s)\n", fn.Name, strings.Join(params, ", "))
		default:
			if isReference(fn.Type) {
				fmt.Fprintf(w, "\treturn %s(_Global.Call(%q%s))\n",
					primType, fn.Name, strings.Join(params, ", "),
				)

			} else {
				fmt.Fprintf(w, "\treturn &%s{Value:_Global.Call(%q%s)}\n",
					primType, fn.Name, strings.Join(params, ", "),
				)
			}
		}
		fmt.Fprintln(w, "}")
	}
	for _, intf := range file.Interfaces {
		fmt.Fprintf(w, "type %s interface{\n", strings.Title(intf.Name))
		for _, fn := range intf.Methods {
			fmt.Fprintf(w, "%s(", strings.Title(fn.Name))
			if len(fn.Params) > 0 {
				params := []string{}
				for _, param := range fn.Params {
					params = append(params,
						fmt.Sprintf("%s %s", getName(param.Name), getParamType(file, param.Type)),
					)
				}
				fmt.Fprint(w, strings.Join(params, ", "))
			}
			fmt.Fprint(w, ")")
			if fn.Type.Name != "void" {
				fmt.Fprintf(w, " %s", getParamType(file, fn.Type))
			}
			fmt.Fprintln(w)
		}
		fmt.Fprintln(w, "}")
	}
	for _, klass := range file.Klasses {
		fmt.Fprintf(w, "type %s struct{\n", strings.Title(klass.Name))
		fmt.Fprintln(w, "\t js.Value")
		fmt.Fprintln(w, "}")

		// for Constructor
		multi := len(klass.Constructor) > 1
		for i, fn := range klass.Constructor {
			primType := getType(file, fn.Type)
			paramType := getParamType(file, fn.Type)
			suffix := ""
			if multi && i > 0 {
				suffix = fmt.Sprintf("%d", i+1)
			}
			fmt.Fprintf(w, "func %s(",
				strings.Title(strings.ReplaceAll(fn.Name+suffix, " ", "")),
			)
			params := []string{}
			paramDefs := []string{}
			for _, param := range fn.Params {
				paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
					getName(param.Name), getParamType(file, param.Type),
				))
				params = append(params, getName(param.Name))
			}
			fmt.Fprint(w, strings.Join(paramDefs, ", "))
			fmt.Fprintf(w, ") %s {\n", paramType)
			fmt.Fprintf(w, "\treturn &%s{Value:get(%q).New(%s)}\n",
				primType, klass.Name, strings.Join(params, ", "),
			)
			fmt.Fprintln(w, "}")
		}

		cn := []rune{}
		for _, r := range klass.Name {
			if unicode.IsUpper(r) {
				cn = append(cn, r)
			}
		}
		name := strings.ToLower(string(cn))
		for _, prop := range klass.Properties {
			primType := getType(file, prop.Type)
			paramType := getParamType(file, prop.Type)
			fmt.Fprintf(w, "func (%s *%s) %s() %s {\n",
				name, strings.Title(klass.Name),
				strings.Title(prop.Name), paramType,
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
			fmt.Fprintf(w, "func (%s *%s) Set%s(v %s) {\n",
				name, strings.Title(klass.Name),
				strings.Title(prop.Name), paramType,
			)
			fmt.Fprintf(w, "\t%s.Set(%q, v)\n", name, prop.Name)
			fmt.Fprintln(w, "}")
		}
		methods := map[string]int{}
		for _, fn := range klass.Methods {
			methods[fn.Name]++
			index := methods[fn.Name]
			suffix := ""
			if index > 1 {
				suffix = fmt.Sprintf("%d", index)
			}
			primType := getType(file, fn.Type)
			paramType := getParamType(file, fn.Type)
			if paramType == "" {
				fmt.Fprintf(w, "func (%s *%s) %s(",
					name, strings.Title(klass.Name),
					strings.Title(fn.Name)+suffix,
				)
				params := []string{""}
				paramDefs := []string{}
				for _, param := range fn.Params {
					paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
						getName(param.Name), getParamType(file, param.Type),
					))
					params = append(params, getName(param.Name))
				}
				fmt.Fprint(w, strings.Join(paramDefs, ", "))
				fmt.Fprintf(w, ") {\n")
				fmt.Fprintf(w, "\t%s.Call(%q%s)\n", name, fn.Name, strings.Join(params, ", "))
				fmt.Fprintln(w, "}")
				continue
			}
			fmt.Fprintf(w, "func (%s *%s) %s(",
				name, strings.Title(klass.Name),
				strings.Title(fn.Name)+suffix,
			)
			params := []string{""}
			paramDefs := []string{}
			for _, param := range fn.Params {
				paramDefs = append(paramDefs, fmt.Sprintf("%s %s",
					getName(param.Name), getParamType(file, param.Type),
				))
				params = append(params, getName(param.Name))
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
		}
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

const template = `package %s

import "github.com/gopherjs/gopherwasm/js"

var _Global = js.Global().Get("THREE")%s

func get(key string) js.Value {
	return _Global.Get(key)
}

func set(key string, v js.Value) {
	return _Global.Set(key, v)
}
`

func writeHeader(pkg, path, fpath string) error {
	sub := []string{}
	for _, p := range strings.Split(path, string(filepath.Separator))[1:] {
		sub = append(sub, fmt.Sprintf(".Get(%q)", p))
	}
	fp, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err := fmt.Fprintf(fp, template, pkg, strings.Join(sub, "")); err != nil {
		return err
	}
	return nil
}

func writeTestFile(pkg, path, fpath string) error {
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
	/*
		b, err := json.MarshalIndent(Files, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(b)
		fmt.Println()
	*/
	for _, file := range Files {
		if _, ok := ExcludeFiles[file.Path]; ok {
			continue
		}
		if len(file.Enums) == 0 && len(file.Funcs) == 0 && len(file.Interfaces) == 0 &&
			len(file.Klasses) == 0 && len(file.Vars) == 0 {
			continue
		}
		packageName := strings.ToLower(filepath.Base(filepath.Dir(file.Path)))
		fpath := filepath.Join(root, strings.TrimSuffix(file.Path, ".d.ts")+".go")
		if file.Module {
			log.Println("export module:", file.Path)
		} else {
			log.Println("export file:", file.Path)
		}
		if packageName == "." {
			packageName = "three"
		}
		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return fmt.Errorf("export failed: %s", err)
		}
		dir := filepath.Dir(file.Path)
		fullpath := filepath.Join(root, dir, "doc.go")
		if err := writeHeader(packageName, filepath.Dir(file.Path), fullpath); err != nil {
			return fmt.Errorf("write header failed: %s", err)
		}
		fullpath = filepath.Join(root, dir, packageName+"_test.go")
		if err := writeTestFile(packageName, filepath.Dir(file.Path), fullpath); err != nil {
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
