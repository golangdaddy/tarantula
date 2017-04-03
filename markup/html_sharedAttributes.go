package g

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  HTML Shared attributes
//  Some attributes are the same across elements, and are included here
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) Min(value int) *ELEMENT { return ele.Attr("min", value) }

func (ele *ELEMENT) Max(value int) *ELEMENT { return ele.Attr("max", value) }

func (ele *ELEMENT) Height(value string) *ELEMENT { return ele.Attr("height", value) }

func (ele *ELEMENT) Src(src string) *ELEMENT { return ele.Attr("src", src) }

func (ele *ELEMENT) Width(eval string) *ELEMENT { return ele.Attr("width", eval) }

func (ele *ELEMENT) OnError(eval string) *ELEMENT { return ele.Attr("onerror", eval) }
