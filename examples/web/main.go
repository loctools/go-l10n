package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"

	"github.com/iafan/go-l10n/loc"
	"github.com/iafan/go-l10n/lochtml"
	"github.com/iafan/go-l10n/locweb"
)

var locpool = loc.NewPool("en")

func indexPage(w http.ResponseWriter, r *http.Request) {
	lang := locweb.GetBestLanguageFromHTTP(locpool, r, "", "lang")
	lc := locpool.GetContext(lang)

	tpl, err := lochtml.Template(lc, template.New("")).ParseFiles("templates/hello.html")
	if err != nil {
		panic(err)
	}

	tpl.ExecuteTemplate(w, "hello.html", map[string]interface{}{
		"NAME":    "John Doe",
		"X":       rand.Int31n(100) + 5, // number of files
		"Y":       rand.Int31n(5) + 1,   // number of folders
		"COMMAND": rand.Int31n(3),       // command index: 0=copy, 1=move, 2=delete
	})
}

func main() {
	http.HandleFunc("/", indexPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
