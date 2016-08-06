package keyphrases

import "testing"

func TestCleanAndNormaliseSentence(t *testing.T) {
	analyser := TextPhrases{Language: "english", NewsSource: true}

	analyser.Init()

	tests := [][]string{
		{"We [are] heartbroken and in \"shock\" over the loss of Brodie Copeland.",
			"We are heartbroken and in shock over the loss of Brodie Copeland"},
		{"We {are} heartbroken and in \"shock\" over the: loss of  Brodie Copeland.",
			"We are heartbroken and in shock over the loss of Brodie Copeland"},
		{" We {are} heartbroken   and in \"shock\" over the: loss of Brodie Copeland! ",
			"We are heartbroken and in shock over the loss of Brodie Copeland"},
	}

	for _, te := range tests {
		testtext := te[0]
		finaltext := te[1]
		if text, _ := analyser.cleanAndNormaliseSentence(testtext); text != finaltext {
			t.Fatalf("For text %s, got %s, expected %s.", testtext, text, finaltext)
		}
	}
}
