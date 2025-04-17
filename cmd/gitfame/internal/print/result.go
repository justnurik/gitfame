package print

import (
	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/programm"
)

func Result(formatType programm.FormatType, config *PrintConfig) error {
	switch formatType {
	case programm.Tabular:
		return Tabular(config)
	case programm.CSV:
		return CSV(config)
	case programm.JSON:
		return JSON(config)
	default:
		return JSONLines(config)
	}
}
