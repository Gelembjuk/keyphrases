package languages

import (
	//	"strings"
	"unicode"
)

type LangClass interface {
	TruncateCommonPhrase(phrase string) string
	TruncateCompanyName(phrase string) string
	CheckIfIsCompanyName(phrase string) bool
	CheckIfIsSimilar(phrase1 string, phrase2 string) uint8
	NormalizeText(text string) string
	CleanNewsMessage(text string) (string, string, error)
	CleanAndNormaliseSentence(sentence string) (string, error)
	IsWord(word string) bool
	RemoveCommonWords(words map[string]int) bool
	IsSimilarWord(word1 string, word2 string) int8
}

type Language struct {
	Name string
}

func stringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func upperCaseFirstLetter(str string) string {
	if len(str) < 1 {
		return str
	}

	out := []rune(str)
	out[0] = unicode.ToUpper(out[0])
	return string(out)
}
