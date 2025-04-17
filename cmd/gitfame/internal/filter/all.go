package filter

import (
	_ "embed"

	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/load"
)

func All(files, extensions, languages, exclude, restrictTo []string) ([]string, error) {
	var err error

	if extensions != nil {
		files = ByExtensions(files, extensions)
	}

	if exclude != nil {
		files, err = ByGlob(files, exclude, true)
		if err != nil {
			return nil, err
		}
	}

	if restrictTo != nil {
		files, err = ByGlob(files, restrictTo, false)
		if err != nil {
			return nil, err
		}
	}

	if languages != nil {
		allLanguages, err := load.WaitLoadLanguageExtensions()
		if err != nil {
			return nil, err
		}

		files = ByLanguages(files, languages, allLanguages)
	}

	return files, err
}
