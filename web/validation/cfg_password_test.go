package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestPasswordWeak(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"hello w": nil,
		"hello world": success,
		"k7!464&5": success,
	}

	vc := PasswordWeak()

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

func TestPasswordHard(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"hello world": nil,
		"57!464&5": nil,
		"b9W7!4a4&5be": success,
	}

	vc := PasswordHard()

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
