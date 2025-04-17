package programm

import (
	"fmt"
	"strings"
)

type OrderType int

const (
	Line OrderType = iota
	Commits
	Files
)

func (o *OrderType) Type() string {
	return "[lines|commits|files]"
}

func (o *OrderType) Set(value string) error {
	switch strings.ToLower(value) {
	case "lines":
		*o = Line
	case "commits":
		*o = Commits
	case "files":
		*o = Files
	default:
		return fmt.Errorf("%s", value)
	}

	return nil
}

func (o *OrderType) String() string {
	switch *o {
	case Line:
		return "lines"
	case Commits:
		return "commits"
	case Files:
		return "files"
	default:
		return "unknown"
	}
}
