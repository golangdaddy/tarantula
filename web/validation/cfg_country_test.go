package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestCountryISO2(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"":	nil,
		"GB": success,
		"ES": success,
		"IN": success,
		"GBR": nil,
		"US": success,
		"USA": nil,
	}

	vc := CountryISO2()

	for test, result := range tests {

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED: "+test)
		}

	}

	for test, result := range tests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED: "+test)
		}

	}
}
