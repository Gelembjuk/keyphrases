package languages

import (
	"errors"
)

type LangClass interface {
	GetName() string
	SetOptions(options map[string]string) error
	CleanNewsMessage(text string) (string, string, error)
	CleanAndNormaliseSentence(sentence string) (string, error)
	IsWord(word string) bool
	RemoveCommonWords(words map[string]int) bool
	IsSimilarWord(word1 string, word2 string) int8
	IsNotUsefulWord(word string) bool
	IsPhraseSubphrase(phrase1 string, phrase2 string) int8
	IsWordModInPhrase(phrase, word string) bool
	GetTypeOfWord(word string) (string, error)
	GetTypeOfWordComplex(word string, prevword string, nextword string) (string, error)
	SimplifyCompanyName(phrase string) string
	SimplifyCompanyNameExt(phrase string) string
}

type Language struct {
	Name string
}

func GetLangObject(language string) (LangClass, error) {
	if language == "" {
		language = "english"
	}

	if language == "english" || language == "en" {
		return new(English), nil
	}

	return nil, errors.New("Unknown Language")

}
