package main

import (
	"fmt"

	"github.com/bassammaged/gosimpleswap/platform"
)

func main() {
	p, err := platform.NewPlatform("testAPiKey")
	if err != nil {
		fmt.Println(err)
	}

	queryParams := platform.NewParams()
	queryParams.Add("symbol", "btc")
	queryParams.Add("apikey", "testthat")

	p.GetRequest("/get_currency", *queryParams)
}
