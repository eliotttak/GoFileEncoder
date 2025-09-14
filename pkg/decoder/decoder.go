package decoder

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/eliotttak/GoFileEncoder/pkg/communFunctions"

	"github.com/sqweek/dialog"
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

	communFunctions.Try(func() error {
		_, err := originalFile.Write(originalFileBlock)
		return err
	}, 3)
}

func Decoder() {
	fmt.Println("Appuyez sur [Entrée] pour sélectionner un fichier...")
	fmt.Scanln()
	var cryptedFilePath string
	var err error

	communFunctions.Try(func() error {
		cryptedFilePath, err = dialog.File().
			Filter("Fichiers binaires encodés (.enc.bin)", "enc.bin").
			Filter("Tous les fichiers", "*").
			Title("Sélectionner un fichier").
			Load()
		return err
	}, 3)

	fmt.Printf("Vous avez sélectionné ce fichier : %s.\n\n", cryptedFilePath)

	fmt.Print("Entrez le mot de passe : ")

	pwd, _ := term.ReadPassword(int(syscall.Stdin))

	fmt.Println()

	fmt.Println("Appuyez sur [Entrée] pour sauvegarder le fichier...")
	fmt.Scanln()

	var originalFileProposition string

	if strings.HasSuffix(cryptedFilePath, ".enc.bin") {
		originalFileProposition = removeExtentions(cryptedFilePath, 2)
	} else {
		originalFileProposition = removeExtentions(cryptedFilePath, 1)
	}

	var originalFilePath string

	communFunctions.Try(func() error {
		originalFilePath, err = dialog.File().
			Title("Sauvegardez un fichier").
			SetStartFile(filepath.Base(originalFileProposition)).
			Save()
		return err
	}, 3)

	fmt.Printf("Vous avez sélectionné ce fichier : %s.\n\n", originalFilePath)

	const blockSize int = 1024 * 16

	var (
		cryptedFile      *os.File
		cryptedFileBlock []byte = make([]byte, blockSize)
		originalFile     *os.File
	)

	communFunctions.Try(func() error {
		cryptedFile, err = os.Open(cryptedFilePath)
		return err
	}, 3)

	communFunctions.Try(func() error {
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
		communFunctions.Try(func() error {
			readBytesNumber, err := cryptedFile.Read(cryptedFileBlock)

			if err == io.EOF && readBytesNumber == 0 {
				isFinished = true
				timeAfter = time.Now()

				timeBetween = timeAfter.Sub(timeBefore)
				fmt.Printf("Fichier encodé en %s.\n", communFunctions.FormatDuration(timeBetween))
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
	communFunctions.Try(func() error {
		var err error
		cryptedFileStats, err = cryptedFile.Stat()
		return err
	}, 3)

	originalFile.Truncate(cryptedFileStats.Size())

}
