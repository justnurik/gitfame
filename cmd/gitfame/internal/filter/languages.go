package filter

import (
	"fmt"
	"os"
	"strings"
)

func ByLanguages(files, languages []string, allLanguages map[string][]string) []string {
	var allowedExtensions []string

	for _, language := range languages {
		langLower := strings.ToLower(language)
		extensions, ok := allLanguages[langLower]

		if !ok {
			fmt.Fprintf(os.Stderr, "warning: неизвестный язык '%s'\n", language)
			continue
		}

		allowedExtensions = append(allowedExtensions, extensions...)
	}

	return ByExtensions(files, allowedExtensions)
}
