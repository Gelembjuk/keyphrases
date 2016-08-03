package languages

type LangClass interface {
	TruncateCommonPhrase(phrase string) string
	TruncateCompanyName(phrase string) string
	CheckIfIsCompanyName(phrase string) bool
	CheckIfIsSimilar(phrase1 string, phrase2 string) uint8
	NormalizeText(text string) string
	CleanNewsMessage(text string) (string, string, error)
}

type Language struct {
	Name string
}
