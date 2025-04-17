package print

import (
	"sort"

	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/programm"
)

type PrintConfig struct {
	authors      []string
	linesCount   []int
	filesCount   []int
	commitsCount []int
}

func cmpSliceGreatEqual(a, b []int) bool {
	for i := range a {
		if a[i] == b[i] {
			continue
		}

		return a[i] > b[i]
	}

	return false
}

func cmpSliceEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func NewConfig(orderType programm.OrderType, authToLineCount, authToCommitCount, authToFileCount map[string]int) *PrintConfig {
	authors := make([]string, 0, len(authToLineCount))
	for author := range authToLineCount {
		authors = append(authors, author)
	}

	comp := func(authTo1 map[string]int, authTo2 map[string]int, authTo3 map[string]int) func(int, int) bool {
		return func(i, j int) bool {
			keyI := []int{authTo1[authors[i]], authTo2[authors[i]], authTo3[authors[i]]}
			keyJ := []int{authTo1[authors[j]], authTo2[authors[j]], authTo3[authors[j]]}

			if cmpSliceEqual(keyI, keyJ) {
				return authors[i] < authors[j]
			}

			return cmpSliceGreatEqual(keyI, keyJ) // >=
		}
	}

	switch orderType {
	case programm.Line:
		sort.Slice(authors, comp(authToLineCount, authToCommitCount, authToFileCount))

	case programm.Files:
		sort.Slice(authors, comp(authToFileCount, authToLineCount, authToCommitCount))

	case programm.Commits:
		sort.Slice(authors, comp(authToCommitCount, authToLineCount, authToFileCount))
	}

	config := PrintConfig{
		authors:      authors,
		linesCount:   make([]int, len(authors)),
		filesCount:   make([]int, len(authors)),
		commitsCount: make([]int, len(authors)),
	}

	for i, author := range authors {
		config.linesCount[i] = authToLineCount[author]
		config.commitsCount[i] = authToCommitCount[author]
		config.filesCount[i] = authToFileCount[author]
	}

	return &config
}
