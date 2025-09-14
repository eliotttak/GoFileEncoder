package communFunctions

import (
	"fmt"
	"log"
	"time"
)

func tern[T any](cond bool, ifTrue T, ifFalse T) T {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

func Try(f func() error, attempts int) {
	for {
		err := f()

		if err == nil {
			return
		} else {
			fmt.Println("Tentative échouée : " + err.Error())
			attempts--
		}

		if attempts == 0 {
			log.Fatal("Trop de tentatives échouées")
		}
	}
}

func FormatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	microseconds := int(duration.Microseconds()) % 1000000
	return fmt.Sprintf(
		"%d heure%s, %02d minute%s et %02d.%06d seconde%s",
		hours, tern(hours >= 2, "s", ""),
		minutes, tern(minutes >= 2, "s", ""),
		seconds, microseconds, tern(seconds >= 2, "s", ""))
}
