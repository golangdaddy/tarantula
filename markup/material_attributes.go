package g

import 	(
		"strconv"
		)

func (ele *ELEMENT) Align(align string) *ELEMENT { return ele.Attr("layout-align", align) }

func (ele *ELEMENT) Flex(value int) *ELEMENT {

	if value < 0 {

		ele.Attr("flex")

	} else {

		ele.Attr("flex", strconv.Itoa(value))

	}

	return ele
}

func (ele *ELEMENT) Layout(layout string) *ELEMENT { return ele.Attr("layout", layout) }

func (ele *ELEMENT) LayoutFill() *ELEMENT { return ele.Attr("layout-fill") }

func (ele *ELEMENT) Wrap() *ELEMENT { return ele.Attr("layout-wrap") }

func (ele *ELEMENT) Size(exp string) *ELEMENT { return ele.Attr("size", exp) }

func (ele *ELEMENT) MdBorderBottom() *ELEMENT { return ele.Attr("md-border-bottom") }

func (ele *ELEMENT) MdSelected(exp string) *ELEMENT { return ele.Attr("md-selected", exp) }

func (ele *ELEMENT) MdStretchTabs(exp string) *ELEMENT { return ele.Attr("md-stretch-tabs", exp) }

func (ele *ELEMENT) MdPlaceholder(exp string) *ELEMENT { return ele.Attr("md-placeholder", exp) }

func (ele *ELEMENT) MdScrollShrink(exp string) *ELEMENT { return ele.Attr("md-scroll-shrink", exp) }
