package locweb

import (
	"net/http"

	"github.com/iafan/go-l10n/loc"
	"golang.org/x/text/language"
)

// GetBestLanguageFromHTTP returns the best language match
// based on the provided list of language tags and the value of
// "Accept-Language" request header and, optionally, explicit language name
// defined in a cookie or a parameter
func GetBestLanguageFromHTTP(lp *loc.Pool, r *http.Request,
	cookieName, paramName string) string {

	if paramName != "" {
		lang := r.URL.Query().Get(paramName)
		if lang != "" && lp.Resources[lang] != nil {
			return lang
		}
	}

	if cookieName != "" {
		cookie, err := r.Cookie(cookieName)
		if err == nil && lp.Resources[cookie.Value] != nil {
			return cookie.Value
		}
	}

	tags := make([]language.Tag, len(lp.Resources))
	i := 0
	for lang := range lp.Resources {
		tags[i] = language.MustParse(lang)
		i++
	}
	matcher := language.NewMatcher(tags)

	tags, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	if err != nil {
		return lp.DefaultLanguage
	}
	tag, _, _ := matcher.Match(tags...)
	return tag.String()
}
