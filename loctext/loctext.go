package loctext

import (
	"text/template"

	"github.com/iafan/go-l10n/loc"
)

// Template takes localization Context and a text template and returns
// an enhanced template with injected translation functions
func Template(lc *loc.Context, tpl *template.Template) *template.Template {
	return tpl.Funcs(template.FuncMap{
		"tr": lc.Tr,
		"tf": lc.Format,
	})
}
