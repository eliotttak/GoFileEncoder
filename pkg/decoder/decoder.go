package decoder

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/eliotttak/GoFileEncoder/pkg/communThings"
	"github.com/eliotttak/GoFileEncoder/pkg/translations"

	"golang.org/x/term"
)

func removeExtentions(path string, nExtensions int) string {
	var result string = path
	for x := 0; x < nExtensions; x++ {
		var ext string = filepath.Ext(result)
		if ext == "" {
			break
		}
		result = strings.TrimSuffix(result, ext)
	}
	return result
}

func decodeByte(cryptedByte byte, pwdByte byte, originalByte *byte) {
	*originalByte = cryptedByte ^ pwdByte
}

func decodeChunk(cryptedChunk []byte, pwd []byte, pwdIndex *int, originalFile *os.File) {

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

		decodeByte(cryptedByte, pwdByte, &originalByte)

		originalFileBlock = append(originalFileBlock, originalByte)

		(*pwdIndex)++
	}

	communThings.Try(func() error {
		_, err := originalFile.Write(originalFileBlock)
		return err
	}, 3)
}

func Decoder() {
	fmt.Println(translations.GetTranslations().PressEnterToSelectFile)
	fmt.Scanln()
	var cryptedFilePath string
	var err error

	communThings.Try(func() error {
		cryptedFilePath, err = communThings.SelectFilePath(
			translations.GetTranslations().SelectFile,
			communThings.SelectFilePathFilters{
				{translations.GetTranslations().EncodedBinFiles, "enc.bin"},
				{translations.GetTranslations().AllFiles, "*"},
			},
			"",
			"",
			communThings.Load,
		)
		return err

	}, 3)

	fmt.Printf(translations.GetTranslations().YouSelectedFile, cryptedFilePath)

	fmt.Print(translations.GetTranslations().EnterPassword)

	pwd, _ := term.ReadPassword(int(syscall.Stdin))

	fmt.Println()

	fmt.Println(translations.GetTranslations().PressEnterToSaveFile)
	fmt.Scanln()

	var originalFileProposition string

	if strings.HasSuffix(cryptedFilePath, ".enc.bin") {
		originalFileProposition = removeExtentions(cryptedFilePath, 2)
	} else {
		originalFileProposition = removeExtentions(cryptedFilePath, 1)
	}

	var originalFilePath string

	communThings.Try(func() error {
		originalFilePath, err = communThings.SelectFilePath(
			translations.GetTranslations().SaveFile,
			communThings.SelectFilePathFilters{},
			filepath.Base(originalFileProposition),
			"",
			communThings.Save,
		)
		return err
	}, 3)

	fmt.Printf(translations.GetTranslations().YouSelectedFile, originalFilePath)

	const blockSize int = 1024 * 16

	var (
		cryptedFile      *os.File
		cryptedFileBlock []byte = make([]byte, blockSize)
		originalFile     *os.File
	)

	communThings.Try(func() error {
		cryptedFile, err = os.Open(cryptedFilePath)
		return err
	}, 3)

	communThings.Try(func() error {
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
		communThings.Try(func() error {
			readBytesNumber, err := cryptedFile.Read(cryptedFileBlock)

			if err == io.EOF && readBytesNumber == 0 {
				isFinished = true
				timeAfter = time.Now()

				timeBetween = timeAfter.Sub(timeBefore)
				fmt.Printf(translations.GetTranslations().FileDecodedIn, communThings.FormatDuration(timeBetween))
				return nil
			}

			return err
		}, 3)

		if isFinished {
			break
		}

		decodeChunk(cryptedFileBlock, pwd, &pwdIndex, originalFile)

	}

	var cryptedFileStats os.FileInfo
	communThings.Try(func() error {
		var err error
		cryptedFileStats, err = cryptedFile.Stat()
		return err
	}, 3)

	originalFile.Truncate(cryptedFileStats.Size())

}
