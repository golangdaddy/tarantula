package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestArrayInterface(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	bodyValues := []interface{}{
		"c",
		3,
		map[string]interface{}{"hello": "world"},
		[]interface{}{10.3, 2.03, 3.3},
	}

	bodyResults := []*bool{
		nil,
		nil,
		nil,
		success,
	}

	vc := ArrayInterface()

	for i, result := range bodyResults {

		if status, _ := vc.BodyFunction(req, bodyValues[i]); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", bodyValues[i])
		}

	}

}
