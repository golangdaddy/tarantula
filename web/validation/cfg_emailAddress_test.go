package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestEmailAddress(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"hello_@_world.com": success,
		"hello@@world.com": nil,
		"957!464&5": nil,
		"hello@world.com": success,
		"alex@cpu.host": success,
	}

	vc := EmailAddress()

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

func TestEmailAddressOptional(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": success,
		"hello_@_world.com": success,
		"hello@@world.com": nil,
		"957!464&5": nil,
		"hello@world.com": success,
		"alex@cpu.host": success,
	}

	vc := EmailAddressOptional()

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
