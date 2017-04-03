package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestLanguageISO2(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"":	nil,
		"EN": success,
	}

	vc := LanguageISO2()

	for test, result := range tests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED")
		}

	}

}
