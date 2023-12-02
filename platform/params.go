package platform

type param struct {
	key   string
	value string
}

// Create an empty params
func newParam(key, value string) param {
	return param{
		key:   key,
		value: value,
	}
}

type Params struct {
	query []param
}

// Create an empty params
func NewParams() *Params {
	return &Params{}
}

// Add inserts a new parameter into []param slice.
// Add expects two value key and value, the both are string.
//
//	params.Add("symbol","btc")
func (p *Params) Add(key, value string) {
	p.query = append(p.query, newParam(key, value))
}
