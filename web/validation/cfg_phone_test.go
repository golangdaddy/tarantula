package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestPhoneNumber(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": success,
		"07463 661890": success,
		"07 776 678765": success,
		"01843 888888": success,
		"01843 h88888": nil,
		"01843 88888": success,
		"09876 54321": success,
	}

	vc := PhoneNumber("GB")

	for test, result := range tests {

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED: "+test)
			return
		}

	}

	for test, result := range tests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED: "+test)
			return
		}

	}
}
