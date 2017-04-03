package validation

import (
	"strings"
	//
	"github.com/golangdaddy/tarantula/web"
)

var validation_countries_map map[string]*Country

// Returns a validation object that checks to see if it can resolve to a country struct
func CountryISO2() *Config {

	min := 2.0
	max := 2.0

	if validation_countries_map == nil {

		validation_countries_map = Countries()

	}

	return NewConfig(
		"US",
		func (req web.RequestInterface, s string) (status *web.ResponseStatus, _ interface{}) {

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToUpper(s)),
			)
			if status != nil {
				return status, nil
			}

			country := validation_countries_map[s]

			if country == nil { status = req.Respond(400, "COUNTRY NOT FOUND: "+s) }

			return status, country
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			s, ok := param.(string); if !ok { return nil, nil }

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToUpper(s)),
			)
			if status != nil {
				return status, nil
			}

			country := validation_countries_map[s]

			if country == nil { status = req.Respond(400, "COUNTRY NOT FOUND: "+s) }

			return status, country
		},
		min,
		max,
	)
}
