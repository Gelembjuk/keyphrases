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

func (obj *TextPhrases) compareSlices(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
