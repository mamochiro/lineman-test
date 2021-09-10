package service

import (
	"lm-test/internals/entity"
	"testing"
)

func TestSummarizeResult(t *testing.T) {
	service := CovidService{}
	dummyResult := &entity.Result{
		Data: []struct {
			ConfirmDate    string      `json:"ConfirmDate"`
			No             interface{} `json:"No"`
			Age            int         `json:"Age"`
			Gender         string      `json:"Gender"`
			GenderEn       string      `json:"GenderEn"`
			Nation         interface{} `json:"Nation"`
			NationEn       string      `json:"NationEn"`
			Province       string      `json:"Province"`
			ProvinceID     int         `json:"ProvinceId"`
			District       interface{} `json:"District"`
			ProvinceEn     string      `json:"ProvinceEn"`
			StatQuarantine int         `json:"StatQuarantine"`
		}{},
	}
	summary := service.SummarizeResult(dummyResult)

	if summary.AgeGroup["N/A"] == 0 {
		t.Error("Error")
	}
}
