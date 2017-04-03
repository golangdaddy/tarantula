package g

import 	(
		"fmt"
		)

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// helper constructors
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

	// <link href="*" rel="icon" type="image/x-icon"></link>

	func FAVICON(href string) *ELEMENT { return LINK().Attr("href", href).Attr("rel", "icon").Attr("type", "image/x-icon") }

	// <link href="*" charset="utf-8" rel="stylesheet" type="text/css"></link>

	func STYLESHEET(href string) *ELEMENT { return LINK().Attr("href", href).Attr("charset", "utf-8").Attr("rel", "stylesheet").Attr("type", "text/css") }

	func INLINESCRIPT(inner string) *ELEMENT { return Ele("script").Inner(inner) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// HTML tags as listed by W3 http://www.w3schools.com/tags/
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

	// <a></a>

	func A(args ...string) *ELEMENT {

		ele := Ele("a")

		for _, arg := range args { ele.Inner(arg) }

		return ele
	}

	// <abbr></abbr>

	func ABBR() *ELEMENT { return Ele("abbr") }

	// <address></address>

	func ADDRESS() *ELEMENT { return Ele("address") }

	// <area></area>

	func AREA() *ELEMENT { return Ele("area") }

	// <article></article>

	func ARTICLE() *ELEMENT { return Ele("article") }

	// <aside></aside>

	func ASIDE() *ELEMENT { return Ele("aside") }

	// <audio></audio>

	func AUDIO() *ELEMENT { return Ele("audio") }

	// <b></b>

	func B(args ...string) *ELEMENT {

		ele := Ele("b")

		for _, arg := range args { ele.Inner(arg) }

		return ele
	}

	// <base></base>

	func BASE() *ELEMENT { return Ele("base") }

	// <bdi></bdi>

	func BDI() *ELEMENT { return Ele("bdi") }

	// <bdo></bdo>

	func BDO() *ELEMENT { return Ele("bdo") }

	// <blockquote></blockquote>

	func BLOCKQUOTE() *ELEMENT { return Ele("blockquote") }

	// <body></body>

	func BODY() *ELEMENT { return Ele("body") }

	// <br></br>

	func BR() *ELEMENT { return Ele("br") }

	// <button></button>

	func BUTTON(args ...string) *ELEMENT {

		ele := Ele("button")

		for _, arg := range args { ele.Inner(arg) }

		return ele
	}

	// <canvas></canvas>

	func CANVAS() *ELEMENT { return Ele("canvas") }

	// <caption></caption>

	func CAPTION() *ELEMENT { return Ele("caption") }

	// <cite></cite>

	func CITE() *ELEMENT { return Ele("cite") }

	// <code></code>

	func CODE(args ...string) *ELEMENT {

		ele := Ele("code")

		for _, arg := range args { ele.Inner(arg) }

		return ele
	}

	// <col></col>

	func COL() *ELEMENT { return Ele("col") }

	// <colgroup></colgroup>

	func COLGROUP() *ELEMENT { return Ele("colgroup") }

	// <datalist></datalist>

	func DATALIST() *ELEMENT { return Ele("datalist") }

	// <dd></dd>

	func DD() *ELEMENT { return Ele("dd") }

	// <del></del>

	func DEL() *ELEMENT { return Ele("del") }

	// <details></details>

	func DETAILS() *ELEMENT { return Ele("details") }

	// <dfn></dfn>

	func DFN() *ELEMENT { return Ele("dfn") }

	// <dialog></dialog>

	func DIALOG() *ELEMENT { return Ele("dialog") }

	// <div></div>

	func DIV(args ...string) *ELEMENT {

		ele := Ele("div")

		for _, arg := range args { return ele.Inner(arg) }

		return ele
	}

	// <dl></dl>

	func DL() *ELEMENT { return Ele("dl") }

	// <dt></dt>

	func DT() *ELEMENT { return Ele("dt") }

	// <em></em>

	func EM() *ELEMENT { return Ele("em") }

	// <embed></embed>

	func EMBED() *ELEMENT { return Ele("embed") }

	// <fieldset></fieldset>

	func FIELDSET() *ELEMENT { return Ele("fieldset") }

	// <figcaption></figcaption>

	func FIGCAPTION() *ELEMENT { return Ele("figcaption") }

	// <figure></figure>

	func FIGURE() *ELEMENT { return Ele("figure") }

	// <footer></footer>

	func FOOTER() *ELEMENT { return Ele("footer") }

	// <form></form>

	func FORM() *ELEMENT { return Ele("form") }

	// <h*></h*>

	func H(args ...interface{}) *ELEMENT {

		ele := Ele("h"+fmt.Sprintf("%v", args[0]))

		if len(args) > 1 {
			ele.Inner(fmt.Sprintf("%v", args[1]))
		}

		return ele
	}

	// <head></head>

	func HEAD() *ELEMENT { return Ele("head") }

	// <header></header>

	func HEADER() *ELEMENT { return Ele("header") }

	// <hr></hr>

	func HR() *ELEMENT { return Ele("hr") }

	// <html></html>

	func HTML() *ELEMENT { return Ele("html") }

	// <i></i>

	func I() *ELEMENT { return Ele("i") }

	// <iframe></iframe>

	func IFRAME() *ELEMENT { return Ele("iframe") }

	// <img></img>

	func IMG() *ELEMENT { return Ele("img") }

	// <input></input>

	func INPUT() *ELEMENT { return Ele("input") }

	// <ins></ins>

	func INS() *ELEMENT { return Ele("ins") }

	// <kbd></kbd>

	func KBD() *ELEMENT { return Ele("kbd") }

	// <keygen></keygen>

	func KEYGEN() *ELEMENT { return Ele("keygen") }

	// <label></label>

	func LABEL() *ELEMENT { return Ele("label") }

	// <legend></legend>

	func LEGEND() *ELEMENT { return Ele("legend") }

	// <li></li>

	func LI(inner ...string) *ELEMENT {
		if len(inner) > 0{
			return Ele("li").Inner(inner[0])
		}else{
			return Ele("li") }
		}

	// <link></link>

	func LINK() *ELEMENT { return Ele("link") }

	// <main></main>

	func MAIN() *ELEMENT { return Ele("main") }

	// <map></map>

	func MAP() *ELEMENT { return Ele("map") }

	// <mark></mark>

	func MARK() *ELEMENT { return Ele("mark") }

	// <menu></menu>

	func MENU() *ELEMENT { return Ele("menu") }

	// <menuitem></menuitem>

	func MENUITEM() *ELEMENT { return Ele("menuitem") }

	// <meta></meta>

	func META() *ELEMENT { return Ele("meta") }

	// <meter></meter>

	func METER() *ELEMENT { return Ele("meter") }

	// <nav></nav>

	func NAV() *ELEMENT { return Ele("nav") }

	// <noscript></noscript>

	func NOSCRIPT() *ELEMENT { return Ele("nav") }

	// <object></object>

	func OBJECT() *ELEMENT { return Ele("object") }

	// <ol></ol>

	func OL() *ELEMENT { return Ele("ol") }

	// <optgroup></optgroup>

	func OPTGROUP() *ELEMENT { return Ele("optgroup") }

	// <option></option>

	func OPTION() *ELEMENT { return Ele("option") }

	// <output></output>

	func OUTPUT() *ELEMENT { return Ele("output") }

	// <p></p>

	func P(args ...string) *ELEMENT {

		ele := Ele("p")

		for _, arg := range args { ele.Inner(arg) }

		return ele
	}

	// <param></param>

	func PARAM() *ELEMENT { return Ele("param") }

	// <pre></pre>

	func PRE(args ...string) *ELEMENT {

		ele := Ele("pre")

		for _, arg := range args { ele.Inner(arg) }

		return ele
	}

	// <progress></progress>

	func PROGRESS() *ELEMENT { return Ele("progress") }

	// <q></q>

	func Q() *ELEMENT { return Ele("q") }

	// <rp></rp>

	func RP() *ELEMENT { return Ele("rp") }

	// <rt></rt>

	func RT() *ELEMENT { return Ele("rt") }

	// <ruby></ruby>

	func RUBY() *ELEMENT { return Ele("ruby") }

	// <s></s>

	func S() *ELEMENT { return Ele("s") }

	// <samp></samp>

	func SAMP() *ELEMENT { return Ele("samp") }

	// <script src="*"></script>

	func SCRIPT(src string) *ELEMENT { return Ele("script").Attr("src", src) }

	// <section></section>

	func SECTION() *ELEMENT { return Ele("section") }

	// <select></select>

	func SELECT() *ELEMENT { return Ele("select") }

	// <small></small>

	func SMALL() *ELEMENT { return Ele("small") }

	// <source src='*' type='*'></source>

	func SOURCE(src, typ string) *ELEMENT { return Ele("source") }

	// <span></span>

		func SPAN(args ...string) *ELEMENT {

			ele := Ele("span")

			for _, arg := range args { ele.Inner(arg) }

			return ele
		}

	// <strong></strong>

	func STRONG() *ELEMENT { return Ele("strong") }

	// <style></style>

	func STYLE() *ELEMENT { return Ele("style") }

	// <sub></sub>

	func SUB() *ELEMENT { return Ele("sub") }

	// <summary></summary>

	func SUMMARY() *ELEMENT { return Ele("summary") }

	// <sup></sup>

	func SUP() *ELEMENT { return Ele("sup") }

	// <table></table>

	func TABLE() *ELEMENT { return Ele("table") }

	// <tbody></tbody>

	func TBODY() *ELEMENT { return Ele("tbody") }

	// <td></td>

	func TD() *ELEMENT { return Ele("td") }

	// <textarea></textarea>

	func TEXTAREA() *ELEMENT { return Ele("textarea") }

	// <tfoot></tfoot>

	func TFOOT() *ELEMENT { return Ele("tfoot") }

	// <th></th>

	func TH() *ELEMENT { return Ele("th") }

	// <thead></thead>

	func THEAD() *ELEMENT { return Ele("thead") }

	// <time></time>

	func TIME() *ELEMENT { return Ele("time") }

	// <title></title>

	func TITLE(args ...string) *ELEMENT {

		ele := Ele("title")

		for _, arg := range args { return ele.Inner(arg) }

		return ele
	}

	// <tr></tr>

	func TR(title string) *ELEMENT { return Ele("tr") }

	// <track></track>

	func TRACK() *ELEMENT { return Ele("track") }

	// <u></u>

	func U() *ELEMENT { return Ele("u") }

	// <ul></ul>

	func UL() *ELEMENT { return Ele("ul") }

	// <var></var>

	func VAR() *ELEMENT { return Ele("var") }

	// <video></video>

	func VIDEO() *ELEMENT { return Ele("video") }

	// <wbr></wbr>

	func WBR() *ELEMENT { return Ele("wbr") }
