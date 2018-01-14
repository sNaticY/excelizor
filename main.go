package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
	lua "github.com/Shopify/go-lua"
)

type parameters struct {
	excelSrc   string
	luaPath    string
	jsonPath   string
	cshapPath  string
	golangPath string
}

var params *parameters

type fileXlsx struct {
	file *excelize.File
	xlsx *Xlsx
}

var loadedFiles map[string]*fileXlsx

func init() {
	params = new(parameters)
	flag.StringVar(&params.excelSrc, "p", "", "[Required] Relative `path` of excel files folder")

	flag.StringVar(&params.luaPath, "lua", "", "path to place exported .lua files, export no .lua files if parameter is missing")
	flag.StringVar(&params.jsonPath, "json", "", "path to place exported .json files, export no .json files if parameter is missing")
	flag.StringVar(&params.cshapPath, "csharp", "", "path to place exported .cs class files, export no .cs files if parameter is missing")
	flag.StringVar(&params.golangPath, "golang", "", "path to place exported .go struct files, export no .go files if parameter is missing")
}

func main() {
	flag.Parse()
	if params.excelSrc == "" || (params.luaPath == "" && params.jsonPath == "" && params.cshapPath == "" && params.golangPath == "") || flag.Arg(0) == "-h" || flag.Arg(0) == "--help" {
		fmt.Println("\nUsage: excelizor -p <path> [-lua=<luaExportPath>] [-json=<luaExportPath>] [-csharp=<luaExportPath>] [-golang=<luaExportPath>]")
		flag.PrintDefaults()
	}

	loadedFiles = make(map[string]*fileXlsx)

	err := filepath.Walk(params.excelSrc, loadFile)
	if err != nil {
		log.Fatalln(err)
	}

	for key, value := range loadedFiles {
		fmt.Print("Parse and export file", key, ". . . ")
		value.xlsx = parseFile(key, value.file)
		exportFile(value.xlsx)
		fmt.Print("Success!\n")
	}

	//x.Print()
	return
}

func loadFile(path string, f os.FileInfo, err error) error {
	if f == nil {
		log.Fatalln(err)
	}
	if f.IsDir() {
		return nil
	}
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		return nil
	}

	loadedFiles[f.Name()] = &fileXlsx{xlsx, nil}

	return err
}

func parseFile(fileName string, file *excelize.File) *Xlsx {
	rows := file.GetRows(file.GetSheetName(1))
	x := new(Xlsx)

	lower, camel := name2lower2Camel(fileName)
	x.Init(lower, camel)
	x.Parse(rows)
	return x
}

func exportFile(x *Xlsx) {
	exporter := new(Exporter)
	exporter.Init()

	if params.luaPath != "" {
		exporter.ExportLua(params.luaPath, x)

	}
	if params.jsonPath != "" {
		exporter.ExportJson(params.jsonPath, x)

	}
	if params.cshapPath != "" {
		exporter.ExportCSharp(params.cshapPath, x)

	}
	if params.golangPath != "" {
		exporter.ExportGolang(params.golangPath, x)

	}
}

func testJson() {
	// jsonBytes, _ := ioutil.ReadFile("exports/test.json")

	// var obj []*exports.SimpleTypes

	// err := json.Unmarshal(jsonBytes, &obj)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", obj)
	// for i, v := range obj {
	// 	fmt.Printf("%v = %v\n", i, v)
	// }
}

func testLua() {
	l := lua.NewState()
	lua.OpenLibraries(l)

	err := lua.DoFile(l, "exports/test.lua")
	if err != nil {
		panic(err)
	}
}
