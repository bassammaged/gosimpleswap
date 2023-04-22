package currency

import (
	"github.com/bassammaged/gosimpleswap/platform"
)

type Currency struct {
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	Network           string `json:"network"`
	HasExtraId        bool   `json:"has_extra_id"`
	ExtraId           string `json:"extra_id"`
	Image             string `json:"image"`
	WarningsFrom      string `json:"warnings_from"`
	Warningsto        string `json:"warnings_to"`
	ValidationAddress string `json:"validation_address"`
	ValidationExtra   string `json:"validation_extra"`
	AddressExplorer   string `json:"address_explorer"`
	TxExplorer        string `json:"tx_explorer"`
	ConfirmationsFrom string `json:"confirmations_from"`
}

// Return info about currency by provided symbol
func (c *Currency) GetCurrency(currencySymbol string, user platform.User) (interface{}, error) {
	// Prepare the URL
	url := platform.URL + "/get_currency?symbol=" + currencySymbol + "&api_key=" + user.ApiKey
	// Send the request
	response, err := platform.SendGetRequest(url)
	if err != nil {
		return nil, err
	}
	statusCode, _ := platform.CheckResponseCode(response)
	object, err := platform.GetResponseBody(statusCode, &c, response)
	if err != nil {
		return nil, err
	}
	return object, nil
}
