package lochtml

import (
	"html/template"

	"github.com/iafan/go-l10n/loc"
)

// Template takes localization Context and an HTML template and returns
// an enhanced template with injected translation functions
func Template(lc *loc.Context, tpl *template.Template) *template.Template {
	return tpl.Funcs(template.FuncMap{
		"tr": lc.Tr,
		"tf": lc.Format,
	})
}
