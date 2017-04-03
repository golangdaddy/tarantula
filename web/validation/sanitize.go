package validation

import (
	"html"
	"strings"
	//
	"github.com/microcosm-cc/bluemonday"
)

func Sanitize(s string) string {
	
	return html.UnescapeString(strings.TrimSpace(bluemonday.StrictPolicy().Sanitize(s)))	
}