package platform

import (
	"fmt"
	"net/url"
)

// GetRequest sends get reguest to Simpleswap and fetch the return response.
// The function expects the URL path to construct the Simpleswap URL.
//
//	platform.GetRequest("/get_currency") // Example
func (p Platform) GetRequest(urlPath string, queryParameters Params) {

	link := p.urlConstructor(urlPath, queryParameters)
	// response, err := http.Get(link.String())
	fmt.Println(link.String())
}

func (p Platform) urlConstructor(urlPath string, queryParameters Params) url.URL {

	// Cosntruct the query parameters
	values := p.queryConstructor(queryParameters)

	// Construct the URL
	constractedURL := url.URL{
		Scheme:     "https",
		Host:       URL,
		Path:       urlPath,
		ForceQuery: true,
		RawQuery:   values,
	}
	return constractedURL
}

// queryConstructor appends the query parameter to url.URL struct.
// It exepcts Params struct.
func (p Platform) queryConstructor(queryParameters Params) string {
	v := url.Values{}
	for _, param := range queryParameters.query {
		v.Add(param.key, param.value)
	}
	return v.Encode()
}
