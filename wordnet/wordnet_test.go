package wordnet

import (
	"reflect"
	"testing"

	"github.com/gelembjuk/keyphrases/helper"
)

func TestGetWordOptions(t *testing.T) {

	wordnet := WordNet{DictLocationDirectory: "dict/"}

	tests := map[string][]string{
		"bet": []string{"v", "n"},
		"car": []string{"n"},
	}

	for word, types := range tests {

		wordoptions, err := wordnet.GetWordOptions(word)

		if err != nil {
			t.Fatalf("WordNet Error %s.", err.Error())
		}

		if !helper.CompareSlices(wordoptions, types) {
			t.Fatalf("For word %s, got slice %s, expected %s.", word, wordoptions, types)
		}
	}
}

func TestGetWordOptionsMap(t *testing.T) {

	wordnet := WordNet{DictLocationDirectory: "dict/"}

	tests := map[string]map[string]int{
		"bet": map[string]int{"n": 2, "v": 3},
		"car": map[string]int{"n": 5},
	}

	for word, types := range tests {

		wordoptions, err := wordnet.GetWordOptionsMap(word)

		if err != nil {
			t.Fatalf("WordNet Error %s.", err.Error())
		}

		if !reflect.DeepEqual(wordoptions, types) {
			t.Fatalf("For word %s, got slice %s, expected %s.", word, wordoptions, types)
		}
	}
}

func TestGetWordSynonims(t *testing.T) {

	wordnet := WordNet{DictLocationDirectory: "dict/"}

	tests := map[string][]string{
		"car": []string{
			"car",
			"auto",
			"automobile",
			"machine",
			"motorcar",
			"railcar",
			"railway_car",
			"railroad_car",
			"gondola",
			"elevator_car",
			"cable_car",
		},
		"bet": []string{
			"bet",
			"stake",
			"stakes",
			"wager",
			"play",
			"count",
			"depend",
			"look",
			"calculate",
			"reckon",
		},
	}

	for word, syns := range tests {

		wordsyns, err := wordnet.GetWordSynonims(word)

		if err != nil {
			t.Fatalf("Error %s.", err.Error())
		}

		if !helper.CompareSlices(wordsyns, syns) {
			t.Fatalf("For word %s, got slice %s, expected %s.", word, wordsyns, syns)
		}
	}
}
