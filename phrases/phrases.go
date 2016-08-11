package phrases

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gelembjuk/keyphrases/helper"
	"github.com/gelembjuk/keyphrases/languages"
	"github.com/gelembjuk/keyphrases/words"
)

type Phrase struct {
	Phrase   string
	Synonims []string
	Count    int
}

var langobj languages.LangClass

func SetLangObject(lang languages.LangClass) {
	langobj = lang
}

func SetLanguage(lang string) {
	langobj, _ = languages.GetLangObject(lang)
}

func (p Phrase) String() string {
	result := p.Phrase

	if len(p.Synonims) > 0 {
		result = result + " (" + strings.Join(p.Synonims, ", ") + ")"
	}

	result = result + " [" + strconv.Itoa(p.Count) + "]"

	return result
}

func GetPhrases(sentences []string, allwords map[string]int) ([]Phrase, error) {
	phrases, _ := getBasicPhrasesHash(sentences, allwords)

	removeCommonPhrases(phrases)

	synonims := findSinonimPhrases(phrases)

	findWordsAsPhrases(phrases, allwords)

	phraseslist := []Phrase{}

	finalphrases := finalFilterPhrases(phrases, 12)

	for _, phrase := range finalphrases {

		if _, ok := synonims[phrase]; !ok {
			synonims[phrase] = []string{}
		}
		phraseext := Phrase{Phrase: phrase, Synonims: synonims[phrase], Count: phrases[phrase]}
		phraseslist = append(phraseslist, phraseext)
	}

	return phraseslist, nil
}

func GetPhrasesShort(sentences []string, allwords map[string]int) ([]string, error) {
	phrases, _ := getBasicPhrasesHash(sentences, allwords)

	removeCommonPhrases(phrases)

	findWordsAsPhrases(phrases, allwords)

	finalphrases := finalFilterPhrases(phrases, 12)

	return finalphrases, nil
}

func trimCommonWords(phrase string, mode int8) string {

	wordslist, _ := words.SplitSentenceForWords(phrase)

	for len(wordslist) > 0 && langobj.IsNotUsefulWord(wordslist[0]) {
		wordslist = wordslist[1:len(wordslist)]
	}

	for len(wordslist) > 0 && langobj.IsNotUsefulWord(wordslist[len(wordslist)-1]) {
		wordslist = wordslist[0 : len(wordslist)-1]
	}
	return strings.Join(wordslist, " ")
}

func getBasicPhrasesHash(sentences []string, allwords map[string]int) (map[string]int, error) {
	// keep a previous words history in  phrase
	allwordsh := map[string][]string{}

	// result list of phrases
	phrases := map[string]int{}

	for _, sentence := range sentences {
		wordslist, _ := words.SplitSentenceForWords(sentence)

		prevphrases := []string{}

		pphrases := []string{}

		for _, word := range wordslist {
			wordaddedtosomephrase := 0
			prevwordaddedtosomephrase := 0

			if len(prevphrases) > 0 {

				for i := 0; i < len(prevphrases); i++ {
					prevword := prevphrases[i]

					addedword := 0

					if _, ok := allwordsh[prevword]; !ok {
						//add this for all secuences

						allwordsh[prevword] = []string{word}
					} else {
						if helper.StringInSlice(word, allwordsh[prevword]) {

							//this is phrase and it occured 2 times now
							// it means in some previous sentence this 2 words also were one after other

							// build possible phrase
							if pphrases[i] == "" {
								pphrases[i] = prevword + " " + word
								prevwordaddedtosomephrase = 1
							} else {
								pphrases[i] = pphrases[i] + " " + word
							}
							addedword = 1
						} else {
							//add to array
							allwordsh[prevword] = append(allwordsh[prevword], word)
						}
					}
					if addedword == 0 && pphrases[i] != "" {
						// phrase building complete. This is final step
						// simplify phrase
						pphrases[i] = trimCommonWords(pphrases[i], 0)

						if words.WordsCount(pphrases[i]) > 1 {
							//this is the end of the phrase

							if _, ok := phrases[pphrases[i]]; ok {
								phrases[pphrases[i]]++
							} else {
								phrases[pphrases[i]] = 2
							}
						}
						pphrases[i] = ""
					}
					if addedword == 1 {
						wordaddedtosomephrase = 1
					}
				}
			}

			if prevwordaddedtosomephrase == 1 {
				//it is needed to remove 1 occurence of this word from list of most used words
				if len(prevphrases) > 0 && prevphrases[len(prevphrases)-1] != "" {
					if _, ok := allwords[prevphrases[len(prevphrases)-1]]; ok {
						allwords[prevphrases[len(prevphrases)-1]]--
					}
				}
			}
			prevphrases = append(prevphrases, word)
			pphrases = append(pphrases, "")

			if wordaddedtosomephrase == 1 {
				//it is needed to remove 1 occurence of this word from list of most used words
				if _, ok := allwords[word]; ok {
					allwords[word]--
				}
			}
		}

		for _, phrase := range pphrases {
			if phrase != "" {
				phrase = trimCommonWords(phrase, 1)
				//this is the end of the phrase
				if words.WordsCount(phrase) > 1 && words.WordsCount(phrase) <= 6 {
					if _, ok := phrases[phrase]; ok {
						phrases[phrase]++
					} else {
						phrases[phrase] = 2
					}
				}
				phrase = ""
			}
		}
	}

	return phrases, nil

}

