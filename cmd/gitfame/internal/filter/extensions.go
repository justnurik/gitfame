package filter

import (
	"path"
)

func ByExtensions(files, extensions []string) []string {
	setExtensions := make(map[string]struct{})

	for _, extension := range extensions {
		setExtensions[extension] = struct{}{}
	}

	var filteredFiles []string

	for _, file := range files {
		if _, ok := setExtensions[path.Ext(file)]; ok {

			filteredFiles = append(filteredFiles, file)
		}
	}

	return filteredFiles
}
