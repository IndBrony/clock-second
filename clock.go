package clock

import "strings"

func StoppedAt(buttonsClicked string) int {
	buttons := strings.Split(buttonsClicked, "")

	counter := 0
	prevButton := "b"

	for _, button := range buttons {
		if button != prevButton {
			prevButton = button
			continue
		}
		switch button {
		case "a":
			counter++
			break
		case "i":
			counter--
			break
		}
	}
	if counter < 1 {
		counter += 60
	}
	return counter
}
