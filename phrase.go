package keyphrases

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gelembjuk/keyphrases/helper"
)

type Phrase struct {
	Phrase   string
	Synonims []string
	Count    int
}

func (p Phrase) String() string {
	result := p.Phrase

	if len(p.Synonims) > 0 {
		result = result + " (" + strings.Join(p.Synonims, ", ") + ")"
	}

	result = result + " [" + strconv.Itoa(p.Count) + "]"

	return result
}

func (obj *TextPhrases) getPhrases(sentences []string, allwords map[string]int) ([]Phrase, error) {
	phrases, _ := obj.getBasicPhrasesHash(sentences, allwords)

	obj.removeCommonPhrases(phrases)

	synonims := obj.findSinonimPhrases(phrases)

	obj.findWordsAsPhrases(phrases, allwords)

	phraseslist := []Phrase{}

	finalphrases := obj.finalFilterPhrases(phrases, 12)

	for _, phrase := range finalphrases {

		if _, ok := synonims[phrase]; !ok {
			synonims[phrase] = []string{}
		}
		phraseext := Phrase{Phrase: phrase, Synonims: synonims[phrase], Count: phrases[phrase]}
		phraseslist = append(phraseslist, phraseext)
	}

	return phraseslist, nil
}

func (obj *TextPhrases) getPhrasesShort(sentences []string, allwords map[string]int) ([]string, error) {
	phrases, _ := obj.getBasicPhrasesHash(sentences, allwords)

	obj.removeCommonPhrases(phrases)

	obj.findWordsAsPhrases(phrases, allwords)

	finalphrases := obj.finalFilterPhrases(phrases, 12)

	return finalphrases, nil
}

func (obj *TextPhrases) trimCommonWords(phrase string, mode int8) string {

	words, _ := obj.splitSentenceForWords(phrase)

	for len(words) > 0 && obj.langobj.IsNotUsefulWord(words[0]) {
		words = words[1:len(words)]
	}

	for len(words) > 0 && obj.langobj.IsNotUsefulWord(words[len(words)-1]) {
		words = words[0 : len(words)-1]
	}
	return strings.Join(words, " ")
}

func (obj *TextPhrases) getBasicPhrasesHash(sentences []string, allwords map[string]int) (map[string]int, error) {

	allwordsh := map[string][]string{}
	phrases := map[string]int{}

	for _, sentence := range sentences {
		words, _ := obj.splitSentenceForWords(sentence)

		prevphrases := []string{}

		pphrases := []string{}

		for _, word := range words {
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
						// simplify phrase
						pphrases[i] = obj.trimCommonWords(pphrases[i], 0)

						if obj.wordsCount(pphrases[i]) > 1 {
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
				phrase = obj.trimCommonWords(phrase, 1)
				//this is the end of the phrase
				if obj.wordsCount(phrase) > 1 && obj.wordsCount(phrase) <= 6 {
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
	/*


		my @phraseslist=();

		foreach my $c(reverse sort {$phrases{$a} <=> $phrases{$b} } keys %phrases){
			next if(!defined $c || $c=~/^\s+$/ || $c eq '');
			my $ptype=$self->getTypeOfPhrase($c);
			if($ptype eq 'n' || $ptype eq 'r' || $ptype eq 's' || $ptype eq 'f'){
				$synonims{$c}=[] unless (defined $synonims{$c});

				push @phraseslist,[$c,$phrases{$c},$synonims{$c},0];
			}
			last if($#phraseslist>12);
		}

		return @phraseslist;
	*/
}

func (obj *TextPhrases) finalFilterPhrases(phrases map[string]int, maxcount int) []string {
	// sort phrases by count
	// and get first maxcount real phrases
	phraseslist := []string{}

	for phrase, _ := range phrases {
		ptype := obj.getTypeOfPhrase(phrase)
		if ptype == "n" || ptype == "r" || ptype == "s" || ptype == "f" {
			phraseslist = append(phraseslist, phrase)
		}

		if maxcount > 0 && len(phraseslist) >= maxcount {
			break
		}
	}

	return phraseslist
}

func (obj *TextPhrases) removeCommonPhrases(phrases map[string]int) bool {

	for p, _ := range phrases {
		hasgood := false

		words, _ := obj.splitSentenceForWords(p)

		for _, w := range words {
			if !obj.langobj.IsNotUsefulWord(w) {
				hasgood = true
			}
		}

		if !hasgood {
			delete(phrases, p)
		}
	}
	return true
}

func (obj *TextPhrases) findSinonimPhrases(phrases map[string]int) map[string][]string {
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

			sres := obj.isSubpraseOfPhrase(phrase1, phrase2)

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

func (obj *TextPhrases) findWordsAsPhrases(phrases map[string]int, allwords map[string]int) {
	// add most used words to phrases list
	// check if this word is not used in another word in most cases
	// aim is to find words that are possible company name

	mostappearphrase := float32(helper.GetBiggestValueInMap(phrases)) / 3.0

	for word, count := range allwords {
		if float32(count) > mostappearphrase {
			// check word type
			// NOTE. this is expensive operation
			wtype := obj.langobj.GetTypeOfWord(word, "", "")

			if wtype == "n" || wtype == "r" || wtype == "s" && wtype == "f" {
				phrases[word] = count
			}
		}
	}

}

func (obj *TextPhrases) normalisePhrase(phrase string) string {
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

func (obj *TextPhrases) isSubpraseOfPhrase(phrase1 string, phrase2 string) int8 {
	nphrase1 := obj.normalisePhrase(phrase1)
	nphrase2 := obj.normalisePhrase(phrase2)

	if nphrase1 == nphrase2 {
		if phrase1 == strings.ToUpper(phrase1) ||
			strings.ToUpper(phrase2[0:1]) == phrase2[0:1] {
			return 1
		}
		return -1
	}

	check := obj.langobj.IsPhraseSubphrase(nphrase1, nphrase2)

	if check != 0 {
		return check
	}

	nphrase1 = obj.trimCommonWords(nphrase1, 0)
	nphrase2 = obj.trimCommonWords(nphrase2, 0)

	check = obj.langobj.IsPhraseSubphrase(nphrase1, nphrase2)

	if check != 0 {
		return check
	}

	return 0
}

func (obj *TextPhrases) isSubpraseOfPhraseExtended(phrase1 string, phrase2 string) int8 {
	result := obj.isSubpraseOfPhrase(phrase1, phrase2)

	if result != 0 {
		return result
	}

	if obj.langobj.IsWord(phrase1) && !obj.langobj.IsWord(phrase2) {
		if obj.langobj.IsWordModInPhrase(phrase2, phrase1) {
			return -1
		}
	}

	if obj.langobj.IsWord(phrase2) && !obj.langobj.IsWord(phrase1) {
		if obj.langobj.IsWordModInPhrase(phrase1, phrase2) {
			return 1
		}
	}

	return 0
}

func (obj *TextPhrases) getTypeOfPhrase(phrase string) string {
	alltypes := []string{}

	words, _ := obj.splitSentenceForWords(phrase)

	l := len(words)

	for i, word := range words {
		t := ""

		if obj.langobj.IsNotUsefulWord(word) {
			t = "b"
		} else {
			prevword := ""
			if i > 0 {
				prevword = words[i-1]
			}
			nextword := ""

			if i < l-1 {
				prevword = words[i+1]
			}
			t = obj.langobj.GetTypeOfWord(word, prevword, nextword)
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
