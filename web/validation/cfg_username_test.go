package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestUsername(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"a": nil,
		"ab": nil,
		"abc": success,
		"dave": success,
		"dave_dave": success,
		"hello world": nil,
		"@dave": nil,
		"@daveydaveydaveydaveydaveydave": nil,
		"daveydaveydaveydaveydaveydave": nil,
	}

	vc := Username(3, 16)

	for test, result := range tests {

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", test)
			return
		}

	}

	for test, result := range tests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", test)
			return
		}

	}
}
