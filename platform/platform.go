package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// URL of simpleswap.io
const URL = "https://api.simpleswap.io"

// Object to keep the 40x responses
type FourHundredStatusCode struct {
	StatusCode  int16  `json:"code"`
	Error       string `json:"error"`
	Description string `json:"description"`
	TraceId     string `json:"trace_id"`
}

func SendGetRequest(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return &http.Response{}, err
	}
	return response, nil
}

// Check the response status code
func CheckResponseCode(response *http.Response) (int, error) {
	// 40x status code
	if response.StatusCode >= 400 && response.StatusCode < 500 {
		return response.StatusCode, fmt.Errorf("failed to communicate, ResponseStatusCode: %v", response.StatusCode)
	} else if response.StatusCode != 200 {
		return response.StatusCode, fmt.Errorf("failed to communicate, ResponseStatusCode: %v", response.StatusCode)
	}
	return response.StatusCode, nil
}

func GetResponseBody(statusCode int, expectedObject interface{}, response *http.Response) (interface{}, error) {
	var resultObject interface{}
	var err error

	if statusCode == 200 {
		resultObject, err = extractResponseBody(&expectedObject, response)
	} else if statusCode >= 400 && statusCode < 500 {
		var object FourHundredStatusCode
		resultObject, err = extractResponseBody(&object, response)
	} else {
		err = fmt.Errorf("failed to communicate, ResponseStatusCode: %v", response.StatusCode)
	}

	if err != nil {
		return nil, err
	}
	return resultObject, nil
}

func extractResponseBody(object interface{}, response *http.Response) (interface{}, error) {
	BodyInBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return object, err
	}
	json.Unmarshal(BodyInBytes, &object)
	return object, nil
}
