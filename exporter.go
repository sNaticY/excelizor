package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	lua "github.com/Shopify/go-lua"
)

var csharpTypeNames map[string]string
var golangTypeNames map[string]string

type Exporter struct {
	luaTableTemplate []byte
	jsonTemplate     []byte
	csharpTemplate   []byte
}

func (e *Exporter) Init() {
	e.luaTableTemplate, _ = ioutil.ReadFile("templates/lua.tmpl")
	e.jsonTemplate, _ = ioutil.ReadFile("templates/json.tmpl")
	e.csharpTemplate, _ = ioutil.ReadFile("templates/csharp.tmpl")

	initCsharpTypeNames()

}

func (e *Exporter) ExportLua(xlsx *Xlsx) {
	tmpl, err := template.New("test").Funcs(sprig.HermeticTxtFuncMap()).Parse(string(e.luaTableTemplate))
	//tmpl.Funcs(sprig.FuncMap())
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create("exports/test.lua")
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, xlsx)
	if err != nil {
		panic(err)
	}

	l := lua.NewState()
	lua.OpenLibraries(l)

	err = lua.DoFile(l, "exports/test.lua")
	if err != nil {
		panic(err)
	}
}

func (e *Exporter) ExportJson(xlsx *Xlsx) {
	tmpl, err := template.New("test").Funcs(sprig.HermeticTxtFuncMap()).Parse(string(e.jsonTemplate))
	//tmpl.Funcs(sprig.FuncMap())
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create("exports/test.json")
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, xlsx)
	if err != nil {
		panic(err)
	}
}

func (e Exporter) ExportCSharp(xlsx *Xlsx) {
	tmpl, err := template.New("test").Funcs(sprig.HermeticTxtFuncMap()).Funcs(genericFuncMap()).Parse(string(e.csharpTemplate))
	//tmpl.Funcs(sprig.FuncMap())
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create("exports/test.cs")
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, xlsx)
	if err != nil {
		panic(err)
	}
}

func genericFuncMap() map[string]interface{} {
	var genericMap = map[string]interface{}{
		"GetCSharpTypeName": parseCsharpType,

		"GetGolangTypeName": func(name string) string {
			return golangTypeNames[name]
		},
	}

	return genericMap
}

func parseCsharpType(longType string) string {
	if csharpTypeNames[longType] != "" {
		return csharpTypeNames[longType]
	}
	first := strings.Index(longType, "<")
	last := strings.LastIndex(longType, ">")
	if first != -1 && last != -1 {
		thistype := longType[:first]
		subType := longType[first+1 : last]
		return csharpTypeNames[thistype] + parseCsharpType(subType) + ">"
	}
	return ""
}

func initCsharpTypeNames() {
	csharpTypeNames = make(map[string]string)
	csharpTypeNames["int"] = "int"
	csharpTypeNames["float"] = "float"
	csharpTypeNames["string"] = "string"
	csharpTypeNames["bool"] = "bool"
	csharpTypeNames["dict"] = "Dictionary<string, "
	csharpTypeNames["list"] = "List<"
}

func initGolangTypeNames() {
	golangTypeNames = make(map[string]string)
	csharpTypeNames["int"] = "int32"
	csharpTypeNames["float"] = "float32"
	csharpTypeNames["string"] = "string"
	csharpTypeNames["bool"] = "bool"
}
