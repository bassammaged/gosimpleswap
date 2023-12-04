package endpoints

import (
	"github.com/bassammaged/gosimpleswap/platform"
)

type Currency struct {
	Name              string   `json:"name"`
	Symbol            string   `json:"symbol"`
	Network           string   `json:"network"`
	ContractAddress   string   `json:"contract_address"`
	HasExtraID        bool     `json:"has_extra_id"`
	ExtraID           string   `json:"extra_id"`
	Image             string   `json:"image"`
	WarningsFrom      []string `json:"warnings_from"`
	WarningsTo        []string `json:"warnings_to"`
	ValidationAddress string   `json:"validation_address"`
	ValidationExtra   string   `json:"validation_extra"`
	AddressExplorer   string   `json:"address_explorer"`
	TxExplorer        string   `json:"tx_explorer"`
	ConfirmationsFrom string   `json:"confirmations_from"`
}

func NewCurrency() *Currency {
	return &Currency{}
}

func (c *Currency) GetCurrencyInfo(configuredPlatform *platform.Platform, currencySymbol string) (result any, err error) {
	// Prepare query parameters
	queryParams := platform.NewParams()
	queryParams.Add("symbol", currencySymbol)
	queryParams.Add("api_key", configuredPlatform.ApiKey)

	res, err := configuredPlatform.GetRequest("/get_currency", *queryParams)
	if err != nil {
		return nil, err
	}

	response, err := configuredPlatform.ResponseReader(c, res)
	if err != nil {
		return nil, err
	}

	return response, nil
}
