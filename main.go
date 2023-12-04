package main

import (
	"fmt"
	"log"

	"github.com/bassammaged/gosimpleswap/endpoints"
	"github.com/bassammaged/gosimpleswap/platform"
)

func main() {
	p, err := platform.NewPlatform("xxxx-xxx-xxxx-xxxx-xxxx")

	if err != nil {
		fmt.Println(err)
	}

	currency := endpoints.NewCurrency()
	result, err := currency.GetCurrencyInfo(p, "btc")
	if err != nil {
		log.Fatal(err, result)
	}

	fmt.Println(result)

}
