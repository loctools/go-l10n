// Package locjson provides the `LoadTranslations()` function to
// load translations from JSON files of the following structure:
//
// {
//     "keyName": {
//         "message": "String to translate",
//         "description": "Description for translators"
//     },
//     ...
// }
package locjson

import (
	"encoding/json"
	"io/ioutil"
)

type translationFile map[string]struct {
	Message     string
	Description string
}

// Load takes the JSON file name and returns string resources
// in a format compatible with loc.Resources map values.
// It will panic if there's a problem with loading/unmarshaling JSON.
func Load(filename string) map[string]string {
	t := make(map[string]string)
	var tf translationFile

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &tf)
	if err != nil {
		panic(err)
	}

	// convert JSON data structure into a destination map format
	for k, v := range tf {
		t[k] = v.Message
	}
	return t
}
