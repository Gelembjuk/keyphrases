package keyphrases

/**
*
 */
import (
	"fmt"
	"os"

	"github.com/gelembjuk/keyphrases/helper"
	"github.com/gelembjuk/keyphrases/languages"
	"github.com/gelembjuk/keyphrases/phrases"
	"github.com/gelembjuk/keyphrases/sentences"
	"github.com/gelembjuk/keyphrases/words"
)

type TextPhrases struct {
	text            string
	Language        string
	LanguageOptions map[string]string
	NewsSource      bool
	langobj         languages.LangClass
}

func (obj *TextPhrases) Init() error {
	var err error
	obj.langobj, err = languages.GetLangObject(obj.Language)

	if err == nil {
		words.SetLangObject(obj.langobj)
		sentences.SetLangObject(obj.langobj)
		phrases.SetLangObject(obj.langobj)

		if len(obj.LanguageOptions) > 0 {

		}
	}

	return err
}

func (obj *TextPhrases) GetKeyPhrases(text string) []string {
	obj.text = text

	//phrases := []string{}

	// 1. Split a text for sentences
	// 2. Normalize sentences

	var sentenceslist []string

	if obj.NewsSource {
		sentenceslist, _ = sentences.SplitTextForSentencesFromNews(text)
	} else {
		sentenceslist, _ = sentences.SplitTextForSentences(text)
	}

	// 3. Get words normalized
	wordslist, _ := words.SplitSentencesForWords(sentenceslist)
	// 4. Get phrases from sentences using words

	phraseslist, _ := phrases.GetPhrases(sentenceslist, wordslist)

	for _, p := range phraseslist {
		fmt.Printf("%s\n", p)
	}
	os.Exit(0)

	return sentenceslist
}

func (obj *TextPhrases) GetKeyWords(text string) []string {
	obj.text = text

	var sentenceslist []string

	if obj.NewsSource {
		sentenceslist, _ = sentences.SplitTextForSentencesFromNews(text)
	} else {
		sentenceslist, _ = sentences.SplitTextForSentences(text)
	}

	wordshash, _ := words.SplitSentencesForWords(sentenceslist)

	words := helper.KeysSortedByValuesReverse(wordshash)

	return words
}
