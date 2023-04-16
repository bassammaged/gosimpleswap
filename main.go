package main

import (
	"fmt"

	"github.com/bassammaged/gosimpleswap/currency"
	"github.com/bassammaged/gosimpleswap/platform"
)

func main() {
	userData := platform.NewUserData("3xeexxxa-8xx4-4dx3-xxx0-606xxxx281xx4")

	currency := currency.Currency{}
	object, err := currency.GetCurrency("BSV", *userData)
	fmt.Println(object)
	if err != nil {
		fmt.Println(err)
	}
}
