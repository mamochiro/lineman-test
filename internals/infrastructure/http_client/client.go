package http_client

import (
	"fmt"
	"io/ioutil"
	"lm-test/internals/config"
	"lm-test/internals/constants"
	"lm-test/internals/utils"
	"log"
	"net/http"
	"time"
)

type HttpClient struct {
	config config.Configuration
}

func NewHttpClient(config config.Configuration) *HttpClient {
	return &HttpClient{
		config: config,
	}
}

// NewClient ...
func (h HttpClient) NewClient() (*http.Client, error) {
	// Setup a http client
	httpClient := &http.Client{
		Timeout: time.Duration(constants.HTTPClientRequestTimeout) * time.Second,
	}
	return httpClient, nil
}

// Request ...
func (h HttpClient) Request(req *http.Request) (*http.Response, error) {
	client, err := h.NewClient()

	if err != nil {
		errMsg := "Can not create HTTP client"
		log.Printf("Error HTTPClient.Request > %s: %s\n", errMsg, err.Error())
		return nil, utils.Error(constants.INTERNAL, errMsg)
	}

	response, err := client.Do(req)
	if err != nil {
		errMsg := "C"
		log.Printf("Error HTTPClient.Request > %s: %s\n", errMsg, err.Error())
		return nil, utils.Error(constants.INTERNAL, errMsg)
	}

	log.Printf("HTTPClient.Request > Request To [%s] %s\n", req.Method, req.URL)

	if response.StatusCode != 200 && response.StatusCode != 204 {
		responseBodyString := ""

		responseBody, err := ioutil.ReadAll(response.Body)
		if err == nil {
			responseBodyString = string(responseBody)
		}

		errMsg := fmt.Sprintf("Request Failure: StatusCode=%d ResponseBody=%s", response.StatusCode, responseBodyString)

		log.Println("Error HTTPClient.Request >", errMsg)
		return nil, utils.Error(constants.INTERNAL, errMsg)
	}

	return response, nil
}
