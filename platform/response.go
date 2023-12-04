package platform

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"slices"
)

// errorResponse
type definedErrorMessage struct {
	Code             int    `json:"code"`
	ErrorMessage     string `json:"error"`
	ErrorDescription string `json:"description"`
	TraceId          string `json:"trace_id"`
}

type undefinedErrorMessage struct {
	StatusCode int
	Content    string
}

// ResponseReader
func (p Platform) ResponseReader(jsonStruct any, response *http.Response) (fetechedData any, err error) {

	// Parse the response in Bytes
	responseInBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Check the response status code
	responseCode := []int{200, 400, 401, 403}

	if slices.Contains(responseCode, response.StatusCode) {
		if response.StatusCode == http.StatusOK {
			parsedResponse := jsonStruct
			if err := p.decodingExpectedJsonContent(parsedResponse, responseInBytes); err != nil {
				return nil, err
			}
			return parsedResponse, nil
		} else {
			parsedResponse := &definedErrorMessage{}
			if err := p.decodingExpectedJsonContent(parsedResponse, responseInBytes); err != nil {
				return nil, err
			}
			return parsedResponse, errors.New(parsedResponse.ErrorDescription)
		}

	} else {
		parsedResponse := &undefinedErrorMessage{StatusCode: response.StatusCode}
		p.decodingUnexpectedErrorContent(parsedResponse, responseInBytes)
		return parsedResponse, errors.New(parsedResponse.Content)
	}
}

// decodingExpectedJsonContent maps JSON content to a struct.
// It excpects struct that will be used for the mapping and responsebody in bytes.
// It returns error in case of any error.
func (p Platform) decodingExpectedJsonContent(jsonStruct interface{}, responseBody []byte) (err error) {
	if err := json.Unmarshal(responseBody, jsonStruct); err != nil {
		return err
	}

	return nil
}

// decodingUnexpectedErrorContent converts the response body of unexpected HTTP status into string.
// It excpects pointer of undefinedErrorMessage struct and responsebody in bytes.
func (p Platform) decodingUnexpectedErrorContent(errorMsg *undefinedErrorMessage, responseBody []byte) {
	errorMsg.Content = string(responseBody)
}
