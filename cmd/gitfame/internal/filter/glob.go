package filter

import (
	"fmt"
	"path/filepath"
)

func ByGlob(files, patterns []string, exclude bool) ([]string, error) {
	var filteredFiles []string

	for _, file := range files {
		oneMatch := false

		for _, pattern := range patterns {
			match, err := filepath.Match(pattern, file)
			if err != nil {
				return nil, fmt.Errorf("некорректный Glob-паттерн: %s: %v", pattern, err)
			}

			if match {
				oneMatch = true
				break
			}
		}

		if (!oneMatch && exclude) || (oneMatch && !exclude) {
			filteredFiles = append(filteredFiles, file)
		}

	}

	return filteredFiles, nil
}
