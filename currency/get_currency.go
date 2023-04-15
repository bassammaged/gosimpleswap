package currency

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const url = "https://api.simpleswap.io/get_currency"

type UserData struct {
	ApiKey string
}

type Currency struct {
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Network          string `json:"network"`
	Has_extra_id     bool   `json:"has_extra_id"`
	Address_explorer string `json:"address_explorer"`
	Extra_id         string `json:"extra_id"`
	Image            string `json:"image"`
}

func NewUserData(apiKey string) *UserData {
	var userData UserData = UserData{
		ApiKey: apiKey,
	}
	return &userData
}

func GetCurrency(currencySymbol string, userData *UserData) (*Currency, error) {
	getUrl := url + "?symbol=" + currencySymbol + "&api_key=" + userData.ApiKey
	response, err := http.Get(getUrl)
	if err != nil {
		return &Currency{}, err
	}

	if err = evaluateTheResponseCode(response); err != nil {
		return &Currency{}, err
	}

	currency, err := getResponseBody(*response)
	if err != nil {
		return &Currency{}, err
	}

	return currency, nil
}

func evaluateTheResponseCode(response *http.Response) error {
	if response.StatusCode != 200 {
		return errors.New("status code is not equal to 200")
	}
	return nil
}

func getResponseBody(response http.Response) (*Currency, error) {
	BodyInBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return &Currency{}, err
	}
	var currency Currency
	json.Unmarshal(BodyInBytes, &currency)

	return &currency, nil
}
