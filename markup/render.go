package g

import 	(
		"fmt"
		"bytes"
		"errors"
		"strings"
		"reflect"
		"encoding/json"
		"html/template"
		//
		"github.com/fatih/structs"
		)

var funcMap template.FuncMap

// template helper functions

func divideFF(a, b float64) float64 { return a / b }

func percentageFF(a, b float64) float64 { return (a / b) * 100 }
func percentageIF(a int, b float64) float64 { return (float64(a) / b) * 100 }

func encodeJSON(data interface{}) template.JS {

	a, err := json.Marshal(data)

	if err != nil { panic(err) }

	return template.JS(a)
}

func float64ToInt(f float64) int { return int(f) }

func float64ToInt64(f float64) int64 { return int64(f) }

func int64ToInt(i int64) int { return int(i) }

// code for rendering ELEMENT

func (ele *ELEMENT) Render() ([]byte, error) {

	// init map of functions available to use within delimiters

	if funcMap == nil {
		funcMap = template.FuncMap{
			"divideFF":				divideFF,
			"percentageFF":			percentageFF,
			"percentageIF":			percentageIF,
			"encodeJSON":			encodeJSON,
			"float64ToInt":			float64ToInt,
			"float64ToInt64":		float64ToInt64,
			"int64ToInt":			int64ToInt,
		}
	}

	for _, child := range ele.Children {

		child.RLock()
			childData := child.data
		child.RUnlock()

		if childData == nil { child.SetData(ele.data) }

	}

	buf := bytes.NewBuffer(nil)

	// inject inline controller before element if exists

	t, err := template.New("x").Delims("<!!", "!!>").Funcs(funcMap).Parse("<script>" + ele.Ctrl + "</script>"); if err != nil { return nil, err }

	ele.RLock()
		if len(ele.Ctrl) > 0 { t.Execute(buf, ele.data) }
	ele.RUnlock()

	if ele.data == nil {

		err := ele.RenderToBuffer(buf, nil); if err != nil { return nil, err }

		return buf.Bytes(), nil

	} else {

		if ele.Repeater != nil {

			array, err := DataToArray(ele.data); if err != nil { return nil, err }

			for _, obj := range array {

				err := ele.RenderToBuffer(buf, obj); if err != nil { return nil, err }

			}

		} else {

			array, err := DataToArray(ele.data); if err != nil { return nil, err }

			err = ele.RenderToBuffer(buf, array[0]); if err != nil { return nil, err }

			return buf.Bytes(), nil

		}

	}

	return buf.Bytes(), nil
}

func ResolveToInterfaceArray(target interface{}) []interface{} {

	switch v := target.(type) {

		case map[string]interface{}:

			return []interface{}{v}


		case []interface{}:

			return v


		case *interface{}:

			return (*v).([]interface{})


		case []string:

			array := make([]interface{}, len(v))

			for i, value := range v {

				array[i] = value
			}

			return array


		case nil:

			b, _ := json.Marshal(v)
			panic("*ELEMENT.Render: NIL REPEATER OBJECT: "+string(b))

			return nil

	}

	panic("*ELEMENT.Render: "+reflect.TypeOf(target).String())
	return nil
}

func DataToArray(data interface{}) ([]map[string]interface{}, error) {

	var array []interface{}

	switch m := data.(type) {

		case []string:

			array = ResolveToInterfaceArray(data)


		case []interface{}:

			array = m


		case []map[string]interface{}:

			array = make([]interface{}, len(m))

			for i, object := range m { array[i] = object }


		case map[string]interface{}:

			return []map[string]interface{}{m}, nil

		default:

			var mm map[string]interface{}

			if data != nil {

				if structs.IsStruct(data) {

					mm = structs.Map(data)

				}

			}

			array = ResolveToInterfaceArray(mm)

	}

	output := make([]map[string]interface{}, len(array))

	for limit, v := range array {

		if v == nil { continue }

		i := map[string]interface{}{}

		switch item := v.(type) {

			case map[string]string:

				for key, value := range item { i[key] = value }

			case map[string]interface{}:

				i = item

			case string:

				i = map[string]interface{}{"value": item}

			case int:

				i = map[string]interface{}{"value": item}
				
			case float64:

				i = map[string]interface{}{"value": item}
				
			case bool:

				i = map[string]interface{}{"value": item}

			case struct{}:

				i = structs.Map(item)

			default:

				return nil, errors.New("INVALID REPATER ARRAY ITEM: "+reflect.TypeOf(v).String())

		}

		output[limit] = i

	}

	return output, nil
}

func (ele *ELEMENT) RenderToBuffer(buf *bytes.Buffer, data map[string]interface{}) error {

	ele.RLock()
		inner := [][]byte{[]byte(ele.Transclude)}
	ele.RUnlock()

	for _, child := range ele.Children {

		child.SetData(data)

		b, err := child.Render(); if err != nil { return errors.New("FAILED TO RENDER CHILD: "+child.Tag+" / "+err.Error()) }

		inner = append(inner, b)
	}

	var classes string
	if len(ele.Classes) > 0 {
		classes = fmt.Sprintf(" class=\"%v\"", strings.Join(ele.Classes, " "))
	}

	// attributes := " aria-label='SSR'"
	// removing this I will manually add to elements that cause error,
	attributes := ""

	for k, v := range ele.Attributes {
		attributes += " " + k
		if len(v) > 0 { attributes += "=\"" + v + "\"" }
	}

	innerHtml := bytes.Join(inner, nil)

	var styles string
	if len(ele.Styles) > 0 {
		var s string
		for k, v := range ele.Styles { s += k + ":" + v + ";" }
		styles = fmt.Sprintf(" style=\"%s\"", s)
	}

	var id string
	if len(ele.Namespace) > 0 { id = fmt.Sprintf(" id=\"%s\"", ele.Namespace) }

	output := fmt.Sprintf("<%s%s%s%s%s>%s</%s>", ele.Tag, id, classes, attributes, styles, innerHtml, ele.Tag)

	t, err := template.New("").Delims(DELIM_L, DELIM_R).Funcs(funcMap).Parse(output); if err != nil { return err }
	err = t.Execute(buf, data); if err != nil { return err }

	return nil
}
