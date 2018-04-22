package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestStringExact(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"1": nil,
		"12": nil,
		"123": success,
		"1234": nil,
	}

	vc := StringExact(3)

	for test, result := range tests {

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED")
			return
		}

	}

	for test, result := range tests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED")
			return
		}

	}

}
