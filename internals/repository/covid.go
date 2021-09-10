package repository

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"lm-test/internals/config"
	"lm-test/internals/entity"
	client "lm-test/internals/infrastructure/http_client"
	"net/http"
)

type CovidRepositoryInterface interface {
	LoadResult(ctx context.Context, id int) error
}

type CovidRepository struct {
	client *client.HttpClient
	config config.Configuration
}

func (r *CovidRepository) LoadResult(result *entity.Result) error {
	req, err := http.NewRequest("GET", r.config.CovidCaseURL, nil)
	if err != nil {
		return err
	}
	res, err := r.client.Request(req)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, result); err != nil {
		return err
	}

	return nil
}

func NewCovidRepository(client *client.HttpClient, config config.Configuration) *CovidRepository {
	return &CovidRepository{
		client: client,
		config: config,
	}
}
