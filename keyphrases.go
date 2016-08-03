package keyphrases

/**
*
 */
import (
	"errors"

	"github.com/gelembjuk/keyphrases/languages"
)

type TextPhrases struct {
	text     string
	Language string
	langobj  languages.LangClass
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

func (obj *TextPhrases) Analyse(text string) []string {
	obj.text = text

	//phrases := []string{}

	// 1. Split a text for sentences
	sentences, _ := obj.splitTextForSentences(text)
	// 2. Normalize sentences
	// 3. Get words normalized
	// 4. Get phrases from sentences using words

	return sentences
}
