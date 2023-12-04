package endpoints

import (
	"strconv"

	"github.com/bassammaged/gosimpleswap/platform"
)

type Pair struct {
	Content []string
}

func NewPair() *Pair {
	return &Pair{}
}

func (p *Pair) GetPairs(configuredPlatform *platform.Platform, fixed bool, currencySymbol string) (result any, err error) {
	// Prepare query parameters
	queryParams := platform.NewParams()
	queryParams.Add("symbol", currencySymbol)
	queryParams.Add("fixed", strconv.FormatBool(fixed))
	queryParams.Add("api_key", configuredPlatform.ApiKey)

	res, err := configuredPlatform.GetRequest("/get_pairs", *queryParams)
	if err != nil {
		return nil, err
	}

	response, err := configuredPlatform.ResponseReader(p, res)
	if err != nil {
		return nil, err
	}

	return response, nil
}
