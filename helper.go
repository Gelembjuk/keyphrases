package keyphrases

import (
	"regexp"
	"strings"
)

func (obj *TextPhrases) cleanTextAfterHTML(text string) (string, error) {
	toreplace := map[string]string{
		"\r":      "",
		"&nbsp;":  " ",
		"&amp;":   "&",
		"&#39;":   "'",
		"&lrm;":   " ",
		"&quot;":  "\"",
		"&#8212;": "-",
	}

	for key, value := range toreplace {
		text = strings.Replace(text, key, value, -1)
	}

	r := regexp.MustCompile("^[\\s\\n\\t]*")

	text = r.ReplaceAllString(text, "")

	r = regexp.MustCompile("[\\s]*?$")

	text = r.ReplaceAllString(text, "")

	return text, nil
}
