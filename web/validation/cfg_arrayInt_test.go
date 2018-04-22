package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestArrayInt(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	bodyValues := []interface{}{
		[]interface{}{"a", "b", "c"},
		[]interface{}{true, false, true},
		[]interface{}{10.3, 2.03, 3.3},
		[]interface{}{3, true, 3.3},
		[]interface{}{3, true, 3},
		[]interface{}{3, 2, 1},
		[]interface{}{3},
	}

	bodyResults := []*bool{
		nil,
		nil,
		nil,
		nil,
		nil,
		success,
		success,
	}

	vc := ArrayInt()

	for i, result := range bodyResults {

		test := bodyValues[i].([]interface{})

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", test)
		}

	}

}
