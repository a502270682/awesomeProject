package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func write() {
	f := excelize.NewFile()
	idx1, idx2 := f.NewSheet("Sheet1"), f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetActiveSheet(idx1)
	f.SetActiveSheet(idx2)
	f.SetCellValue("Sheet1", "A2", "Hello world.")
	f.SetCellValue("Sheet2", "B2", 100)
	// Set active sheet of the workbook.

	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func read() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet2", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func main() {
}

