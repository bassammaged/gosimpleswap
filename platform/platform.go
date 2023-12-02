package platform

// URL of [simpleswap.io]
//
// [simpleswap.io]: https://simpleswap.io
const URL string = "api.simpleswap.io"

type Platform struct {
	apiKey string
}

// NewPlatform creates a new platform struct with assigned API key.
// Apply security code principle by passing the API key as enviornment variable.
// It returns platform pointer.
func NewPlatform(apiKey string) (*Platform, error) {
	// Create a new struct
	p := &Platform{
		apiKey: apiKey,
	}

	// Test the api Key

	return p, nil
}
