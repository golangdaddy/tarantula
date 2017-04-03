package g

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Colors, Styles, and Shadows - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) FillStyle(eval string) *ELEMENT { return ele.Attr("fillStyle", eval) }

func (ele *ELEMENT) StrokeStyle(eval string) *ELEMENT { return ele.Attr("strokeStyle", eval) }

func (ele *ELEMENT) ShadowColor(eval string) *ELEMENT { return ele.Attr("shadowColor", eval) }

func (ele *ELEMENT) ShadowBlur(eval string) *ELEMENT { return ele.Attr("shadowBlur", eval) }

func (ele *ELEMENT) ShadowOffsetX(eval string) *ELEMENT { return ele.Attr("shadowOffsetX", eval) }

func (ele *ELEMENT) ShadowOffsetY(eval string) *ELEMENT { return ele.Attr("shadowOffsetY", eval) }

func (ele *ELEMENT) CreateLinearGradient(eval string) *ELEMENT { return ele.Attr("createLinearGradient()", eval) }

func (ele *ELEMENT) CreatePattern(eval string) *ELEMENT { return ele.Attr("createPattern()", eval) }

func (ele *ELEMENT) CreateRadialGradient(eval string) *ELEMENT { return ele.Attr("createRadialGradient()", eval) }

func (ele *ELEMENT) AddColorStop(eval string) *ELEMENT { return ele.Attr("addColorStop()", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Line styles - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) LineCap(eval string) *ELEMENT { return ele.Attr("lineCap", eval) }

func (ele *ELEMENT) LineJoin(eval string) *ELEMENT { return ele.Attr("lineJoin", eval) }

func (ele *ELEMENT) LineWidth(eval string) *ELEMENT { return ele.Attr("lineWidth", eval) }

func (ele *ELEMENT) MiterLimit(eval string) *ELEMENT { return ele.Attr("miterLimit", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Rectangles - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) Rect(eval string) *ELEMENT { return ele.Attr("rect()", eval) }

func (ele *ELEMENT) FillRect(eval string) *ELEMENT { return ele.Attr("fillRect()", eval) }

func (ele *ELEMENT) StrokeRect(eval string) *ELEMENT { return ele.Attr("strokeRect()", eval) }

func (ele *ELEMENT) ClearRect(eval string) *ELEMENT { return ele.Attr("clearRect()", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Paths - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) Fill(eval string) *ELEMENT { return ele.Attr("fill()", eval) }

func (ele *ELEMENT) Stroke(eval string) *ELEMENT { return ele.Attr("stroke()", eval) }

func (ele *ELEMENT) BeginPath(eval string) *ELEMENT { return ele.Attr("beginPath()", eval) }

func (ele *ELEMENT) MoveTo(eval string) *ELEMENT { return ele.Attr("moveTo()", eval) }

func (ele *ELEMENT) ClosePath(eval string) *ELEMENT { return ele.Attr("closePath()", eval) }

func (ele *ELEMENT) LineTo(eval string) *ELEMENT { return ele.Attr("lineTo()", eval) }

func (ele *ELEMENT) Clip(eval string) *ELEMENT { return ele.Attr("clip()", eval) }

func (ele *ELEMENT) QuadraticCurveTo(eval string) *ELEMENT { return ele.Attr("quadraticCurveTo()", eval) }

func (ele *ELEMENT) BezierCurveTo(eval string) *ELEMENT { return ele.Attr("bezierCurveTo()", eval) }

func (ele *ELEMENT) Arc(eval string) *ELEMENT { return ele.Attr("arc()", eval) }

func (ele *ELEMENT) ArcTo(eval string) *ELEMENT { return ele.Attr("arcTo()", eval) }

func (ele *ELEMENT) IsPointInPath(eval string) *ELEMENT { return ele.Attr("isPointInPath()", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Transformations - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) Scale(eval string) *ELEMENT { return ele.Attr("scale()", eval) }

func (ele *ELEMENT) Rotate(eval string) *ELEMENT { return ele.Attr("rotate()", eval) }

// Removed because of conflict
// func (ele *ELEMENT) Translate(eval string) *ELEMENT { return ele.Attr("translate()", eval) }

func (ele *ELEMENT) Transform(eval string) *ELEMENT { return ele.Attr("transform()", eval) }

func (ele *ELEMENT) SetTransform(eval string) *ELEMENT { return ele.Attr("setTransform()", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Text - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) Font(eval string) *ELEMENT { return ele.Attr("font", eval) }

func (ele *ELEMENT) TextAlign(eval string) *ELEMENT { return ele.Attr("textAlign", eval) }

func (ele *ELEMENT) TextBaseline(eval string) *ELEMENT { return ele.Attr("textBaseline", eval) }

func (ele *ELEMENT) FillText(eval string) *ELEMENT { return ele.Attr("fillText()", eval) }

func (ele *ELEMENT) StrokeText(eval string) *ELEMENT { return ele.Attr("strokeText()", eval) }

func (ele *ELEMENT) MeasureText(eval string) *ELEMENT { return ele.Attr("measureText()", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Image drawing - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) DrawImage(eval string) *ELEMENT { return ele.Attr("drawImage()", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Pixel manipulation - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

// Moved to html_sharedAttributes.go because of conflict
// func (ele *ELEMENT) Width(eval string) *ELEMENT { return ele.Attr("width", eval) }

// Moved to html_sharedAttributes.go because of conflict
// func (ele *ELEMENT) Height(eval string) *ELEMENT { return ele.Attr("height", eval) }

func (ele *ELEMENT) Data(eval string) *ELEMENT { return ele.Attr("data", eval) }

func (ele *ELEMENT) CreateImageData(eval string) *ELEMENT { return ele.Attr("createImageData()", eval) }

func (ele *ELEMENT) GetImageData(eval string) *ELEMENT { return ele.Attr("getImageData()", eval) }

func (ele *ELEMENT) PutImageData(eval string) *ELEMENT { return ele.Attr("putImageData()", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Compositing - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) GlobalAlpha(eval string) *ELEMENT { return ele.Attr("globalAlpha", eval) }

func (ele *ELEMENT) GlobalCompositeOperation(eval string) *ELEMENT { return ele.Attr("globalCompositeOperation", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  Other - http://www.w3schools.com/tags/ref_canvas.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) Save(eval string) *ELEMENT { return ele.Attr("save()", eval) }

func (ele *ELEMENT) Restore(eval string) *ELEMENT { return ele.Attr("restore()", eval) }

func (ele *ELEMENT) CreateEvent(eval string) *ELEMENT { return ele.Attr("createEvent()", eval) }

func (ele *ELEMENT) GetContext(eval string) *ELEMENT { return ele.Attr("getContext()", eval) }

func (ele *ELEMENT) ToDataURL(eval string) *ELEMENT { return ele.Attr("toDataURL()", eval) }
