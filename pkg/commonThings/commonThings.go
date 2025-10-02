// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commonThings

import (
	"errors"
	"fmt"
	"log"
	"path"
	"time"

	"github.com/christianhujer/isheadless"
	"github.com/eliotttak/GoFileEncoder/pkg/translate"
	"github.com/kjk/goey/dialog"
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

	translations := translate.GetTranslations()

	if isheadless.IsHeadless() {
		fmt.Print(translations.CommonThings.NoGUI)
		var path string
		fmt.Scanln(&path)
		return path, nil
	} else {
		switch actionType {
		case Load:
			var dialogBox *dialog.OpenFile = dialog.NewOpenFile()

			if title != "" {
				dialogBox = dialogBox.WithTitle(title)
			}

			for _, filter := range filters {
				dialogBox = dialogBox.AddFilter(filter[0], filter[1])
			}

			if startFile != "" || startDir != "" {
				dialogBox = dialogBox.WithFilename(path.Join(startDir, startFile))
			}

			result, err := dialogBox.Show()

			if result == "" {
				err = errors.New("no file selected")
			}

			return result, err
		case Save:
			var dialogBox *dialog.SaveFile = dialog.NewSaveFile()

			if title != "" {
				dialogBox = dialogBox.WithTitle(title)
			}

			for _, filter := range filters {
				dialogBox = dialogBox.AddFilter(filter[0], filter[1])
			}

			if startFile != "" || startDir != "" {
				dialogBox = dialogBox.WithFilename(path.Join(startDir, startFile))
			}

			result, err := dialogBox.Show()

			if result == "" {
				err = errors.New("no file selected")
			}

			return result, err
		default:
			return "", fmt.Errorf("'actionType' must be either \"save\" ('Save') or \"load\" ('Load'), so \"%s\" is incorrect", actionType)
		}
	}
}

func Tern[T any](cond bool, ifTrue T, ifFalse T) T {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

func Try(f func() error, attempts int) {
	translations := translate.GetTranslations()

	for {
		err := f()

		if err == nil {
			return
		} else {
			fmt.Printf(translations.CommonThings.EchoedAttempt, err.Error())
			attempts--
		}

		if attempts == 0 {
			log.Fatal(translations.CommonThings.ToManyEchoedAttempts)
		}
	}
}

func FormatDuration(duration time.Duration) string {
	translations := translate.GetTranslations()

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	microseconds := int(duration.Microseconds()) % 1000000
	return fmt.Sprintf(
		translations.CommonThings.Time,
		hours, Tern(hours >= 2, translations.CommonThings.Hours, translations.CommonThings.Hour),
		minutes, Tern(minutes >= 2, translations.CommonThings.Minutes, translations.CommonThings.Minute),
		seconds, microseconds, Tern(seconds >= 2, translations.CommonThings.Seconds, translations.CommonThings.Second))
}
