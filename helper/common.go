package helper

import (
	"regexp"
	"strings"
	"unicode"
)

func CleanTextAfterHTML(text string) (string, error) {
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

func CompareSlices(s1 []string, s2 []string) bool {
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

func StringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func UpperCaseFirstLetter(str string) string {
	if len(str) < 1 {
		return str
	}

	out := []rune(str)
	out[0] = unicode.ToUpper(out[0])
	return string(out)
}

func GetFirstLetter(str string) string {
	if len(str) < 1 {
		return str
	}

	out := []rune(str)
	return string(out[0])
}

func GetIndexOfMaxInSlice(slice []int) int {
	max := 0
	index := -1

	for i, v := range slice {
		if v > max {
			index = i
			max = v
		}
	}

	return index
}

func AverageFloat32(arr []float32) float32 {
	if len(arr) == 0 {
		return 0
	}

	total := float32(0.0)

	for _, v := range arr {
		total += v
	}

	return total / float32(len(arr))
}
