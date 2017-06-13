package keyphrases

/**
*
 */
import (
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

type Phrase phrases.Phrase
type PhrasesList phrases.PhrasesList
type InPhrase phrases.InPhrase

func (obj *TextPhrases) Init() error {
	var err error
	obj.langobj, err = languages.GetLangObject(obj.Language)

	if err == nil {
		if len(obj.LanguageOptions) > 0 {
			obj.langobj.SetOptions(obj.LanguageOptions)
		}

		words.SetLangObject(obj.langobj)
		sentences.SetLangObject(obj.langobj)
		phrases.SetLangObject(obj.langobj)

	}

	return err
}

func (obj *TextPhrases) GetKeyPhrases(text string) PhrasesList {
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

	phraseslisttmp, _ := phrases.GetPhrases(sentenceslist, wordslist)

	phraseslist := PhrasesList{}

	for _, i := range phraseslisttmp {
		phraseslist = append(phraseslist, i)
	}

	return phraseslist
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

// returns inclusion of predefined key phrases from the text
// predefined phrases is a hash array where each phrase has a list of sinonims
func (obj *TextPhrases) GetKeyPhrasesFromList(text string, inphrases []InPhrase) PhrasesList {
	obj.text = text

	var sentenceslist []string

	if obj.NewsSource {
		sentenceslist, _ = sentences.SplitTextForSentencesFromNews(text)
	} else {
		sentenceslist, _ = sentences.SplitTextForSentences(text)
	}
	// convert to internal list
	inphrasesconv := []phrases.InPhrase{}

	for _, ph := range inphrases {
		phconv := phrases.InPhrase{ph.Phrase, ph.Synonims}
		inphrasesconv = append(inphrasesconv, phconv)
	}

	phraseslisttmp, _ := phrases.GetPhrasesByPredefinedList(sentenceslist, inphrasesconv)

	phraseslist := PhrasesList{}

	for _, i := range phraseslisttmp {
		phraseslist = append(phraseslist, i)
	}

	return phraseslist
}
