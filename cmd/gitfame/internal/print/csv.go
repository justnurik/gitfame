package print

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CSV(config *PrintConfig) error {
	w := csv.NewWriter(os.Stdout)

	if err := w.Write([]string{"Name", "Lines", "Commits", "Files"}); err != nil {
		return err
	}

	for i := range config.authors {
		record := []string{
			config.authors[i],
			fmt.Sprintf("%d", config.linesCount[i]),
			fmt.Sprintf("%d", config.commitsCount[i]),
			fmt.Sprintf("%d", config.filesCount[i]),
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}

	w.Flush()
	return w.Error()
}
