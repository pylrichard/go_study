package main

import (
	"fmt"
	
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func writeExcel() {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet2")
	//向单元格中写入值
	f.SetCellValue("Sheet2", "A2", "go.cn")
	f.SetCellValue("Sheet1", "B2", 100)
	//设置文件打开后显示的Sheet，0表示Sheet1
	f.SetActiveSheet(index)
	if err := f.SaveAs("1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func readExcel() {
	f, err := excelize.OpenFile("1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取单元格内容
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	//获取Sheet1中所有的行
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func addChart() {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	//添加图表
	if err := f.AddChart("Sheet1", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet1!$A$2",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$2:$D$2"
        },
        {
            "name": "Sheet1!$A$3",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$3:$D$3"
        },
        {
            "name": "Sheet1!$A$4",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SaveAs("2.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func addPicture() {
	f, err := excelize.OpenFile("2.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := f.AddPicture("Sheet1", "A2", "logo.png", ""); err != nil {
		fmt.Println(err)
	}
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	writeExcel()
	readExcel()
	addChart()
}