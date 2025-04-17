package print

import (
	"encoding/json"
	"os"
)

func JSONLines(config *PrintConfig) error {
	encoder := json.NewEncoder(os.Stdout)

	for i := range config.authors {
		entry := map[string]any{
			"name":    config.authors[i],
			"lines":   config.linesCount[i],
			"commits": config.commitsCount[i],
			"files":   config.filesCount[i],
		}

		if err := encoder.Encode(entry); err != nil {
			return err
		}
	}

	return nil
}
