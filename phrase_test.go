package keyphrases

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPhrases(t *testing.T) {
	analyser := getAnalyserForTesting()

	text := `Whenever Mario Draghi clears a hurdle on his path to higher inflation, a new one appears.
Just as the 19-nation economy sends encouraging signals that challenges from Brexit to terrorism won’t derail the modest recovery, a new decline in oil prices is casting a shadow over an expected pick-up in inflation. With growth not strong enough to generate price pressures, the European Central Bank president may have to revise his outlook yet again.
Inflation remains far below the ECB’s 2 percent goal after more than two years of unprecedented stimulus and isn’t seen reaching it before 2018. Staff will begin to draw up fresh forecasts in mid-August, and while officials are in no rush to adjust or expand their 1.7 trillion-euro ($1.9 trillion) quantitative-easing plan in September, economists predict Draghi will have to ease policy before the end of the year.
“Now that the euro-area economy seems to have shrugged off the Brexit vote, focus will again shift on inflation, against the background of those negative news from oil prices,” said Johannes Gareis, an economist at Natixis in Frankfurt. “Yes, the ECB has managed to dispel deflation fears, but all the uncertainty means inflation will stay lower for longer -- and Draghi will have to take notice.”`

	sentences, _ := analyser.splitTextForSentences(text)

	words, _ := analyser.splitSentencesForWords(sentences)

	phrases, _ := analyser.getBasicPhrasesHash(sentences, words)
	fmt.Println(phrases)

	expected := map[string]int{
		"new one appears":         2,
		"revise his outlook":      2,
		"Brexit oil prices":       2,
		"take notice":             3,
		"economy the Brexit":      2,
		"economy inflation":       2,
		"Brexit inflation":        2,
		"oil prices":              6,
		"economy from oil prices": 2,
	}

	if !reflect.DeepEqual(phrases, expected) {
		t.Fatalf("Got %s, expected %s.", phrases, expected)
	}

}
