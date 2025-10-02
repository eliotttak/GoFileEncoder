// Copyright 2025 Eliott Takvorian
//
// This file is part of GoFileEncoder.
//
// GoFileEncoder is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.
package translate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type (
	generalTranslationsType struct {
		EncodedBinFiles        string
		SelectFile             string
		SaveFile               string
		PressEnterToSelectFile string
		YouSelectedFile        string
		EnterPassword          string
		PressEnterToSaveFile   string
		AllFiles               string
	}
	commonThingsTranslationsType struct {
		NoGUI                string
		EchoedAttempt        string
		ToManyEchoedAttempts string
		Hour                 string
		Hours                string
		Minute               string
		Minutes              string
		Second               string
		Seconds              string
		Time                 string
	}
	introTranslationsType struct {
		WhatWouldYouLikeToDo string
		EncodeInitial        string
		DecodeInitial        string
		LicenseInitial       string
	}

	encodingTranslationsType struct {
		YesInitial          string
		NoInitial           string
		IsPwdOk             string
		CrackedIn           string
		MoreThan290         string
		PasswordValidated   string
		FileEncodedIn       string
		WhichEncodingMethod string
	}

	decodingTranslationsType struct {
		FileDecodedIn           string
		WhichEncodingMethodUsed string
	}

	translationsType struct {
		General      generalTranslationsType
		CommonThings commonThingsTranslationsType
		Intro        introTranslationsType
		Encoding     encodingTranslationsType
		Decoding     decodingTranslationsType
	}
)

func tern[T any](cond bool, ifTrue T, ifFalse T) T {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

func extractLangPart(filename string) string {
	reg := regexp.MustCompile("[a-z]{2}-[A-Z]{2}")

	return reg.FindString(filename)
}

var translations translationsType = translationsType{}

func getExecutableDir() string {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	execName := filepath.Base(execPath)

	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		log.Fatal(err)
	}

	execDir, _ := strings.CutSuffix(execPath, execName)

	return execDir
}

func listTranslationFiles() []string {
	var translationFiles []string

	dir := getExecutableDir()
	pattern := `^translate-[a-z]{2}-[A-Z]{2}\.json$`
	re := regexp.MustCompile(pattern)

	allFiles, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range allFiles {
		if !file.IsDir() && re.MatchString(file.Name()) {
			translationFiles = append(translationFiles, filepath.Join(dir, file.Name()))
		}
	}

	return translationFiles
}

func GetTranslations() translationsType {
	if translations != (translationsType{}) {
		return translations
	} else {
		var englishTranslations translationsType = translationsType{
			General: generalTranslationsType{
				EncodedBinFiles:        "Encoded binary files",
				SaveFile:               "Save a file",
				SelectFile:             "Select a file",
				PressEnterToSelectFile: "Press [Enter] to select a file...",
				YouSelectedFile:        "You selected this file: %s.\n\n",
				EnterPassword:          "Enter the password: ",
				PressEnterToSaveFile:   "Press [Enter] to save the file...",
				AllFiles:               "All files",
			},
			CommonThings: commonThingsTranslationsType{
				NoGUI:                "No graphical interface has been detected on your computer. Enter the absolute path to your file: ",
				EchoedAttempt:        "Echoed attempt: %s\n",
				ToManyEchoedAttempts: "Too many echoed attempts",
				Hour:                 "hour",
				Hours:                "hours",
				Minute:               "minute",
				Minutes:              "minutes",
				Second:               "second",
				Seconds:              "seconds",
				Time:                 "%d %s, %02d %s and %02d.%06d %s",
			},
			Intro: introTranslationsType{
				WhatWouldYouLikeToDo: "What would you like to do?\n - Encode a file (e)\n - Decode a file (d)\n - Read the license (l)\n",
				EncodeInitial:        "e",
				DecodeInitial:        "d",
				LicenseInitial:       "l",
			},
			Encoding: encodingTranslationsType{
				YesInitial:          "y",
				NoInitial:           "n",
				IsPwdOk:             "Is it right for you?",
				CrackedIn:           "The password would be cracked in %s\n",
				MoreThan290:         "more than 90 years.",
				PasswordValidated:   "The password has been validated.",
				FileEncodedIn:       "File encoded in %s.\n",
				WhichEncodingMethod: "Which encoding method would you like to use?\n - Only the XOR encryption (x)\n - Both XOR and Rotate encryption (recommended) (xr)\n",
			},
			Decoding: decodingTranslationsType{
				FileDecodedIn:           "File decoded in %s.\n",
				WhichEncodingMethodUsed: "Which encryption method has been used to encode this file?\n - Only the XOR encryption (x)\n - Both XOR and Rotate encryption (xr)\n",
			},
		}

		translationFiles := listTranslationFiles()

		switch len(translationFiles) {
		case 0:
			translations = englishTranslations
			return translations
		case 1:
			translationFile, err := os.Open(translationFiles[0])
			defer translationFile.Close()

			if err != nil {
				log.Fatal(err)
			}

			var jsonFile []byte

			jsonFile, err = io.ReadAll(translationFile)
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(jsonFile, &translations)
			if err != nil {
				log.Fatal(err)
			}
			return translations
		default:
			var availableLanguages []string

			for _, fileName := range translationFiles {
				availableLanguages = append(availableLanguages, extractLangPart(fileName))
			}

			fmt.Println("There are several available languages; please choose one.")

			for i, language := range availableLanguages {
				fmt.Printf(" - %s (%d)\n", language, i+1)
			}

			var rep int
			for !(rep >= 1 && rep <= len(availableLanguages)) {
				fmt.Printf("(1/%s%d)>>>", tern(len(availableLanguages) >= 3, ".../", ""), len(availableLanguages))
				var sRep string
				fmt.Scanln(&sRep)

				rep, _ = strconv.Atoi(sRep)
				time.Sleep(300 * time.Millisecond)
			}

			usedFileName := translationFiles[rep-1]

			var usedFile *os.File

			usedFile, err := os.Open(usedFileName)

			if err != nil {
				log.Fatal(err)
			}

			var jsonFile []byte

			jsonFile, err = io.ReadAll(usedFile)
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(jsonFile, &translations)
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(300 * time.Millisecond)

			return translations
		}

	}
}
