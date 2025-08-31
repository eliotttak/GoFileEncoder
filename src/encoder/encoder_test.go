package encoder

import (
	"testing"
	"time"
)

func Test_formatDuration(t *testing.T) {
	s := formatDuration(time.Hour + 3*time.Minute + 14*time.Second + 1524*time.Millisecond)
	if s != "1 heure(s), 03 minute(s) et 15.524 seconde(s)" {
		t.Error("Test failed with " + s)
	}
}
