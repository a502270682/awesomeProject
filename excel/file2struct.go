package main

import (
	"bytes"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/liangzibo/go-excel/lzbExcel"
	"github.com/mitchellh/mapstructure"
)

type UploadTravelInfo struct {
	TravelType      string `json:"TravelType" csv:"Travel Type(One option)" index:"0"`
	TravellerType   string `json:"TravellerType" csv:"Traveller Type" index:"1"`
	TravelStartDate string `json:"TravelStartDate" csv:"Travel Start Date" index:"2"`
	TravelEndDate   string `json:"TravelEndDate" csv:"Travel End Date" index:"3"`
}

func parseRowsToSlice(fileInfo []byte) ([]UploadTravelInfo, error) {
	file, err := excelize.OpenReader(bytes.NewReader(fileInfo))
	if err != nil {
		return nil, err
	}
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}
	var out []UploadTravelInfo
	err = lzbExcel.NewExcelStructDefault().SetPointerStruct(&UploadTravelInfo{}).RowsAllProcess(rows, func(maps map[string]interface{}) error {
		var ptr UploadTravelInfo
		// map 转 结构体
		if err := mapstructure.Decode(maps, &ptr); err != nil {
			return err
		}
		out = append(out, ptr)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}
