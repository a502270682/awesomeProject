package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strings"
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
	f, err := excelize.OpenFile("/Users/guofeiyang/Downloads/test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	rows, err := f.GetRows("Sheet1")
	for idx := 2; idx <= len(rows); idx++ {
		cellIdx := fmt.Sprintf("C%d", idx)
		cell, err := f.GetCellValue("Sheet1", cellIdx)
		if err != nil {
			fmt.Println(err)
			return
		}
		newOne := strings.Trim(cell, "\"")
		err = f.SetCellValue("Sheet1", cellIdx, newOne)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if err := f.SaveAs("/Users/guofeiyang/Downloads/test.xlsx"); err != nil {
		fmt.Println(err)
	}

}

type Test struct {
}

func main() {
	var res []*Test
	if res == nil {
		fmt.Println(1)
	}
	//
	//policies := make([]int, 210001)
	//i := 0
	//for i = 0; i < len(policies); i += 10000 {
	//	var policyRange []int
	//	if i+10000 >= len(policies) {
	//		policyRange = policies[i:]
	//	} else {
	//		policyRange = policies[i : i+10000]
	//	}
	//	fmt.Println(len(policyRange))
	//	//fmt.Println(i)
	//}
	//fmt.Println(i)
}
