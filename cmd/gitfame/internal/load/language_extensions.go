package load

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed language_extensions.json
var languageExtensions []byte

type language struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Extensions []string `json:"extensions"`
}

func loadLanguageExtensions(languageExtensions []byte) (map[string][]string, error) {
	allLanguages := make(map[string][]string)
	var languages []language

	if err := json.Unmarshal(languageExtensions, &languages); err != nil {
		return nil, fmt.Errorf("не удалось разобрать JSON: %v", err)
	}

	for _, language := range languages {
		allLanguages[strings.ToLower(language.Name)] = language.Extensions
	}

	return allLanguages, nil
}

var allLanguages map[string][]string
var chanErr chan error

func StartLoadLanguageExtensions() {
	chanErr = make(chan error, 1)

	go func() {
		var err error
		allLanguages, err = loadLanguageExtensions(languageExtensions)
		chanErr <- err
	}()
}

func WaitLoadLanguageExtensions() (map[string][]string, error) {
	err := <-chanErr
	close(chanErr)

	if err != nil {
		return nil, fmt.Errorf("не удалось скачать language_extensions.json: %v", err)
	}

	return allLanguages, nil
}
