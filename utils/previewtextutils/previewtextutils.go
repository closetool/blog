package previewtextutils

import (
	"github.com/anaskhan96/soup"
	"github.com/closetool/blog/utils/collectionsutils"
)

func GetText(html string, length int) string {
	if html == "" {
		return ""
	}
	text := soup.HTMLParse(html).FullText()
	if length <= 0 {
		return text
	} else {
		return collectionsutils.Abbreviate(text, length)
	}
}
