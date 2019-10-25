package clock

import "strings"

//StoppedAt returns the number which the seconds needle stopped
//input = string of letter "a" for right button and/or "i" for left button
//ex. StoppedAt("aa") => 1
//    StoppedAt("aaii") => 60
//    StoppedAt("aiia") => 59
func StoppedAt(buttonsClicked string) int {
	//to get the individual char from string
	buttons := strings.Split(buttonsClicked, "")

	counter := 0
	prevButton := "b"

	//for every button pressed
	for _, button := range buttons {
		//if the button didn't get pressed more than once the ignore it
		if button != prevButton {
			prevButton = button
			continue
		}
		//if "a" turn clockwise, else the other way around
		switch button {
		case "a":
			counter++
			break
		case "i":
			counter--
			break
		}
	}

	//if it turns counter clockwise add 60 to get the actual position
	if counter < 1 {
		counter += 60
	}
	return counter
}
