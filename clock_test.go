package clock

import "testing"

func TemplateClockTest(t *testing.T, input string, expectedOutput int) {
	if output := StoppedAt(input); output != expectedOutput {
		t.Errorf("Test Failed on input %s; Expecting %v but got %v", input, expectedOutput, output)
	}
}

func TestClock(t *testing.T) {
	TemplateClockTest(t, "ai", 60)
	TemplateClockTest(t, "aa", 1)
	TemplateClockTest(t, "aaa", 2)
	TemplateClockTest(t, "aai", 1)
	TemplateClockTest(t, "aaii", 60)

	TemplateClockTest(t, "ii", 59)
	TemplateClockTest(t, "iii", 58)
	TemplateClockTest(t, "iia", 59)
}
