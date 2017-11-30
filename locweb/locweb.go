package locweb

import (
	"net/http"
	"sort"

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

	header := r.Header.Get("Accept-Language")

	// If the header is empty, just return the default language.
	if header == "" {
		return lp.DefaultLanguage
	}

	// Otherwise, use language matcher to determine the best language.
	// Sort language tags alphabetically, but place the default language
	// to the beginning of the list, since the ordering of tags
	// is important for the matcher.
	langNames := make([]string, len(lp.Resources))
	i := 0
	for lang := range lp.Resources {
		if lang == lp.DefaultLanguage {
			langNames[i] = "" // use an empty string to make it the first one
		} else {
			langNames[i] = lang
		}
		i++
	}
	sort.Strings(langNames)
	if langNames[0] == "" {
		langNames[0] = lp.DefaultLanguage // restore the actual language value
	}

	tags := make([]language.Tag, len(langNames))
	for i := range langNames {
		tags[i] = language.MustParse(langNames[i])
	}
	matcher := language.NewMatcher(tags)

	tags, _, err := language.ParseAcceptLanguage(header)
	if err != nil {
		return lp.DefaultLanguage
	}
	tag, _, _ := matcher.Match(tags...)
	return tag.String()
}
