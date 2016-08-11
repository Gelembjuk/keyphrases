package wordnet

import (
	"testing"

	"github.com/gelembjuk/keyphrases/helper"
)

func TestGetWordOptions(t *testing.T) {

	wordnet := WordNet{DictLocationDirectory: "dict/"}

	const finaltext = "FINAL TEXT"

	tests := map[string][]string{
		"bet": []string{"n", "v"},
		"car": []string{"n"},
	}

	for word, types := range tests {

		wordoptions, err := wordnet.GetWordOptions(word)

		if err != nil {
			t.Fatalf("WordNet Error %s.", err.Error())
		}

		if !helper.CompareSlices(wordoptions, types) {
			t.Fatalf("For word %s, got slice %s, expected.", word, wordoptions, types)
		}
	}
}
