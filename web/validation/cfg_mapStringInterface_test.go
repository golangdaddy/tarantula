package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestMapStringInterface(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	bodyValues := []interface{}{
		map[string]interface{}{"hello": "world"},
		"c",
		3,
		[]interface{}{10.3, 2.03, 3.3},
	}

	bodyResults := []*bool{
		success,
		nil,
		nil,
		nil,
	}

	vc := MapStringInterface()

	for i, result := range bodyResults {

		if status, _ := vc.BodyFunction(req, bodyValues[i]); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", bodyValues[i])
		}

	}

}
