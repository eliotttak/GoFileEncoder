package encoder

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sqweek/dialog"
)

func formatDuration(duration time.Duration) string {
	var hours int = int(duration.Hours())
	var minutes int = int(duration.Minutes()) % 60
	var seconds int = int(duration.Seconds()) % 60
	var milliseconds int = int(duration.Milliseconds()) % 1000
	return fmt.Sprintf("%d heure(s), %02d minute(s) et %02d.%03d seconde(s)", hours, minutes, seconds, milliseconds)
}

func Encoder() {
	fmt.Println("Appuyez sur [Entrée] pour sélectionner un fichier...")
	fmt.Scanln()
	var originalFilePath string
	var err error
	var attempts int = 0
	for {
		originalFilePath, err = dialog.File().Title("Sélectionner un fichier").Load()

		if err == nil {
			break
		} else {
			fmt.Println("Tentative échouée : " + err.Error())
			attempts++
		}

		if attempts == 3 {
			fmt.Println("Trop de tentatives échouées")
			return
		}
	}

	fmt.Printf("Vous avez sélectionné ce fichier : %s.\n\n", originalFilePath)

	var pwdStr string

	fmt.Print("Entrez le mot de passe : ")
	fmt.Scanf("%s\n", &pwdStr)

	var pwd []byte = []byte(pwdStr)

	var pwdLen int = len(pwd)

	fmt.Println("Appuyez sur [Entrée] pour sauvegarder le fichier...")
	fmt.Scanln()
	var cryptedFilePath string
	for {
		cryptedFilePath, err = dialog.File().
			Title("Sauvegardez un fichier").
			Filter("Fichiers binaires encodés (.enc.bin)", "enc.bin").
			Filter("Tous les fichiers", "*").
			SetStartFile(filepath.Base(originalFilePath + ".enc.bin")).
			Save()

		if err == nil {
			break
		} else {
			fmt.Println("Tentative échouée : " + err.Error())
			attempts++
		}

		if attempts == 3 {
			fmt.Println("Trop de tentatives échouées")
			return
		}
	}

	fmt.Printf("Vous avez sélectionné ce fichier : %s.\n\n", cryptedFilePath)

	const blockSize int = 1024 * 16

	var (
		blockStart        int64 = 0
		originalFile      *os.File
		originalFileBlock []byte = make([]byte, blockSize)
		cryptedFile       *os.File
	)

	attempts = 0
	for {
		originalFile, err = os.Open(originalFilePath)

		if err == nil {
			defer originalFile.Close()
			break
		} else {
			fmt.Println("Tentative échouée : " + err.Error())
			attempts++
		}

		if attempts == 3 {
			fmt.Println("Trop de tentatives échouées")
			return
		}
	}

	attempts = 0
	for {
		cryptedFile, err = os.Create(cryptedFilePath)

		if err == nil {
			defer cryptedFile.Close()
			break
		} else {
			fmt.Println("Tentative échouée : " + err.Error())
			attempts++
		}

		if attempts == 3 {
			fmt.Println("Trop de tentatives échouées")
			return
		}
	}

	var (
		timeBefore  time.Time
		timeAfter   time.Time
		timeBetween time.Duration
	)

	for {
		attempts = 0
		for {
			_, err = originalFile.Seek(blockStart, io.SeekStart)
			if err == nil {
				break
			} else {
				fmt.Println("Tentative échouée : " + err.Error())
				attempts++
			}

			if attempts == 3 {
				fmt.Println("Trop de tentatives échouées")
				return
			}
		}

		attempts = 0
		for {
			_, err = originalFile.Read(originalFileBlock)
			if err == nil {
				break
			} else if err == io.EOF {
				timeAfter = time.Now()

				timeBetween = timeAfter.Sub(timeBefore)
				fmt.Printf("Fichier encodé en %s.\n", formatDuration(timeBetween))
				return
			} else {
				fmt.Println("Tentative échouée : " + err.Error())
				attempts++
			}

			if attempts == 3 {
				fmt.Println("Trop de tentatives échouées")
				return
			}

		}

		var cryptedFileBlock []byte = []byte{}

		var originalByte byte

		var pwdIndex int = 0

		timeBefore = time.Now()

		for _, originalByte = range originalFileBlock {
			if pwdIndex >= pwdLen {
				pwdIndex = 0
			}

			var pwdByte byte = pwd[pwdIndex]
			var cryptedByte byte = originalByte ^ pwdByte

			cryptedFileBlock = append(cryptedFileBlock, cryptedByte)

			pwdIndex++
		}

		attempts = 0
		for {
			_, err = cryptedFile.Seek(0, io.SeekEnd)

			if err == nil {
				break
			} else {
				fmt.Println("Tentative échouée : " + err.Error())
				attempts++
			}

			if attempts == 3 {
				fmt.Println("Trop de tentatives échouées")
				return
			}
		}

		attempts = 0
		for {
			_, err = cryptedFile.Write(cryptedFileBlock)

			if err == nil {
				break
			} else {
				fmt.Println("Tentative échouée : " + err.Error())
				attempts++
			}

			if attempts == 3 {
				fmt.Println("Trop de tentatives échouées")
				return
			}
		}
		blockStart += 16 * 1024

	}

}
