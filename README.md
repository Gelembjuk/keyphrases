## KeyPhrases with Go 

Golang package to extract key phrases from a text.

A function analyses a given text and returns a list of key phrases. This version works with English only.
However, there is a way to extend for other languages.

The package uses the [WordNet](https://wordnet.princeton.edu/) dictionary to work with english texts. No need to install the WordNet. You need only "dict/" folder and set a path to it when create analiser object. 

### Installation

go get github.com/gelembjuk/keyphrases

### Example 1

Get key phrases from a text saved in a file .

```go
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

### Example 2

Get key phrases from a web page. We use the package github.com/gelembjuk/articletext to gextract a text from a web page

```go
package main

import (
	"fmt"
	"os"

	"github.com/gelembjuk/articletext"
	"github.com/gelembjuk/keyphrases"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No URL provided")
		os.Exit(1)
	}
	// get URL from command line argument
	url := os.Args[1]
	
	// get text from this web page
	text, err := articletext.GetArticleTextFromUrl(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	// print a text to a console
	fmt.Println(text)
	
	// Create a text analyser object. It requires a path to WordNet dictionary directory
	 
	analyser := keyphrases.TextPhrases{Language: "english",
		LanguageOptions: map[string]string{"wordnetdirectory": "/home/roman/Projects/Go/WordNet/dict"}}

	// this is required procedure to initialise analyser 
	analyser.Init()

	// get key phrases
	phrases := analyser.GetKeyPhrases(text)

	for _, phrase := range phrases {
		fmt.Println(phrase)
	}
}
```

### Author

Roman Gelembjuk (@gelembjuk)

LinkedIn: [https://linkedin.com/in/gelembjuk](https://linkedin.com/in/gelembjuk)