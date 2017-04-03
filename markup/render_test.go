package g

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Address string
}

func TestRender(t *testing.T) {
	
	person := &Person{
		Name: 			"Joe Bloggs",
		Address:		"12 Bloggs Ave",
	}

	person1 := map[string]interface{}{
		"Name": 		"Joe Bloggs iuqdgowjeg d",
		"Address":		"12qwjdhil qwduhq w Bloggs Ave",
	}

	person2 := map[string]interface{}{
		"Name": 		"Jasdjkahsdboe Bloggs iuqdgowjeg d",
		"Address":		"12qwjdhil qwwkadhjqduhq w Bloggs Ave",
	}

	array := []interface{}{
		person,
		person1,
		person2,
	}

		ele := DIV().Add(
			DIV().Add(
			DIV().Add(
				SPAN(Delims("Name")),
				SPAN(Delims("Address")),
			),
			),
		)

	for _, v := range array {
		
		e := ele.New().SetData(v)

		b, err := e.Render(); if err != nil { t.Error(err) }

		fmt.Println(string(b))

	}


}