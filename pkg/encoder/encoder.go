package encoder

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/eliotttak/GoFileEncoder/pkg/communFunctions"

	"github.com/sqweek/dialog"
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

	communFunctions.Try(func() error {
		_, err := cryptedFile.Write(cryptedFileBlock)
		return err
	}, 3)
}

func Encoder() {
	fmt.Println("Appuyez sur [Entrée] pour sélectionner un fichier...")
	fmt.Scanln()
	var originalFilePath string

	communFunctions.Try(func() error {
		var err error
		originalFilePath, err = dialog.File().Title("Sélectionner un fichier").Load()
		return err
	}, 3)

	fmt.Printf("Vous avez sélectionné ce fichier : %s.\n\n", originalFilePath)

	fmt.Print("Entrez le mot de passe : ")
	pwd, _ := term.ReadPassword(int(syscall.Stdin))

	fmt.Println()

	fmt.Println("Appuyez sur [Entrée] pour sauvegarder le fichier...")
	fmt.Scanln()
	var cryptedFilePath string

	communFunctions.Try(func() error {
		var err error
		cryptedFilePath, err = dialog.File().
			Title("Sauvegardez un fichier").
			Filter("Fichiers binaires encodés (.enc.bin)", "enc.bin").
			Filter("Tous les fichiers", "*").
			SetStartFile(filepath.Base(originalFilePath + ".enc.bin")).
			Save()
		return err
	}, 3)

	fmt.Printf("Vous avez sélectionné ce fichier : %s.\n\n", cryptedFilePath)

	const chunkSize int64 = 1024 * 16

	var (
		chunkStart        int64 = 0
		originalFile      *os.File
		originalFileChunk []byte = make([]byte, chunkSize)
		cryptedFile       *os.File
	)

	communFunctions.Try(func() error {
		var err error
		originalFile, err = os.Open(originalFilePath)
		return err
	}, 3)

	communFunctions.Try(func() error {
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

		communFunctions.Try(func() error {
			readBytesNumber, err := originalFile.Read(originalFileChunk)

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

		encodeChunk(originalFileChunk, pwd, &pwdIndex, cryptedFile)

		chunkStart += chunkSize
	}

	var originalFileStats os.FileInfo
	communFunctions.Try(func() error {
		var err error
		originalFileStats, err = originalFile.Stat()
		return err
	}, 3)

	cryptedFile.Truncate(originalFileStats.Size())
}
