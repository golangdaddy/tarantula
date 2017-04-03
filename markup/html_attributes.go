package g

import 	(
		"fmt"
		"strings"
		"strconv"
		)

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Global attributes - http://www.w3schools.com/tags/ref_standardattributes.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) AccessKey(value string) *ELEMENT {

	if len(value) != 1 { panic("*ELEMENT.AccessKey") }

	return ele.Attr("accesskey", value)
}

func (ele *ELEMENT) Class(args ...string) *ELEMENT {

	for _, arg := range args {

		ele.Lock()
			ele.Classes = append(ele.Classes, strings.Split(arg, " ")...)
		ele.Unlock()

	}

	return ele
}

func (ele *ELEMENT) ContentEditTable(state bool) *ELEMENT { return ele.Attr("contentedittable", fmt.Sprintf("%s", state)) }

func (ele *ELEMENT) ContextMenu(value string) *ELEMENT { return ele.Attr("contextmenu", value) }

func (ele *ELEMENT) Dir(value string) *ELEMENT {

	switch value {

		case "rtl":
		case "ltr":
		case "auto":

		default:	panic("*ELEMENT.Dir")

	}

	return ele.Attr("dir", value)
}

func (ele *ELEMENT) Draggable(state bool) *ELEMENT { return ele.Attr("draggable", fmt.Sprintf("%s", state)) }

func (ele *ELEMENT) Dropzone(value string) *ELEMENT {

	switch value {

		case "copy":
		case "move":
		case "link":

		default:	panic("*ELEMENT.Dropzone")

	}

	return ele.Attr("dropzone", value)
}

func (ele *ELEMENT) Hidden() *ELEMENT { return ele.Attr("hidden") }

func (ele *ELEMENT) Id(value string) *ELEMENT { return ele.Attr("id", value) }

func (ele *ELEMENT) Lang(value string) *ELEMENT { return ele.Attr("lang", value) }

func (ele *ELEMENT) SpellCheck(state bool) *ELEMENT { return ele.Attr("spellcheck", fmt.Sprintf("%s", state)) }

func (ele *ELEMENT) Style(s string) *ELEMENT {

	p := strings.Split(s, ";")

	for _, a := range p {

		if len(a) < 3 { continue }

		z := strings.Split(a, ":")

		if len(z) != 2 { continue }

		key := strings.TrimSpace(z[0])
		value := strings.TrimSpace(z[1])

		if len(key) == 0 || len(value) == 0 { continue }

		ele.Lock()
			ele.Styles[key] = value
		ele.Unlock()

	}

	return ele
}

func (ele *ELEMENT) TabIndex(number int) *ELEMENT { return ele.Attr("tabindex", strconv.Itoa(number)) }

func (ele *ELEMENT) Title(value string) *ELEMENT { return ele.Attr("title", value) }

func (ele *ELEMENT) Translate(state bool) *ELEMENT {

	if state { return ele.Attr("translate", "yes") }

	return ele.Attr("translate", "no")
}

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Non-global attributes
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) Href(value string) *ELEMENT { return ele.Attr("href", value) }

func (ele *ELEMENT) Target(value string) *ELEMENT { return ele.Attr("target", value) }

func (ele *ELEMENT) Type(value string) *ELEMENT { return ele.Attr("type", value) }

func (ele *ELEMENT) Color(value string) *ELEMENT { return ele.Attr("color", value) }

// Moved to html_sharedAttributes.go because of conflict
// func (ele *ELEMENT) Height(value string) *ELEMENT { return ele.Attr("height", value) }

func (ele *ELEMENT) Placeholder(value string) *ELEMENT { return ele.Attr("placeholder", value) }

func (ele *ELEMENT) Value(value string) *ELEMENT { return ele.Attr("value", value) }

// Moved to html_sharedAttributes.go because of conflict
// func (ele *ELEMENT) Width(value string) *ELEMENT { return ele.Attr("width", value) }