func finalFilterPhrases(phrases map[string]int, maxcount int) []string {
	// sort phrases by count
	// and get first maxcount real phrases
	phraseslist := []string{}

	for phrase, _ := range phrases {
		ptype := getTypeOfPhrase(phrase)
		if ptype == "n" || ptype == "r" || ptype == "s" || ptype == "f" {
			phraseslist = append(phraseslist, phrase)
		}

		if maxcount > 0 && len(phraseslist) >= maxcount {
			break
		}
	}

	return phraseslist
}

func removeCommonPhrases(phrases map[string]int) bool {

	for p, _ := range phrases {
		hasgood := false

		wordslist, _ := words.SplitSentenceForWords(p)

		for _, w := range wordslist {
			if !langobj.IsNotUsefulWord(w) {
				hasgood = true
			}
		}

		if !hasgood {
			delete(phrases, p)
		}
	}
	return true
}

func findSinonimPhrases(phrases map[string]int) map[string][]string {
	sinonims := map[string][]string{}

	remove := []string{}

	for phrase1, _ := range phrases {
		if helper.StringInSlice(phrase1, remove) {
			continue
		}

		for phrase2, _ := range phrases {
			if phrase1 == phrase2 {
				continue
			}
			if helper.StringInSlice(phrase2, remove) {
				continue
			}

			sres := isSubpraseOfPhrase(phrase1, phrase2)

			if sres > 0 {
				phrases[phrase1] += phrases[phrase2]

				sinonims[phrase1] = append(sinonims[phrase1], phrase2)

				remove = append(remove, phrase2)
			} else if sres < 0 {
				phrases[phrase2] += phrases[phrase1]

				sinonims[phrase2] = append(sinonims[phrase2], phrase1)

				remove = append(remove, phrase1)

				break
			}
		}
	}

	for _, phrase := range remove {
		delete(phrases, phrase)
	}
	return sinonims
}

func findWordsAsPhrases(phrases map[string]int, allwords map[string]int) {
	// add most used words to phrases list
	// check if this word is not used in another word in most cases
	// aim is to find words that are possible company name

	mostappearphrase := float32(helper.GetBiggestValueInMap(phrases)) / 3.0

	for word, count := range allwords {
		if float32(count) > mostappearphrase {
			// check word type
			// NOTE. this is expensive operation
			wtype := langobj.GetTypeOfWord(word, "", "")

			if wtype == "n" || wtype == "r" || wtype == "s" && wtype == "f" {
				phrases[word] = count
			}
		}
	}

}

func normalisePhrase(phrase string) string {
	phrase = strings.ToLower(phrase)

	replace := [][]string{
		{"\\s\\s+", " "},
		{"^\\s+", ""},
		{"\\s+$", ""},
	}

	for _, template := range replace {
		r := regexp.MustCompile(template[0])

		phrase = r.ReplaceAllString(phrase, template[1])
	}

	return phrase
}

func isSubpraseOfPhrase(phrase1 string, phrase2 string) int8 {
	nphrase1 := normalisePhrase(phrase1)
	nphrase2 := normalisePhrase(phrase2)

	if nphrase1 == nphrase2 {
		if phrase1 == strings.ToUpper(phrase1) ||
			strings.ToUpper(phrase2[0:1]) == phrase2[0:1] {
			return 1
		}
		return -1
	}

	check := langobj.IsPhraseSubphrase(nphrase1, nphrase2)

	if check != 0 {
		return check
	}

	nphrase1 = trimCommonWords(nphrase1, 0)
	nphrase2 = trimCommonWords(nphrase2, 0)

	check = langobj.IsPhraseSubphrase(nphrase1, nphrase2)

	if check != 0 {
		return check
	}

	return 0
}

func isSubpraseOfPhraseExtended(phrase1 string, phrase2 string) int8 {
	result := isSubpraseOfPhrase(phrase1, phrase2)

	if result != 0 {
		return result
	}

	if langobj.IsWord(phrase1) && !langobj.IsWord(phrase2) {
		if langobj.IsWordModInPhrase(phrase2, phrase1) {
			return -1
		}
	}

	if langobj.IsWord(phrase2) && !langobj.IsWord(phrase1) {
		if langobj.IsWordModInPhrase(phrase1, phrase2) {
			return 1
		}
	}

	return 0
}

func getTypeOfPhrase(phrase string) string {
	alltypes := []string{}

	wordslist, _ := words.SplitSentenceForWords(phrase)

	l := len(wordslist)

	for i, word := range wordslist {
		t := ""

		if langobj.IsNotUsefulWord(word) {
			t = "b"
		} else {
			prevword := ""
			if i > 0 {
				prevword = wordslist[i-1]
			}
			nextword := ""

			if i < l-1 {
				prevword = wordslist[i+1]
			}
			t = langobj.GetTypeOfWord(word, prevword, nextword)
		}

		alltypes = append(alltypes, t)
	}

	if helper.StringInSlice("n", alltypes) && helper.StringInSlice("v", alltypes) {
		return "s"
	}

	torder := []string{"f", "n", "a", "v", "c", "r", "t"}

	for _, t := range torder {
		if helper.StringInSlice(t, alltypes) {
			return t
		}
	}

	return "r"
}
