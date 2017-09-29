About go-l10n
=============

Lightweight yet powerful continuous localization solution for Go, based on
[Serge](https://serge.io/) and [Plurr](https://github.com/iafan/Plurr).

Advantages
----------

 1. Provides solution for both web-based and desktop applications.
 2. Can be used with different resource file formats, thanks to Serge.
    Currently offers the support for JSON resource files and
    string maps in the native .go format. Both formats support
    developer comments for translators to get extra context.
 3. Supports named placeholders, plurals and conditionals via
    Plurr-formatted strings (see the
    [interactive demo](http://iafan.github.io/plurr-demo/)).
 4. Support for translation and formatting functions in `html/template`
    and `text/template`.
 5. For web applications, an ability to auto-detect the best language
    from the browser, and to force the language using a cookie or
    a query string parameter.
 6. Comes with ready-to-use Serge configuration files to allow for
    seamless continuous localization process (see
    [examples](https://github.com/iafan/go-l10n/tree/master/examples/)).
 7. Supports pseudotranslations out of the box. Pseudotranslation is a way
    to "translate" your application by applying some algorithm to each
    source string; it is useful to test if your application looks good
    with other locales without waiting for actual translations.

 7. The code is split into small packages to minimize external dependencies.

Example Code
------------

There are two example projects that showcase approaches to localizing
[web applications](https://github.com/iafan/go-l10n/tree/master/examples/web/) and
[desktop applications](https://github.com/iafan/go-l10n/tree/master/examples/desktop/).
Below is a condensed example to illustrate the typical usage.

`strings.go` (master resource file):
```go
package main

func init() {
    locpool.Resources["en"] = map[string]string{
        // Page title
        "Hello": "Hello!",

        // {N} is the number of messages
        "YouHaveNMessages": "You have {N} {N_PLURAL:message|messages}",
    }
}
```

`strings-ru.go` (localized resource file):
```go
package main

func init() {
    locpool.Resources["ru"] = map[string]string{
        // Page title
        "Hello": "Здравствуйте!",

        // {N} is the number of messages
        "YouHaveNMessages": "У вас {N} {N_PLURAL:сообщение|сообщения|сообщений}",
    }
}
```

Code:
```go
package main

import (
    "github.com/iafan/Plurr/go/plurr"
    "github.com/iafan/go-l10n/loc"
)

// Create a global localization pool which will be populated
// by resource files; use English as a default (fallback) language
var locpool = loc.NewPool("en")

func main() {
  // Get Russian localization context
  lc := locpool.GetContext("ru")

  // Get translation by key name:
  name := lc.Tr("Hello")

  // get translation by key name, then format it using Plurr:
  hello := lc.Format("YouHaveNMessages", plurr.Params{"N": 5})

  ...
}
```
