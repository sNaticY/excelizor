package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type parameters struct {
	excelSrc   string
	luaPath    string
	jsonPath   string
	cshapPath  string
	golangPath string
	tag        string
}

var params *parameters

type fileXlsx struct {
	file *excelize.File
	xl   *xlsx
}

var loadedFiles map[string]*fileXlsx

func init() {
	params = new(parameters)
	flag.StringVar(&params.excelSrc, "p", "", "[Required] Relative `path` of excel files folder")

	flag.StringVar(&params.tag, "tag", "", "only field with this tag or empty string will be exported")
	flag.StringVar(&params.luaPath, "lua", "", "path to place exported .lua files, export no .lua files if parameter is missing")
	flag.StringVar(&params.jsonPath, "json", "", "path to place exported .json files, export no .json files if parameter is missing")
	flag.StringVar(&params.cshapPath, "csharp", "", "path to place exported .cs class files, export no .cs files if parameter is missing")
	flag.StringVar(&params.golangPath, "golang", "", "path to place exported .go struct files, export no .go files if parameter is missing")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "test" {
		params.excelSrc = "./excels/"
		params.tag = "client"
		params.luaPath = "./exports/"
		params.jsonPath = "./exports/"
		params.cshapPath = "./exports/"
		params.golangPath = "./exports/"
	} else if params.excelSrc == "" || (params.luaPath == "" && params.jsonPath == "" && params.cshapPath == "" && params.golangPath == "") || flag.Arg(0) == "help" {
		fmt.Println("Usage: excelizor -p <path> [-lua=<luaExportPath>] [-json=<jsonExportPath>] [-csharp=<csharpExportPath>] [-golang=<golangExportPath>] [-tag=<tag>] ")
		fmt.Println("       excelizor (help|test)")
		flag.PrintDefaults()
		return
	}

	loadedFiles = make(map[string]*fileXlsx)

	err := filepath.Walk(params.excelSrc, loadFile)
	if err != nil {
		log.Fatalln(err)
	}

	for key, value := range loadedFiles {
		parseExcel(key, value)
	}

	//x.Print()
	return
}

func parseExcel(name string, excel *fileXlsx) {
	output := "\nParse and export file [ " + name + " ] "
	fmt.Print(output)

	for i := 0; i < 60-len(output); i++ {
		fmt.Print(".")
	}

	if excel.xl == nil {
		excel.xl = parseFile(name, excel.file)
		exportFile(excel.xl)
	}
	fmt.Print(" Success!")
}

func loadFile(path string, f os.FileInfo, err error) error {
	if f == nil {
		log.Fatalln(err)
	}
	if f.IsDir() {
		return nil
	}
	xl, err := excelize.OpenFile(path)
	if err != nil {
		return nil
	}

	loadedFiles[f.Name()] = &fileXlsx{xl, nil}

	return err
}

func parseFile(fileName string, file *excelize.File) *xlsx {
	var data [][]string
	sheetName := file.GetSheetName(1)

	x := new(xlsx)
	if sheetName == "Vertical" {
		data = convertToVertical(file.GetRows(sheetName))
	} else {
		data = file.GetRows(sheetName)
	}

	x.ParentClass = ""
	if strings.Contains(sheetName, "|") {
		i := strings.LastIndex(sheetName, "|")
		j := strings.LastIndex(sheetName, ".")
		x.ParentClass = sheetName[i+1 : j]
	}

	lower, camel := name2lower2Camel(fileName)
	x.Init(lower, camel)
	x.Parse(data)
	return x
}

func exportFile(x *xlsx) {
	e := new(exporter)
	e.Init()

	if params.luaPath != "" {
		e.ExportLua(params.luaPath, x)

	}
	if params.jsonPath != "" {
		e.ExportJSON(params.jsonPath, x)

	}
	if params.cshapPath != "" {
		e.ExportCSharp(params.cshapPath, x)

	}
	if params.golangPath != "" {
		e.ExportGolang(params.golangPath, x)

	}
}
