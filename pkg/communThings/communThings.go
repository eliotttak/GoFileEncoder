package communThings

import (
	"fmt"
	"log"
	"time"

	"github.com/christianhujer/isheadless"
	"github.com/eliotttak/GoFileEncoder/pkg/translations"
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
		fmt.Print(translations.GetTranslations().NoGUI)
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
			fmt.Printf(translations.GetTranslations().EchoedAttempt, err.Error())
			attempts--
		}

		if attempts == 0 {
			log.Fatal(translations.GetTranslations().ToManyEchoedAttempts)
		}
	}
}

func FormatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	microseconds := int(duration.Microseconds()) % 1000000
	return fmt.Sprintf(
		translations.GetTranslations().Time,
		hours, tern(hours >= 2, translations.GetTranslations().Hours, translations.GetTranslations().Hour),
		minutes, tern(minutes >= 2, translations.GetTranslations().Minutes, translations.GetTranslations().Minute),
		seconds, microseconds, tern(seconds >= 2, translations.GetTranslations().Seconds, translations.GetTranslations().Second))
}
