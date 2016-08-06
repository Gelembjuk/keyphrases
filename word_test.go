package keyphrases

import "testing"

func TestSplitSentenceForWords(t *testing.T) {
	analyser := TextPhrases{Language: "english", NewsSource: true}

	analyser.Init()

	tests := map[string][]string{}

	tests["We are heartbroken and in shock over  the loss of Brodie  Copeland"] = []string{"We", "are", "heartbroken", "and", "in", "shock", "over", "the", "loss", "of", "Brodie", "Copeland"}

	for testtext, finalwords := range tests {

		if words, _ := analyser.splitSentenceForWords(testtext); !analyser.compareSlices(words, finalwords) {
			t.Fatalf("For text %s, got %s, expected %s.", testtext, words, finalwords)
		}
	}
}
