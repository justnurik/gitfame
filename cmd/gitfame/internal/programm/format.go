package programm

import (
	"fmt"
	"strings"
)

type FormatType int

const (
	Tabular FormatType = iota
	CSV
	JSON
	JSONLines
)

func (f *FormatType) Type() string {
	return "[tabular|csv|json|json-lines]"
}

func (f *FormatType) Set(value string) error {
	switch strings.ToLower(value) {
	case "tabular":
		*f = Tabular
	case "csv":
		*f = CSV
	case "json":
		*f = JSON
	case "json-lines":
		*f = JSONLines
	default:
		return fmt.Errorf("%s", value)
	}
	return nil
}

func (f FormatType) String() string {
	switch f {
	case Tabular:
		return "tabular"
	case CSV:
		return "csv"
	case JSON:
		return "json"
	case JSONLines:
		return "json-lines"
	default:
		return "unknown"
	}
}
