# GO Simpleswap



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/bassammaged/gosimpleswap">
    <img src="assets/images/SimpleSwap.svg" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Go Simpleswap.io</h3>

  <p align="center">
    Unofficial GO package for Simpleswap.io
    <br />
    <a href="https://github.com/othneildrew/gosimpleswap/docs"><strong>Explore the docs »</strong></a>
    ·
    <a href="https://github.com/bassammaged/gosimpleswap/issues">Report Bug</a>
    ·
    <a href="https://github.com/bassammaged/gosimpleswap/issues">Request Feature</a>
  </p>

  # Features
  
  - get_currency supports now.

  # Example

  ```
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
  ```

</div>