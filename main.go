package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("excels/SimpleTypes.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	// cell := xlsx.GetCellValue("Sheet1", "B2")
	// fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")
	x := new(Xlsx)
	x.Init("SimpleTypes")
	x.Parse(rows)

	exporter := new(Exporter)
	exporter.Init()

	exporter.ExportLua(x)
	exporter.ExportJson(x)
	exporter.ExportCSharp(x)
	exporter.ExportGolang(x)

	//x.Print()
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }
	return
}
