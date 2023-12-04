package platform

import (
	"net/http"
	"net/url"
)

// GetRequest sends get reguest to Simpleswap and fetch the return response.
// The function expects the URL path to construct the Simpleswap URL.
//
//	platform.GetRequest("/get_currency") // Example
func (p Platform) GetRequest(urlPath string, queryParameters Params) (*http.Response, error) {

	// Construct the query
	link := p.getUrlConstructor(urlPath, queryParameters)

	// Send http Get request
	response, err := http.Get(link.String())

	// Evaluate the errors
	if err != nil {
		return nil, err
	}
	// Returns the HTTP response in a successful attempt.
	return response, nil
}

// getUrlConstructor builds the Simpleswap URL up.
// It expects urlPath as string and query parameters as Params struct.
// It returns Simpleswap url in url.URL struct
func (p Platform) getUrlConstructor(urlPath string, queryParameters Params) url.URL {

	// Cosntruct the query parameters
	values := p.getQueryConstructor(queryParameters)

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

// getQueryConstructor appends the query parameter to url.URL struct.
// It exepcts Params struct.
func (p Platform) getQueryConstructor(queryParameters Params) string {
	v := url.Values{}
	for _, param := range queryParameters.query {
		v.Add(param.key, param.value)
	}
	return v.Encode()
}
