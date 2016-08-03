package keyphrases

import (
	"strings"

	"gopkg.in/neurosnap/sentences.v1"
	"gopkg.in/neurosnap/sentences.v1/data"
)

func (obj *TextPhrases) splitTextForSentences(text string) ([]string, error) {
	// prepare tokenizer
	sentenceslist := []string{}

	langfile := "data/" + obj.Language + ".json"

	b, err := data.Asset(langfile)

	if err != nil {
		return sentenceslist, err
	}

	// load the training data
	training, err := sentences.LoadTraining(b)

	if err != nil {
		return sentenceslist, err
	}

	// create the default sentence tokenizer
	tokenizer := sentences.NewSentenceTokenizer(training)

	text, _ = obj.cleanTextAfterHTML(text)

	sentencesobjs := tokenizer.Tokenize(text)

	for _, s := range sentencesobjs {
		sentence := s.Text

		// remove last symbol of a sentence if it is a dot or so
		if len(sentence) < 3 {
			continue
		}

		lastsymbol := sentence[len(sentence)-1:]

		if strings.ContainsAny(lastsymbol, ".?!)]:") {
			sentence = sentence[0 : len(sentence)-1]
		}

		sentenceslist = append(sentenceslist, sentence)
	}

	return sentenceslist, nil
}
