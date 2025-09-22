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
		FileEncodedIn string
	}

	decodingTranslationsType struct {
		FileDecodedIn string
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
				EncodedBinFiles:        "Encoded binary files (.enc.bin)",
				SaveFile:               "Save a file",
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
				FileEncodedIn: "File encoded in %s.\n",
			},
			Decoding: decodingTranslationsType{
				FileDecodedIn: "File decoded in %s.\n",
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
