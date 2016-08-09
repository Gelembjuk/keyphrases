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
	IsSimilarWord(word1 string, word2 string) int8
	IsNotUsefulWord(word string) bool
	IsPhraseSubphrase(phrase1 string, phrase2 string) int8
	IsWordModInPhrase(phrase, word string) bool
	GetTypeOfWord(word string, prevword string, nextword string) string
}

type Language struct {
	Name string
}
