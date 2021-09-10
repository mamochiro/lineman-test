package service

import "lm-test/internals/entity"

type CovidService struct{}

func NewCovidService() *CovidService {
	return &CovidService{}
}

func (c *CovidService) SummarizeResult(result *entity.Result) *entity.Summary {
	province := map[string]int{
		"N/A": 0,
	}
	ageGroup := map[string]int{
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
		"N/A":   0,
	}

	for _, data := range result.Data {
		if data.Province == "" {
			data.Province = "N/A"
		}
		province[data.Province]++
		age := data.Age
		if age >= 0 && age <= 30 {
			ageGroup["0-30"]++
		} else if age >= 31 && age <= 60 {
			ageGroup["31-60"]++
		} else if age >= 61 {
			ageGroup["61+"]++
		} else {
			ageGroup["N/A"]++
		}
	}

	return &entity.Summary{
		Province: province,
		AgeGroup: ageGroup,
	}
}
