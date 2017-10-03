package loc

import (
	"github.com/iafan/Plurr/go/plurr"
)

// Context holds a localization context
// (a combination of a language, a corresponding Plurr formatter,
// and a link back to parent localization Pool object)
type Context struct {
	pool  *Pool
	lang  string
	plurr *plurr.Plurr
}

// GetLanguage returns the current language set for the context.
func (lc *Context) GetLanguage() string {
	return lc.lang
}

// GetContext returns a context for another language
// withing the same localization pool. This can be useful to
// switch the language within the same pool (think of it
// as of a `SetLanguage()` function which returns a new context
// without changing the previous one)
func (lc *Context) GetContext(lang string) *Context {
	return lc.pool.GetContext(lang)
}

// Tr returns a translated version of the string
func (lc *Context) Tr(key string) string {
	// try preferred language
	s, ok := lc.pool.Resources[lc.lang][key]
	if ok {
		return s
	}

	// try default language
	s, ok = lc.pool.Resources[lc.pool.DefaultLanguage][key]
	if ok {
		return s
	}

	// if everything else failed, return the key as is
	return key
}

// Format returns a formatted version of the string
func (lc *Context) Format(key string, params plurr.Params) string {
	s, err := lc.plurr.Format(lc.Tr(key), params)
	if err != nil {
		return key
	}

	return s
}
