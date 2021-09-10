package controller

import (
	"github.com/gin-gonic/gin"
	"lm-test/internals/entity"
	"lm-test/internals/repository"
	"lm-test/internals/service"
	"net/http"
)

type CovidController struct {
	covidService    *service.CovidService
	covidRepository *repository.CovidRepository
}

func NewCovidController(covidService *service.CovidService, covidRepository *repository.CovidRepository) *CovidController {
	return &CovidController{
		covidService:    covidService,
		covidRepository: covidRepository,
	}
}

func (co *CovidController) Summary(c *gin.Context) {
	result := &entity.Result{}
	if err := co.covidRepository.LoadResult(result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	summary := co.covidService.SummarizeResult(result)
	c.JSON(http.StatusOK, summary)
}
