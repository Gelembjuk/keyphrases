package keyphrases

/**
*
 */
import (
	"errors"
	"fmt"
	"os"

	"github.com/gelembjuk/keyphrases/helper"
	"github.com/gelembjuk/keyphrases/languages"
)

type TextPhrases struct {
	text       string
	Language   string
	NewsSource bool
	langobj    languages.LangClass
}

func (obj *TextPhrases) Init() error {
	if obj.Language == "" {
		obj.Language = "english"
	}

	if obj.Language == "english" {
		obj.langobj = new(languages.English)
	} else {
		return errors.New("Unknown Language")
	}

	return nil
}

func (obj *TextPhrases) GetKeyPhrases(text string) []string {
	obj.text = text

	//phrases := []string{}

	// 1. Split a text for sentences
	// 2. Normalize sentences
	sentences, _ := obj.splitTextForSentences(text)
	// 3. Get words normalized
	words, _ := obj.splitSentencesForWords(sentences)
	// 4. Get phrases from sentences using words

	phraseslist, _ := obj.getPhrases(sentences, words)

	for _, p := range phraseslist {
		fmt.Printf("%s\n", p)
	}
	os.Exit(0)

	return sentences
}

func (obj *TextPhrases) GetKeyWords(text string) []string {
	obj.text = text

	sentences, _ := obj.splitTextForSentences(text)

	wordshash, _ := obj.splitSentencesForWords(sentences)

	words := helper.KeysSortedByValuesReverse(wordshash)

	return words
}

func getAnalyserForTesting() TextPhrases {
	analyser := TextPhrases{Language: "english", NewsSource: true}

	analyser.Init()

	return analyser
}
