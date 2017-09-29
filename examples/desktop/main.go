package main

import (
	"fmt"

	"github.com/iafan/Plurr/go/plurr"
	"github.com/iafan/go-l10n/loc"
	"github.com/iafan/go-l10n/locjson"
)

func main() {
	// create localization pool with "en" as a default language
	// and load string resources
	lp := loc.NewPool("en")
	lp.Resources["en"] = locjson.Load("strings/strings.json")
	lp.Resources["ru"] = locjson.Load("strings/strings-ru.json")
	lp.Resources["en-xa"] = locjson.Load("strings/strings-en-xa.json")

	// Print output in English
	printSomeMessages(lp.GetContext("en"))

	// Print output in Russian
	printSomeMessages(lp.GetContext("ru"))

	// Print pseudotranslated output
	printSomeMessages(lp.GetContext("en-xa"))
}

func printSomeMessages(lc *loc.Context) {
	fmt.Println(lc.Tr("Hello"))

	for n := 0; n <= 5; n++ {
		fmt.Println(lc.Format("YouHaveNMessages", plurr.Params{"N": n}))
	}

	fmt.Println("---")
}
