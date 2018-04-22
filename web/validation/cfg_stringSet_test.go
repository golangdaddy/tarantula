package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestStringSet(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"a": nil,
		"hello": success,
		"world": success,
		"hello world": nil,
	}

	vc := StringSet("hello", "world")

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
