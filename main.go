package main

import (
	"fmt"

	"github.com/bassammaged/gosimpleswap/currency"
)

func main() {
	userData := currency.NewUserData("xxx-xxx-xxxxxx-xxxxxx-xxxxxxx")

	currency, err := currency.GetCurrency("BSV", userData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(currency)
}
