package print

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func Tabular(config *PrintConfig) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 1, ' ', 0)

	fmt.Fprintln(w, "Name\tLines\tCommits\tFiles")

	for i, author := range config.authors {
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\n",
			author,
			config.linesCount[i],
			config.commitsCount[i],
			config.filesCount[i])
	}

	return w.Flush()
}
