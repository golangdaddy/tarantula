package g

import 	(
		"fmt"
		"sync"
		)

const 	(
		DELIM_L = "<!!"
		DELIM_R = "!!>"
		)

func Delims(key string) string { return DELIM_L + " ." + key + " " + DELIM_R }
func Interpolate(expression string) string { return "{{ " + expression + " }}" }

type Renderer interface{
	Render() string
}

type ELEMENT struct {
	parent *ELEMENT
	Tag string
	Namespace string
	Classes []string
	Styles map[string]string
	Children []*ELEMENT
	Repeater *Repeater
	Ctrl string
	Attributes map[string]string
	Transclude string
	RepeaterTemplate bool
	RepeaterTemplateChild bool
	RepeaterTemplateName string
	RepeaterTemplateMode string
	data interface{}
	sync.RWMutex
}

func Ele(tagType string) *ELEMENT {
	return &ELEMENT{
		Tag: 			tagType,
		Classes:		[]string{},
		Styles:			map[string]string{},
		Children:		[]*ELEMENT{},
		Attributes:		map[string]string{},
	}
}

// add data to the element's scope

// locks the element, sets the data and then unlocks the element
func (ele *ELEMENT) SetData(data interface{}) *ELEMENT {

	ele.Lock()
		ele.data = data
	ele.Unlock()

	return ele
}

// access the element's scope data

func (ele *ELEMENT) GetData() interface{} {

	ele.RLock()
		data := ele.data
	ele.RUnlock()

	return data
}

func (ele *ELEMENT) New() *ELEMENT {

	x := *ele
	y := &x

	classes := []string{}

	for _, item := range y.Classes { classes = append(classes, item) }

	y.Classes = classes


	styles := map[string]string{}

	for k, v := range y.Styles { styles[k] = v }

	y.Styles = styles


	attributes := map[string]string{}

	for k, v := range y.Attributes { attributes[k] = v }

	y.Attributes = attributes

	y.data = nil

	return y
}

// anything added via this function will appear inside the tag

func (ele *ELEMENT) Inner(s string, args ...interface{}) *ELEMENT {

	ele.Lock()
		ele.Transclude += fmt.Sprintf(s, args...)
	ele.Unlock()

	return ele
}

func (ele *ELEMENT) Attr(args ...interface{}) *ELEMENT {

	key, ok := args[0].(string); if !ok { panic("FIRST ARGUMENT TO .Attr(...) MUST BE string!") }

	switch len(args) {

		case 1:

			ele.Lock()
				ele.Attributes[key] += ""
			ele.Unlock()

		case 2:

			ele.Lock()
				ele.Attributes[key] += fmt.Sprintf("%v", args[1])
			ele.Unlock()

	}

	return ele
}

func (ele *ELEMENT) RAttr(key, value string) *ELEMENT {

	ele.Lock()
		ele.Attributes[key] = value
	ele.Unlock()

	return ele
}


// for multiple attribute diretives that don't take values like 'select-on-click'

func (ele *ELEMENT) AttrDirs(directiveNames ...string) *ELEMENT {

	for _, arg := range directiveNames{ ele.Attr(arg) }

	return ele
}

// adds a child element inside the parent

func (ele *ELEMENT) Add(childs ...*ELEMENT) *ELEMENT {

	ele.RLock()
		children := ele.Children
	ele.RUnlock()

	if children == nil { children = []*ELEMENT{} }

	for _, child := range childs {

		child.Lock()
			child.parent = ele
		child.Unlock()

		// add the new child

		ele.Lock()
			ele.Children = append(ele.Children, child)
		ele.Unlock()
	}

	return ele
}

// repeater for iterating through arrays

type Repeater struct {
	Limit int
	Order string
	data []map[string]interface{}
}

func (ele *ELEMENT) Repeat(limit int, data interface{}) *ELEMENT {

	array, err := DataToArray(data); if err != nil { panic("FAILED TO SET DATA") }

	ele.Lock()

		ele.Repeater = &Repeater{
			Limit:			limit,
			data:			array,
		}

	ele.Unlock()

	return ele
}
