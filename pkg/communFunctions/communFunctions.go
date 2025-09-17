package communFunctions

import (
	"fmt"
	"log"
	"time"

	"github.com/christianhujer/isheadless"
	"github.com/sqweek/dialog"
)

type (
	SelectFilePathFilter  [2]string
	SelectFilePathFilters []SelectFilePathFilter
)

const (
	Save string = "save"
	Load string = "load"
)

func SelectFilePath(
	title string,
	filters SelectFilePathFilters,
	startFile string,
	startDir string,
	actionType string,
) (string, error) {

	if isheadless.IsHeadless() {
		fmt.Print("Aucune interface graphique n'a été détectée sur votre ordinateur. Entrez le chemin absolu vers votre fichier : ")
		var path string
		fmt.Scanln(&path)
		return path, nil
	} else {
		var popup *dialog.FileBuilder = dialog.File()

		if title != "" {
			popup = popup.Title(title)
		}

		for _, filter := range filters {
			popup = popup.Filter(filter[0], filter[1])
		}

		if startFile != "" {
			popup = popup.SetStartFile(startFile)
		}

		if startDir != "" {
			popup = popup.SetStartDir(startDir)
		}

		switch actionType {
		case Save:
			return popup.Save()
		case Load:
			return popup.Load()
		default:
			return "", fmt.Errorf("'actionType' must be either \"save\" ('Save') or \"load\" ('Load'), so \"%s\" is incorrect", actionType)
		}
	}
}

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
