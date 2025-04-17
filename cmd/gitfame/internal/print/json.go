package print

import (
	"encoding/json"
	"os"
)

func JSON(config *PrintConfig) error {
	data := make([]map[string]any, len(config.authors))

	for i := range config.authors {
		data[i] = map[string]any{
			"name":    config.authors[i],
			"lines":   config.linesCount[i],
			"commits": config.commitsCount[i],
			"files":   config.filesCount[i],
		}
	}

	encoder := json.NewEncoder(os.Stdout)
	return encoder.Encode(data)
}
