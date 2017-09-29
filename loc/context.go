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
