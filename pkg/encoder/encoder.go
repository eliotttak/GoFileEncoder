package encoder

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/eliotttak/GoFileEncoder/pkg/commonThings"
	"github.com/eliotttak/GoFileEncoder/pkg/translate"

	"golang.org/x/term"
)

func encodeByte(originalByte byte, pwdByte byte, cryptedByte *byte) {
	*cryptedByte = originalByte ^ pwdByte
}

func encodeChunk(originalChunk []byte, pwd []byte, pwdIndex *int, cryptedFile *os.File) {

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

		encodeByte(originalByte, pwdByte, &cryptedByte)

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

	fmt.Print(translations.General.EnterPassword)
	pwd, _ := term.ReadPassword(int(syscall.Stdin))

	fmt.Println()

	fmt.Println(translations.General.PressEnterToSaveFile)
	fmt.Scanln()
	var cryptedFilePath string

	commonThings.Try(func() error {
		var err error
		cryptedFilePath, err = commonThings.SelectFilePath(
			translations.General.SaveFile,
			commonThings.SelectFilePathFilters{
				{translations.General.EncodedBinFiles, "enc.bin"},
				{translations.General.AllFiles, "*"},
			},
			filepath.Base(originalFilePath+".enc.bin"),
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

		encodeChunk(originalFileChunk, pwd, &pwdIndex, cryptedFile)

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
