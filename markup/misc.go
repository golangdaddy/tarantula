package g

import 	(
		"fmt"
		)

func TAG(args ...string) *ELEMENT {

	ele := MdROW("space-between center").Class("ngTag")
	// .Attr("md-ink-ripple", "true").Style("position:relative;")

	// only supposed to take 1 arg or 0 so this makes sense to return if there is 1 arg. If not, it just returns ele
	for _, arg := range args { return ele.Class(arg) }

	return ele
}

// 430x260 is the card media area size
func LLMAP(size string, zoom int, latKey, lngKey string) *ELEMENT {

	prefix := "https://media.leadinglocally.com/api/map/" + fmt.Sprintf("%v/%v/", size, zoom)

	return IMG().NgSrc(prefix + latKey + "/" + lngKey)
}

func LLPreviewTag(display, tagType, endorsements string ) *ELEMENT {

	component := TAG().Class("ngPreviewTag "+tagType).Add(
		SPAN(display).Class("display").Align("center center"),
	)

	e := endorsements

	if len(e) > 0 {
		component.Add(
			MdICON("favorite").Class("nudge-3 light"),
			SPAN(e).Class("endorsements"),
		)
	}

	return component
}

func LLLocationTag(location, flag string) *ELEMENT {

	return MdROW("center center").Class("ngLocationTag").Add(
		MdICON("location_on").Class("nudge-5 orange"),
		SPAN().Class("flag").BackgroundImage("https://media.leadinglocally.com/api/flag/small/"+flag, "cover", "center"),
		SPAN(location).Class("location-display"),
	)
}
