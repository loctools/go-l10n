package loc

import (
	"github.com/iafan/Plurr/go/plurr"
)

// Resources is an array that holds string resources for a particular language
type Resources map[string]string

// Pool object holds all string resources and language-specific contexts.
// Pool can be created either for an entire project, or for each logical
// part of the application that has its own set of strings.
type Pool struct {
	DefaultLanguage string
	Resources       map[string]Resources
	contexts        map[string]*Context
}

// NewPool creates a new localization pool object
func NewPool(defaultLanguage string) *Pool {
	return &Pool{
		DefaultLanguage: defaultLanguage,
		Resources:       make(map[string]Resources),
		contexts:        make(map[string]*Context),
	}
}

// GetContext returns a translation context based on the provided language.
// Contexts are thread-safe.
func (lp *Pool) GetContext(lang string) *Context {
	if lp.contexts[lang] == nil {
		lp.contexts[lang] = &Context{
			pool:  lp,
			lang:  lang,
			plurr: plurr.New().SetLocale(lang),
		}
	}

	return lp.contexts[lang]
}
