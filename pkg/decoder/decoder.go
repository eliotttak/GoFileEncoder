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
package decoder

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/eliotttak/GoFileEncoder/pkg/commonThings"
	"github.com/eliotttak/GoFileEncoder/pkg/translate"

	"golang.org/x/term"
)

func removeExtentions(path string, nExtensions int) string {
	var result string = path
	for range nExtensions {
		var ext string = filepath.Ext(result)
		if ext == "" {
			break
		}
		result = strings.TrimSuffix(result, ext)
	}
	return result
}

func decodeByte(cryptedByte byte, pwdByte byte, originalByte *byte, method string) error {
	var toReturn byte = cryptedByte

	if method != "x" && method != "r" && method != "xr" {
		return fmt.Errorf("'method' must be either \"x\", \"r\" or \"xr\", so \"%s\" is incorrect", method)
	}

	// This part make a byte rotation
	if strings.Contains(method, "r") {
		rotateN := pwdByte % 8
		toReturn = (toReturn >> rotateN) | (toReturn << (8 - rotateN))
	}

	if strings.Contains(method, "x") {
		toReturn = toReturn ^ pwdByte
	}

	*originalByte = toReturn

	return nil
}

func decodeChunk(cryptedChunk []byte, pwd []byte, pwdIndex *int, originalFile *os.File, method string) error {
	var (
		pwdLen            int    = len(pwd)
		originalFileBlock []byte = []byte{}
	)

	for _, cryptedByte := range cryptedChunk {
		if *pwdIndex >= pwdLen {
			*pwdIndex = 0
		}

		var pwdByte byte = pwd[*pwdIndex]
		var originalByte byte

		err := decodeByte(cryptedByte, pwdByte, &originalByte, method)

		if err != nil {
			return err
		}

		originalFileBlock = append(originalFileBlock, originalByte)

		(*pwdIndex)++
	}

	commonThings.Try(func() error {
		_, err := originalFile.Write(originalFileBlock)
		return err
	}, 3)

	return nil
}

func Decoder() {

	var method string

	translations := translate.GetTranslations()

	fmt.Println(translations.General.PressEnterToSelectFile)
	fmt.Scanln()
	var cryptedFilePath string
	var err error

	commonThings.Try(func() error {
		cryptedFilePath, err = commonThings.SelectFilePath(
			translations.General.SelectFile,
			commonThings.SelectFilePathFilters{
				{translations.General.EncodedBinFiles, "*.gfe*"},
				{translations.General.AllFiles, "*"},
			},
			"",
			"",
			commonThings.Load,
		)
		return err

	}, 3)

	fmt.Printf(translations.General.YouSelectedFile, cryptedFilePath)

	fmt.Print(translations.General.EnterPassword)

	pwd, _ := term.ReadPassword(int(syscall.Stdin))

	fmt.Println()

	fmt.Println(translations.General.PressEnterToSaveFile)
	fmt.Scanln()

	var originalFileProposition string

	splittedFilePath := strings.Split(cryptedFilePath, ".")

	ext := splittedFilePath[len(splittedFilePath)-1]

	switch ext {
	case "enc.bin": // Deprecated! Only for compatibility with v1.1.0 and older.
		originalFileProposition = removeExtentions(cryptedFilePath, 2)
		method = "x"
	case "gfe1":
		originalFileProposition = removeExtentions(cryptedFilePath, 1)
		method = "x"
	case "gfe2":
		fmt.Println("xr")
		originalFileProposition = removeExtentions(cryptedFilePath, 1)
		method = "xr"
	default:
		originalFileProposition = cryptedFilePath
		var method string

		fmt.Print(translations.Decoding.WhichEncodingMethodUsed)
		for !(method == "x" || method == "xr") {
			time.Sleep(300 * time.Millisecond)

			fmt.Printf("(x/xr)>>> ")
			fmt.Scanf("%s\n", &method)
			method = strings.ToLower(method)
			time.Sleep(300 * time.Millisecond)
		}
	}

	var originalFilePath string

	commonThings.Try(func() error {
		originalFilePath, err = commonThings.SelectFilePath(
			translations.General.SaveFile,
			commonThings.SelectFilePathFilters{},
			filepath.Base(originalFileProposition),
			"",
			commonThings.Save,
		)
		return err
	}, 3)

	fmt.Printf(translations.General.YouSelectedFile, originalFilePath)

	const blockSize int = 1024 * 16

	var (
		cryptedFile      *os.File
		cryptedFileBlock []byte = make([]byte, blockSize)
		originalFile     *os.File
	)

	commonThings.Try(func() error {
		cryptedFile, err = os.Open(cryptedFilePath)
		return err
	}, 3)

	commonThings.Try(func() error {
		originalFile, err = os.Create(originalFilePath)
		return err
	}, 3)

	var (
		timeBefore  time.Time = time.Now()
		timeAfter   time.Time
		timeBetween time.Duration
	)

	var pwdIndex int = 0

	for {

		isFinished := false
		commonThings.Try(func() error {
			readBytesNumber, err := cryptedFile.Read(cryptedFileBlock)

			if err == io.EOF && readBytesNumber == 0 {
				isFinished = true
				timeAfter = time.Now()

				timeBetween = timeAfter.Sub(timeBefore)
				fmt.Printf(translations.Decoding.FileDecodedIn, commonThings.FormatDuration(timeBetween))
				return nil
			}

			return err
		}, 3)

		if isFinished {
			break
		}

		decodeChunk(cryptedFileBlock, pwd, &pwdIndex, originalFile, method)

	}

	var cryptedFileStats os.FileInfo
	commonThings.Try(func() error {
		var err error
		cryptedFileStats, err = cryptedFile.Stat()
		return err
	}, 3)

	originalFile.Truncate(cryptedFileStats.Size())

}
