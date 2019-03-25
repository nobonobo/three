package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	prefix       = os.Getenv("PWD") + "/three.js/src/"
	outputPrefix = "github.com/nobonobo/three"
	outputDir    = filepath.Clean(filepath.Join(os.Getenv("PWD"), "..", ".."))
)

var (
	fileStack  []*File
	klassStack []*Klass
)

func getDefineAndPos(file *File) *DefineAndPos {
	dir, name := filepath.Split(file.Path)
	dir = filepath.Clean(dir)
	if dir == "." {
		dir = ""
	}
	name = strings.TrimSuffix(name, ".d.ts")
	_, parent := filepath.Split(dir)
	return &DefineAndPos{
		Pos:    dir,
		Parent: parent,
		Name:   name,
	}
}

func walk(tab int, node *Define) {
	switch node.Kind {
	case 0:
		// top
	case ExternalModule:
		if !strings.HasPrefix(node.OriginalName, prefix) {
			return
		}
		srcPath := node.OriginalName[len(prefix):]
		file := &File{
			Path:    srcPath,
			Imports: map[string]struct{}{},
		}
		Files = append(Files, file)
		fileStack = append(fileStack, file)
		defer func() {
			fileStack = fileStack[:len(fileStack)-1]
		}()
	case Module:
		file := fileStack[len(fileStack)-1]
		dir, _ := path.Split(file.Path)
		file = &File{
			Path:    path.Join(dir, node.Name, strings.ToLower(node.Name)+".d.ts"),
			Imports: map[string]struct{}{},
			Module:  true,
		}
		Files = append(Files, file)
		fileStack = append(fileStack, file)
		defer func() {
			fileStack = fileStack[:len(fileStack)-1]
		}()
	case Function:
		file := fileStack[len(fileStack)-1]
		file.Funcs = append(file.Funcs, node.Signatures...)
	case Enumeration:
		if len(klassStack) == 0 {
			file := fileStack[len(fileStack)-1]
			file.Enums = append(file.Enums, &Enum{
				Name: node.Name,
			})
		} else {
			klass := klassStack[len(klassStack)-1]
			klass.Enums = append(klass.Enums, &Enum{
				Name: node.Name,
			})
		}
	case EnumerationMember:
		var enum *Enum
		if len(klassStack) == 0 {
			file := fileStack[len(fileStack)-1]
			enum = file.Enums[len(file.Enums)-1]
		} else {
			klass := klassStack[len(klassStack)-1]
			enum = klass.Enums[len(klass.Enums)-1]
		}
		enum.Members = append(enum.Members, node.Name)
	case Variable:
		file := fileStack[len(fileStack)-1]
		file.Vars = append(file.Vars, &Prop{
			Name: node.Name,
			Type: node.Type,
		})
	case Interface:
		file := fileStack[len(fileStack)-1]
		klass := &Klass{
			Name: node.Name,
		}
		file.Interfaces = append(file.Interfaces, klass)
		klassStack = append(klassStack, klass)
		defer func() {
			klassStack = klassStack[:len(klassStack)-1]
		}()
	case Class:
		file := fileStack[len(fileStack)-1]
		klass := &Klass{
			Name: node.Name,
		}
		file.Klasses = append(file.Klasses, klass)
		klassStack = append(klassStack, klass)
		defer func() {
			klassStack = klassStack[:len(klassStack)-1]
		}()
	case Constructor:
		klass := klassStack[len(klassStack)-1]
		klass.Constructor = append(klass.Constructor, node.Signatures...)
	case Method:
		klass := klassStack[len(klassStack)-1]
		klass.Methods = append(klass.Methods, node.Signatures...)
	case Property:
		klass := klassStack[len(klassStack)-1]
		klass.Properties = append(klass.Properties, &Prop{
			Name: node.Name,
			Type: node.Type,
		})
	case TypeLiteral, TypeAlias:
		// not supported
	default:
		log.Fatalln("unknown:", node.KindString, "from", node.Kind)
	}
	for _, child := range node.Children {
		walk(tab+1, child)
	}
	if len(fileStack) > 0 {
		file := fileStack[len(fileStack)-1]
		defPos := getDefineAndPos(file)
		defPos.Interface = node.Kind == Interface
		Links[node.ID] = defPos
	}
}

func main() {
	var node *Define
	fp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.NewDecoder(fp).Decode(&node); err != nil {
		if jerr, ok := err.(*json.SyntaxError); ok {
			log.Fatalf("%s in offset: %d", jerr.Error(), jerr.Offset)
		}
		log.Fatal(err)
	}
	walk(0, node)
	log.Println("export:", outputDir)
	if err := Export(outputDir); err != nil {
		log.Fatal(err)
	}
}
