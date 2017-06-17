package words

import (
	"regexp"
	"sort"
	"strings"

	"github.com/gelembjuk/keyphrases/languages"
)

var langobj languages.LangClass

func SetLangObject(lang languages.LangClass) {
	langobj = lang
}

func SetLanguage(lang string) {
	langobj, _ = languages.GetLangObject(lang)
}

func SplitSentencesForWords(sentences []string) (map[string]int, error) {

	allwords := map[string]int{}

	sort.Strings(sentences)

	for _, sentence := range sentences {
		words, _ := SplitSentenceForWords(sentence)

		for _, word := range words {

			if _, ok := allwords[word]; ok {
				allwords[word]++
			} else {
				allwords[word] = 1
			}
		}

	}
	langobj.RemoveCommonWords(allwords)

	joinSimilarWords(allwords)

	return allwords, nil
}

func WordsCount(sentence string) uint {
	words, _ := SplitSentenceForWords(sentence)

	return uint(len(words))
}

func SplitSentenceForWords(sentence string) ([]string, error) {

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
		if langobj.IsWord(w) {
			words = append(words, w)
		}
	}

	return words, nil
}

func joinSimilarWords(words map[string]int) bool {

	wordslist := []string{}

	for word, _ := range words {
		wordslist = append(wordslist, word)
	}

	sort.Strings(wordslist)

	remove := [][]string{}

	checkIfRemoved := func(word string) bool {
		for _, w := range remove {
			if w[0] == word || w[1] == word {
				return true
			}
		}
		return false
	}

	for i, word1 := range wordslist {

		if checkIfRemoved(word1) {
			continue
		}

		for j, word2 := range wordslist {
			if i == j {
				continue
			}

			if checkIfRemoved(word2) {
				continue
			}
			similar := langobj.IsSimilarWord(word1, word2)

			if similar == 1 {
				remove = append(remove, []string{word2, word1})
			} else if similar == -1 {
				remove = append(remove, []string{word1, word2})
				continue
			}
		}
	}

	for _, w := range remove {
		words[w[1]] += words[w[0]]
		delete(words, w[0])
	}

	return true
}
