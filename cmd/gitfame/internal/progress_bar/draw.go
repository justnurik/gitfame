package progressbar

import "fmt"

// todo: package progressbar
func draw(progress, total int) {
	const barWidth = 100 //! TODO

	percent := float64(progress) / float64(total)
	filled := int(barWidth * percent)

	fmt.Printf("\r[")
	for i := 0; i < barWidth; i++ {
		if i < filled {
			fmt.Print("=")
		} else if i == filled {
			fmt.Print(">")
		} else {
			fmt.Print(" ")
		}
	}

	fmt.Printf("] %3.0f%%", percent*100)
}
