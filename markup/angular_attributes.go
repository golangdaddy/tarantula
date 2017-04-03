package g

import 	(
		)

func (ele *ELEMENT) NgApp(eval string) *ELEMENT { return ele.Attr("ng-app", eval) }

func (ele *ELEMENT) NgBind(eval string) *ELEMENT { return ele.Attr("ng-bind", eval) }

func (ele *ELEMENT) NgController(args ...string) *ELEMENT {

	switch len(args) {

		case 1:		return ele.Attr("ng-controller", args[0])

		case 2:
          ele.Lock()
						ele.Ctrl = args[1]
					ele.Unlock()
					return ele.Attr("ng-controller", args[0])

	}

	panic("*ELEMENT.NgController")

}

func (ele *ELEMENT) NgClass(eval string) *ELEMENT { return ele.Attr("ng-class", eval) }

func (ele *ELEMENT) NgChecked(eval string) *ELEMENT { return ele.Attr("ng-checked", eval) }

func (ele *ELEMENT) NgCloak() *ELEMENT { return ele.Attr("ng-cloak") }

func (ele *ELEMENT) NgDisabled(eval string) *ELEMENT { return ele.Attr("ng-disabled", eval) }

func (ele *ELEMENT) NgEnter(eval string) *ELEMENT { return ele.Attr("ng-enter", eval) }

func (ele *ELEMENT) NgIf(eval string) *ELEMENT { return ele.Attr("ng-if", eval) }

func (ele *ELEMENT) NgInclude(eval string) *ELEMENT { return ele.Attr("ng-include", eval) }

func (ele *ELEMENT) NgInit(eval string) *ELEMENT { return ele.Attr("ng-init", eval) }

func (ele *ELEMENT) NgHide(eval string) *ELEMENT { return ele.Attr("ng-hide", eval) }

func (ele *ELEMENT) NgModel(eval string) *ELEMENT { return ele.Attr("ng-model", eval) }

func (ele *ELEMENT) NgOptions(eval string) *ELEMENT { return ele.Attr("ng-options", eval) }

func (ele *ELEMENT) NgRepeat(eval string) *ELEMENT { return ele.Attr("ng-repeat", eval) }

func (ele *ELEMENT) NgReadOnly(eval string) *ELEMENT { return ele.Attr("ng-readonly", eval) }

func (ele *ELEMENT) NgShow(eval string) *ELEMENT { return ele.Attr("ng-show", eval) }

func (ele *ELEMENT) NgSrc(src string) *ELEMENT { return ele.Attr("ng-src", src) }

func (ele *ELEMENT) NgStyle(eval string) *ELEMENT { return ele.Attr("ng-style", eval) }

func (ele *ELEMENT) NgSubmit(eval string) *ELEMENT { return ele.Attr("ng-submit", eval) }

func (ele *ELEMENT) NgSwitch(eval string) *ELEMENT { return ele.Attr("ng-switch", eval) }

func (ele *ELEMENT) NgTransclude(eval string) *ELEMENT { return ele.Attr("ng-transclude", eval) }

func (ele *ELEMENT) NgValue(eval string) *ELEMENT { return ele.Attr("ng-value", eval) }
