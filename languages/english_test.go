package languages

import "testing"

func TestCleanNewsMessagePrefix(t *testing.T) {
	const finaltext = "FINAL TEXT"

	tests := []string{
		"AAA BBB--",
		"AAA BBB- -",
		"AAA BBB- ",
		"By John Smith - ",
		"By John Smith-- ",
	}

	eng := English{}

	for _, te := range tests {
		testtext := te + finaltext
		if text, _, _ := eng.cleanNewsMessagePrefix(testtext); text != finaltext {
			t.Fatalf("For text %s, got %s.", testtext, text)
		}
	}
}
