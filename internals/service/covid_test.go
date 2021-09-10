package service

import (
	"gopkg.in/go-playground/assert.v1"
	"lm-test/internals/entity"
	"testing"
)

func TestSummarizeResult(t *testing.T) {
	service := NewCovidService()
	t.Run("Test From Dummy Result", func(t *testing.T) {
		dummyResult := &entity.Result{
			Data: []entity.Data{
				entity.Data{
					Age:        51,
					Gender:     "หญิง",
					Province:   "Phrae",
					ProvinceID: 46,
				},
				entity.Data{
					Age:        18,
					Gender:     "ชาย",
					Province:   "Chumphon",
					ProvinceID: 12,
				},
				entity.Data{
					Age:        70,
					Gender:     "ชาย",
					Province:   "Bangkok",
					ProvinceID: 46,
				},
				entity.Data{
					Age:        7,
					Gender:     "ชาย",
					Province:   "Bangkok",
					ProvinceID: 32,
				},
			},
		}
		summary := service.SummarizeResult(dummyResult)

		assert.Equal(t, summary.Province["Bangkok"], 2)
		assert.Equal(t, summary.Province["Phrae"], 1)
		assert.Equal(t, summary.Province["Chumphon"], 1)
		assert.Equal(t, summary.Province["N/A"], 0)

		assert.Equal(t, summary.AgeGroup["N/A"], 0)
		assert.Equal(t, summary.AgeGroup["0-30"], 2)
		assert.Equal(t, summary.AgeGroup["31-60"], 1)
		assert.Equal(t, summary.AgeGroup["61+"], 1)
	})

	t.Run("Test Empty Result", func(t *testing.T) {
		dummyEmptyResult := &entity.Result{}
		summary := service.SummarizeResult(dummyEmptyResult)

		assert.Equal(t, summary.Province["N/A"], 0)

		assert.Equal(t, summary.AgeGroup["N/A"], 0)
	})
}
