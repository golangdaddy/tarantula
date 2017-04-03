package g

import 	(
		"fmt"
		)

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Helper functions
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

// <div layout="row" class="gaml_row"></div

func MdROW(args ...string) *ELEMENT {

	ele := Ele("div").Layout("row").Class("g_row")

	for _, arg := range args { return ele.Align(arg) }

	return ele
}

// <div layout="column" class="gaml_row"></div

func MdCOL(args ...string) *ELEMENT {

	ele := Ele("div").Layout("column").Class("g_col")

	for _, arg := range args { return ele.Align(arg) }

	return ele
}

// <ng-md-icon size='24' icon="*"></ng-md-icon>

func MdICON(args ...string) *ELEMENT {
  if len(args) == 2 {
  return Ele("ng-md-icon").Attr("size", args[1]).Attr("icon", args[0])
  }
  return Ele("ng-md-icon").Attr("size", "24").Attr("icon", args[0])
}

func MdINPUTCONTAINER() *ELEMENT { return Ele("md-input-container") }

func MdINPUT(args ...string) *ELEMENT {

	label := args[0]
	model := args[1]

	input := Ele("input").Attr("type", "text").Attr("ng-model", model)
	if len(args) == 4 { input.Attr(args[2], args[3]) }

	return MdINPUTCONTAINER().Add(
		Ele("label").Inner(label),
		input,
	)
}

func MdINPUTINT(args ...string) *ELEMENT {

	label := args[0]
	model := args[1]

	input := Ele("input").Type("number").NgModel(model)
	if len(args) == 4 { input.Attr(args[2], args[3]) }

	return MdINPUTCONTAINER().Add(
		Ele("label").Inner(label),
		input,
	)
}

func MdINPUTFLOAT(args ...string) *ELEMENT {

	label := args[0]
	model := args[1]

	input := Ele("input").Type("number").Attr("step", "any").NgModel(model)
	if len(args) == 4 { input.Attr(args[2], args[3]) }

	return MdINPUTCONTAINER().Add(
		Ele("label").Inner(label),
		input,
	)
}

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Angular Material ELEMENTs - sorted alphabetically
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

// <md-button></md-button>

func MdBUTTON(args ...string) *ELEMENT {

	ele := Ele("md-button")

	for _, arg := range args { return ele.Inner(arg) }

	return ele
}
//<md-checkbox></md-checkbox>

func MdCHECKBOX(model string) *ELEMENT { return Ele("md-checkbox").Attr("ng-model", model) }

//<md-subheader></md-subheader>
func MdSUBHEADER() *ELEMENT { return Ele("md-subheader") }

// <md-content></md-content>

func MdCONTENT() *ELEMENT { return Ele("md-content") }

//<md-datepicker></md-datepicker>

func MdDATEPICKER() *ELEMENT { return Ele("md-datepicker") }

// <md-divider></md-divider>

func MdDIVIDER(args ...int) *ELEMENT {

	ele := Ele("md-divider")

	for _, arg := range args { return ele.Style(fmt.Sprintf("margin:%vpx 0px", arg)) }

	return ele
}

// <md-menu></md-menu>

func MdMENU() *ELEMENT { return Ele("md-menu") }

// <md-menu-content></md-menu-content>

func MdMENUCONTENT() *ELEMENT { return Ele("md-menu-content") }

// <md-menu-item></md-menu-item>

func MdMENUITEM() *ELEMENT { return Ele("md-menu-item") }

// <md-tab></md-tab>

func MdTAB(label string) *ELEMENT { return Ele("md-tab").Attr("label", label) }

// <md-tabs></md-tabs>

func MdTABS() *ELEMENT { return Ele("md-tabs") }

// <md-tab-label></md-tab-label>
func MdTABLABEL(label string) *ELEMENT { return Ele("md-tab-label").Inner(label) }

// <md-tab-body></md-tab-body>
func MdTABBODY() *ELEMENT { return Ele("md-tab-body")}

// <md-tab-content> </md-tab-content>
func MdTABCONTENT() *ELEMENT { return Ele("md-tab-content") }

//<md-progress-circular md-mode="*" value="*"></md-progress-circular>

func MdPROGRESSCIRCULAR(mode, value string) *ELEMENT {

	if len(value) == 0 { return Ele("md-progress-circular").Attr("mode", mode) }

	return Ele("md-progress-circular").Attr("mode", mode).Attr("value", value)
}

//<md-progress-linear md-mode="*" value="*"></md-progress-circular>

func MdPROGRESSLINEAR(mode, value string) *ELEMENT {

	if len(value) == 0 { return Ele("md-progress-linear").Attr("mode", mode) }

	return Ele("md-progress-linear").Attr("mode", mode).Attr("value", value)
}

//<md-slider ng-submit="someFunc()"></md-slider>

func MdSLIDER() *ELEMENT { return Ele("md-slider") }

//<md-slider-container ng-submit="someFunc()"></md-slider-container>

func MdSLIDERCONTAINER() *ELEMENT { return Ele("md-slider-container") }

// <md-select ng-model="model"></md-select>

func MdSELECT(model, placeholder string) *ELEMENT { return Ele("md-select").NgModel(model).Placeholder(placeholder) }

// <md-option></md-option>

func MdOPTION() *ELEMENT { return Ele("md-option") }

// <md-sidenav></md-sidenav>

func MdSIDENAV(componentId string) *ELEMENT { return Ele("md-sidenav").Attr("md-component-id", componentId) }

// <md-toolbar></md-toolbar>

func MdTOOLBAR() *ELEMENT { return Ele("md-toolbar").Class("md-theme-light") }

//<md-tooltip></md-tooltip>

func MdTOOLTIP(label string) *ELEMENT { return Ele("md-tooltip").Inner(label) }

// <md-dialog></md-dialog>
func MdDIALOG() *ELEMENT { return Ele("md-dialog")}

// <md-dialog-content> </md-dialog-content>
func MdDIALOGCONTENT() *ELEMENT { return Ele("md-dialog-content")}

// <md-dialog-actions> </md-dialog-actions>
func MdDIALOGACTIONS() *ELEMENT { return Ele("md-dialog-actions")}

// <md-list> </md-list>
func MdLIST() *ELEMENT { return Ele("md-list")}

// <md-listiitem> </md-list-item>
func MdLISTITEM() *ELEMENT { return Ele("md-list-item")}
