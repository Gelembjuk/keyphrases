## KeyPhrases with Go 

Golang package to extract key phrases from a text

A function analyses a given text and returns a list of key phrases. This version works with English only.
However, there is a way to extend for other languages

### Installation

go get github.com/gelembjuk/keyphrases

### Example 

```
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gelembjuk/keyphrases"
)

func main() {
	textfile := "inputtext.txt"
	filecontents, _ := ioutil.ReadFile(textfile)

	text := string(filecontents)

	analyser := keyphrases.TextPhrases{Language: "english"}

	analyser.Init()

	phrases := analyser.GetKeyPhrases(text)

	for _, phrase := range phrases {
		fmt.Println(phrase)
	}
}
```

### Author

Roman Gelembjuk (@gelembjuk)

LinkedIn: [https://linkedin.com/in/gelembjuk](https://linkedin.com/in/gelembjuk)