package validation

import (
	"testing"
	//
	"github.com/golangdaddy/tarantula/web"
)

func TestString(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"a": success,
		"hello world": nil,
		"957!464&5": success,
		"hello world iweufghoqiuweh oqiwhe fpiqhw fihqwe": nil,
		"hello world iweufghoqiuweh oqiwhe fpiqhw fihqwehello world iweufghoqiuweh oqiwhe fpiqhw fihqwehello world iweufghoqiuweh oqiwhe fpiqhw fihqwehello world iweufghoqiuweh oqiwhe fpiqhw fihqwe": nil,
	}

	vc := String(1, 9)

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

func TestStringSet(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"happy": nil,
		"raining": nil,
		"fish": nil,
		"yes": success,
		"no": success,
		"maybe": success,
	}

	vc := StringSet("yes", "no", "maybe")

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

func TestStringSplit(t *testing.T) {

	req := web.NewTestInterface("GET", "/")

	tests := map[string][]string{
		" yes , no , maybe ": []string{"yes", "no", "maybe"},
		"raining": []string{"raining"},
	}

	vc := StringSplit(",")

	for test, result := range tests {

		status, output := vc.PathFunction(req, test); if status != nil {
			t.Error("FAILED")
			return
		}

		if len(result) != len(output.([]string)) {
			t.Error("FAILED")
			return
		}

		for x, v := range output.([]string) {

			if result[x] != v {
				t.Error("FAILED")
				return
			}

		}
	}

	for test, result := range tests {

		status, output := vc.BodyFunction(req, test); if status != nil {
			t.Error("FAILED")
			return
		}

		if len(result) != len(output.([]string)) {
			t.Error("FAILED")
			return
		}

		for x, v := range output.([]string) {

			if result[x] != v {
				t.Error("FAILED")
				return
			}

		}
	}
}
