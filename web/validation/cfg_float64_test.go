package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestFloat64(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	pathTests := map[string]*bool{
		"1780877.9797": success,
		"hello world": nil,
		"67957!464&5": nil,
	}

	bodyTests := map[interface{}]*bool{
		1780877.9797: success,
		"hello world": nil,
		67957: nil,
	}

	vc := Float64()

	for test, result := range pathTests {

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", test)
		}

	}

	for test, result := range bodyTests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", test)
		}

	}

}
