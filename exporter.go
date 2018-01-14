package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

var csharpTypeNames map[string]string
var golangTypeNames map[string]string

type Exporter struct {
	luaTableTemplate []byte
	jsonTemplate     []byte
	csharpTemplate   []byte
	golangTemplate   []byte
}

func (e *Exporter) Init() {
	e.luaTableTemplate, _ = ioutil.ReadFile("templates/lua.tmpl")
	e.jsonTemplate, _ = ioutil.ReadFile("templates/json.tmpl")
	e.csharpTemplate, _ = ioutil.ReadFile("templates/csharp.tmpl")
	e.golangTemplate, _ = ioutil.ReadFile("templates/golang.tmpl")

	initCsharpTypeNames()
	initGolangTypeNames()
}

func (e *Exporter) ExportLua(folder string, xlsx *Xlsx) {
	tmpl, err := template.New("luaExport").Funcs(sprig.HermeticTxtFuncMap()).Parse(string(e.luaTableTemplate))
	//tmpl.Funcs(sprig.FuncMap())
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create(path.Join(folder, xlsx.FileName+".lua"))
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, xlsx)
	if err != nil {
		panic(err)
	}
}

func (e *Exporter) ExportJson(folder string, xlsx *Xlsx) {
	tmpl, err := template.New("jsonExport").Funcs(sprig.HermeticTxtFuncMap()).Parse(string(e.jsonTemplate))
	//tmpl.Funcs(sprig.FuncMap())
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create(path.Join(folder, xlsx.FileName+".json"))
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, xlsx)
	if err != nil {
		panic(err)
	}
}

func (e *Exporter) ExportCSharp(folder string, xlsx *Xlsx) {
	tmpl, err := template.New("csharpExport").Funcs(sprig.HermeticTxtFuncMap()).Funcs(genericFuncMap()).Parse(string(e.csharpTemplate))
	//tmpl.Funcs(sprig.FuncMap())
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create(path.Join(folder, xlsx.Name+".cs"))
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, xlsx)
	if err != nil {
		panic(err)
	}
}

func (e *Exporter) ExportGolang(folder string, xlsx *Xlsx) {
	tmpl, err := template.New("golangExport").Funcs(sprig.HermeticTxtFuncMap()).Funcs(genericFuncMap()).Parse(string(e.golangTemplate))
	//tmpl.Funcs(sprig.FuncMap())
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create(path.Join(folder, xlsx.FileName+".go"))
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
		"GetGolangTypeName": parseGolangType,
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

func parseGolangType(longType string) string {
	if golangTypeNames[longType] != "" {
		return golangTypeNames[longType]
	}
	first := strings.Index(longType, "<")
	last := strings.LastIndex(longType, ">")
	if first != -1 && last != -1 {
		thistype := longType[:first]
		subType := longType[first+1 : last]
		return golangTypeNames[thistype] + parseGolangType(subType)
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
	golangTypeNames["int"] = "int32"
	golangTypeNames["float"] = "float32"
	golangTypeNames["string"] = "string"
	golangTypeNames["bool"] = "bool"
	golangTypeNames["dict"] = "map[string]"
	golangTypeNames["list"] = "[]"
}
