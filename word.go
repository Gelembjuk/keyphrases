package keyphrases

import (
	"regexp"
	"strings"
)

func (obj *TextPhrases) splitSentencesForWords(sentences []string) (map[string]int, error) {

	allwords := map[string]int{}

	for _, sentence := range sentences {
		words, _ := obj.splitSentenceForWords(sentence)

		for _, word := range words {
			if _, ok := allwords[word]; ok {
				allwords[word]++
			} else {
				allwords[word] = 1
			}
		}

	}

	obj.langobj.RemoveCommonWords(allwords)

	return allwords, nil
}

func (obj *TextPhrases) splitSentenceForWords(sentence string) ([]string, error) {
	words := []string{}

	replace := [][]string{
		{"[():,;.]", " "},
		{"\\.\\.\\.", " "},
		{"\\s\\s+", " "},
		{"^\\s+", ""},
		{"\\s+$", ""},
	}

	for _, template := range replace {
		r := regexp.MustCompile(template[0])

		sentence = r.ReplaceAllString(sentence, template[1])
	}

	twords := strings.Split(sentence, " ")

	for _, w := range twords {
		if obj.langobj.IsWord(w) {
			words = append(words, w)
		}
	}

	return words, nil
}
