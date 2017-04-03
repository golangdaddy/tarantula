package httpclient

import (
	"fmt"
	"testing"
)

func TestMethods(t *testing.T) {

	client := NewClient()

	t.Run(
		"test GET",
		func (t *testing.T) {

			b, err := client.Get("https://www.google.co.uk", nil)
			if err != nil {
				t.Error(err)
				return
			}

			fmt.Println(string(b))

		},
	)

	t.Run(
		"test GET fail",
		func (t *testing.T) {

			dest := map[string]interface{}{}
			_, err := client.Get("https://www.google.co.uk", &dest)
			if err == nil {
				t.Error(fmt.Errorf("FAILED TO FAIL!"))
				return
			}

			fmt.Println(err)

		},
	)

}
