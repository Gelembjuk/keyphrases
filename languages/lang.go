package languages

type LangClass interface {
	TruncateCommonPhrase(phrase string) string
	TruncateCompanyName(phrase string) string
	CheckIfIsCompanyName(phrase string) bool
	CheckIfIsSimilar(phrase1 string, phrase2 string) uint8
	NormalizeText(text string) string
	CleanNewsMessage(text string) (string, string, error)
	CleanAndNormaliseSentence(sentence string) (string, error)
	IsWord(word string) bool
	RemoveCommonWords(words map[string]int) bool
}

type Language struct {
	Name string
}

func stringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
