package languages

import (
	"fmt"
	"regexp"
)

type English struct {
	Lang Language
}

func (lang English) TruncateCommonPhrase(phrase string) string {
	return phrase
}

func (lang English) TruncateCompanyName(phrase string) string {
	return phrase
}
func (lang English) CheckIfIsCompanyName(phrase string) bool {
	return true
}
func (lang English) CheckIfIsSimilar(phrase1 string, phrase2 string) uint8 {
	return 0
}
func (lang English) NormalizeText(text string) string {
	return ""
}

func (lang *English) CleanNewsMessage(text string) (string, string, error) {
	removed := ""
	/*
		$data=~s/Inc\.\s*?\(/Inc (/g;
		$data=~s/\([^)]+\)//g;
		$data=~s/'s\s/ /g;
	*/

	return text, removed, nil
}

func (lang *English) cleanNewsMessagePrefix(text string) (string, string, error) {

	removed := ""

	templates := [][]string{
		{"^(.{3,30})\\s?-(-|\\s|-\\s)(.*?)$", "1", "3"},
		{"^By (.{3,40}) --? (.*?)$", "1", "2"},
		{"^([A-Z0-9 ]{3,30})\\s?—\\s?(.*?)$", "1", "2"},
		{"^([A-Z0-9 ]{3,30}):\\s(.*?)$", "1", "2"},
		{"^(.{3,30},[A-Z0-9 /]{3,30}):\\s(.*?)$", "1", "2"},
		{"^(In brief):\\s(.*?)$", "1", "2"},
		{"^(\\(.{3,30}\\)\\s.{3,25})-(.*?)$", "1", "2"},
		{"^(.{10,40}\\s/.{3,25}/)\\s--\\s(.*?)$", "1", "2"},
		{"^(\\[.{3,25}\\])\\s?(.*?)$", "1", "2"},
	}

	for _, template := range templates {
		r := regexp.MustCompile(template[0])

		matched := r.FindStringSubmatch(text)
		fmt.Println(matched)
	}

	return text, removed, nil
}
