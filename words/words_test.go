package words

import (
	"reflect"
	"testing"

	"github.com/gelembjuk/keyphrases/helper"
)

func TestSplitSentenceForWords(t *testing.T) {
	SetLanguage("english")

	tests := map[string][]string{}

	tests["We are heartbroken and in shock over  the loss of Brodie  Copeland"] = []string{"We", "are", "heartbroken", "and", "in", "shock", "over", "the", "loss", "of", "Brodie", "Copeland"}

	for testtext, finalwords := range tests {

		if words, _ := SplitSentenceForWords(testtext); !helper.CompareSlices(words, finalwords) {
			t.Fatalf("For text %s, got %s, expected %s.", testtext, words, finalwords)
		}
	}
}

func TestJoinSimilarWords(t *testing.T) {
	SetLanguage("english")

	words := map[string]int{
		"testword1": 2,
		"US":        3,
		"USA":       4,
		"testword2": 3,
		"xxx":       2,
		"testword3": 2,
		"zzz":       4,
	}

	result := map[string]int{
		"USA":       7,
		"testword1": 7,
		"xxx":       2,
		"zzz":       4,
	}

	joinSimilarWords(words)

	if !reflect.DeepEqual(result, words) {
		t.Fatalf("List not equel, got %s, expected %s.", words, result)
	}

}
