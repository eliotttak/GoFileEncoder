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
package encoder

import (
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/eliotttak/GoFileEncoder/pkg/commonThings"
	"github.com/eliotttak/GoFileEncoder/pkg/translate"
	"github.com/nbutton23/zxcvbn-go"

	"golang.org/x/term"
)

func encodeByte(originalByte byte, pwdByte byte, cryptedByte *byte, method string) error {
	var toReturn byte = originalByte

	if method != "x" && method != "r" && method != "xr" {
		return fmt.Errorf("'method' must be either \"x\", \"r\" or \"xr\", so \"%s\" is incorrect", method)
	}

	if strings.Contains(method, "x") {
		toReturn = toReturn ^ pwdByte
	}

	// This part make a byte rotation
	if strings.Contains(method, "r") {
		var rotateN byte = pwdByte % 8
		toReturn = (toReturn << rotateN) | (toReturn >> (8 - rotateN)) // If leftPart = 00100000 and rightPart = 00001001, so originalByte = 00101001

	}

	*cryptedByte = toReturn

	return nil
}

func encodeChunk(originalChunk []byte, pwd []byte, pwdIndex *int, cryptedFile *os.File, method string) {
	var (
		pwdLen           int    = len(pwd)
		cryptedFileBlock []byte = []byte{}
	)

	for _, originalByte := range originalChunk {
		if *pwdIndex >= pwdLen {
			*pwdIndex = 0
		}

		var pwdByte byte = pwd[*pwdIndex]
		var cryptedByte byte

		encodeByte(originalByte, pwdByte, &cryptedByte, method)

		cryptedFileBlock = append(cryptedFileBlock, cryptedByte)

		(*pwdIndex)++
	}

	commonThings.Try(func() error {
		_, err := cryptedFile.Write(cryptedFileBlock)
		return err
	}, 3)
}

func Encoder() {
	translations := translate.GetTranslations()

	fmt.Println(translations.General.PressEnterToSelectFile)
	fmt.Scanln()

	var originalFilePath string

	commonThings.Try(func() error {
		var err error
		originalFilePath, err = commonThings.SelectFilePath(
			translations.General.SelectFile,
			commonThings.SelectFilePathFilters{},
			"",
			"",
			commonThings.Load,
		)
		return err
	}, 3)

	fmt.Printf(translations.General.YouSelectedFile, originalFilePath)

	var method string

	fmt.Print(translations.Encoding.WhichEncodingMethod)
	for !(method == "x" || method == "xr") {
		time.Sleep(300 * time.Millisecond)

		fmt.Printf("(x/xr)>>> ")
		fmt.Scanf("%s\n", &method)
		method = strings.ToLower(method)
		time.Sleep(300 * time.Millisecond)
	}

	var pwd []byte

	for {
		fmt.Print(translations.General.EnterPassword)
		pwd, _ = term.ReadPassword(int(syscall.Stdin))

		scoreFl64 := zxcvbn.PasswordStrength(string(pwd), []string{}).CrackTime

		score := time.Duration(scoreFl64 * math.Pow10(9))

		fmt.Println()

		fmt.Printf(translations.Encoding.CrackedIn, commonThings.Tern(scoreFl64 > 9223372036.854775807, translations.Encoding.MoreThan290, commonThings.FormatDuration(score)))

		fmt.Println(translations.Encoding.IsPwdOk)
		var rep string
		for !(rep == translations.Encoding.YesInitial || rep == translations.Encoding.NoInitial) {
			time.Sleep(300 * time.Millisecond)

			fmt.Printf(
				"(%s/%s)>>>",
				translations.Encoding.YesInitial,
				translations.Encoding.NoInitial,
			)
			fmt.Scanf("%s\n", &rep)
			rep = strings.ToLower(rep)
			time.Sleep(300 * time.Millisecond)
		}

		if rep == translations.Encoding.YesInitial {
			break
		}
	}

	fmt.Println()

	fmt.Println(translations.General.PressEnterToSaveFile)
	fmt.Scanln()

	ext := commonThings.Tern(method == "x", ".gfe1", "gfe2") // .gfe1 files are encoded with only one step (XOR), and .gfe2 files are encoded with two steps (XOR and Rotate)

	var cryptedFilePath string

	commonThings.Try(func() error {
		var err error
		cryptedFilePath, err = commonThings.SelectFilePath(
			translations.General.SaveFile,
			commonThings.SelectFilePathFilters{
				{translations.General.EncodedBinFiles, ext},
				{translations.General.AllFiles, "*"},
			},
			filepath.Base(originalFilePath+"."+ext),
			"",
			commonThings.Save,
		)
		return err
	}, 3)

	fmt.Printf(translations.General.YouSelectedFile, cryptedFilePath)

	const chunkSize int64 = 1024 * 16

	var (
		chunkStart        int64 = 0
		originalFile      *os.File
		originalFileChunk []byte = make([]byte, chunkSize)
		cryptedFile       *os.File
	)

	commonThings.Try(func() error {
		var err error
		originalFile, err = os.Open(originalFilePath)
		return err
	}, 3)

	commonThings.Try(func() error {
		var err error
		cryptedFile, err = os.Create(cryptedFilePath)
		return err
	}, 3)

	var (
		timeBefore  time.Time
		timeAfter   time.Time
		timeBetween time.Duration
	)

	timeBefore = time.Now()

	var pwdIndex int = 0

	for {
		isFinished := false

		commonThings.Try(func() error {
			readBytesNumber, err := originalFile.Read(originalFileChunk)

			if err == io.EOF && readBytesNumber == 0 {
				isFinished = true
				timeAfter = time.Now()

				timeBetween = timeAfter.Sub(timeBefore)
				fmt.Printf(translations.Encoding.FileEncodedIn, commonThings.FormatDuration(timeBetween))
				return nil
			}
			return err
		}, 3)

		if isFinished {
			break
		}

		encodeChunk(originalFileChunk, pwd, &pwdIndex, cryptedFile, method)

		chunkStart += chunkSize
	}

	var originalFileStats os.FileInfo
	commonThings.Try(func() error {
		var err error
		originalFileStats, err = originalFile.Stat()
		return err
	}, 3)

	cryptedFile.Truncate(originalFileStats.Size())
}
