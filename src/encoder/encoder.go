package encoder

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/sqweek/dialog"
	"golang.org/x/term"
)

func formatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	milliseconds := int(duration.Milliseconds()) % 1000
	return fmt.Sprintf("%d heure(s), %02d minute(s) et %02d.%03d seconde(s)", hours, minutes, seconds, milliseconds)
}

func try(f func() error, attempts int) {
	for {
		err := f()

		if err == nil {
			return
		} else {
			fmt.Println("Tentative échouée : " + err.Error())
			attempts--
		}

		if attempts == 0 {
			log.Fatal("Trop de tentatives échouées")
		}
	}
}

func Encoder() {
	fmt.Println("Appuyez sur [Entrée] pour sélectionner un fichier...")
	fmt.Scanln()
	var originalFilePath string
	var err error

	try(func() error {
		originalFilePath, err = dialog.File().Title("Sélectionner un fichier").Load()
		return err
	}, 3)

	fmt.Printf("Vous avez sélectionné ce fichier : %s.\n\n", originalFilePath)

	fmt.Print("Entrez le mot de passe : ")
	pwd, _ := term.ReadPassword(int(syscall.Stdin))

	var pwdLen int = len(pwd)

	fmt.Println("Appuyez sur [Entrée] pour sauvegarder le fichier...")
	fmt.Scanln()
	var cryptedFilePath string

	try(func() error {
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

	try(func() error {
		originalFile, err = os.Open(originalFilePath)
		return err
	}, 3)

	try(func() error {
		cryptedFile, err = os.Create(cryptedFilePath)
		return err
	}, 3)

	var (
		timeBefore  time.Time
		timeAfter   time.Time
		timeBetween time.Duration
	)

	for {

		try(func() error {
			_, err = originalFile.Seek(chunkStart, io.SeekStart)
			return err
		}, 3)

		isFinished := false

		try(func() error {
			_, err = originalFile.Read(originalFileChunk)

			if err == io.EOF {
				isFinished = true
				timeAfter = time.Now()

				timeBetween = timeAfter.Sub(timeBefore)
				fmt.Printf("Fichier encodé en %s.\n", formatDuration(timeBetween))
				return nil
			}
			return err
		}, 3)

		if isFinished {
			return
		}

		var cryptedFileChunk []byte = []byte{}

		var originalByte byte

		var pwdIndex int = 0

		timeBefore = time.Now()

		for _, originalByte = range originalFileChunk {
			if pwdIndex >= pwdLen {
				pwdIndex = 0
			}

			var pwdByte byte = pwd[pwdIndex]
			var cryptedByte byte = originalByte ^ pwdByte

			cryptedFileChunk = append(cryptedFileChunk, cryptedByte)

			pwdIndex++
		}

		try(func() error {
			_, err := cryptedFile.Write(cryptedFileChunk)
			return err
		}, 3)

		chunkStart += chunkSize

	}

}
