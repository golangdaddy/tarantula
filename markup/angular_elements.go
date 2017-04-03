package g

import 	(
		)

// <ng-include src="*"></ng-include>

func NgINCLUDE(src string) *ELEMENT { return Ele("ng-include").Attr("src", src) }

func NgSWITCH(on string) *ELEMENT { return Ele("ng-switch").Attr("on", on) }

func NgVIEW() *ELEMENT { return Ele("ng-switch") }

func NgPLURALIZE(count, when, offset string) *ELEMENT { return Ele("ng-pluralize").Attr("count", count).Attr("when", when).Attr("offset", offset) }

func NgDIRECTIVE(name string) *ELEMENT { return Ele(name) }

func NgTHUMB(file, width string) *ELEMENT {
  return Ele("div").Attr("ng-thumb", "{file: "+file+", width: "+width+"}")
}
