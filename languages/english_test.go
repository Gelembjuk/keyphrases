package languages

import "testing"

func TestCleanNewsMessagePrefix(t *testing.T) {
	const finaltext = "FINAL TEXT"

	tests := []string{
		"AAA BBB--",
		"AAA BBB-- ",
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

func TestCleanNewsMessage(t *testing.T) {

	tests := [][]string{
		{"AAA BBB--FINAL TEXT", "FINAL TEXT"},
		{"AAA BBB-- FINAL TEXT", "FINAL TEXT"},
		{"AAA BBB- FINAL TEXT", "FINAL TEXT"},
		{"By John Smith - FINAL TEXT", "FINAL TEXT"},
		{"By John Smith-- FINAL TEXT", "FINAL TEXT"},
		{"By John Smith-- FINAL TEXT", "FINAL TEXT"},
		{"By John Smith-- Inc.(FINAL TEXT", "Inc (FINAL TEXT"},
		{"By John Smith-- Inc. (FINAL TEXT", "Inc (FINAL TEXT"},
		{"By John Smith-- FINAL Inc. (TEXT", "FINAL Inc (TEXT"},
		{"AAA BBB- FINAL (Xxxxxx) TEXT", "FINAL  TEXT"},
		{"AAA BBB- FINAL ( abd abc ) TEXT", "FINAL  TEXT"},
		{"AAA BBB-- FINAL TEXT John's", "FINAL TEXT John"},
		{"AAA BBB-- FINAL John's TEXT", "FINAL John TEXT"},
	}

	eng := English{}

	for _, te := range tests {
		testtext := te[0]
		finaltext := te[1]
		if text, _, _ := eng.CleanNewsMessage(testtext); text != finaltext {
			t.Fatalf("For text %s, got %s, expected %s.", testtext, text, finaltext)
		}
	}
}
