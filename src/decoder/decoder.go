package decoder

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func Decoder() {
	fmt.Println("Appuyez sur [Entrée] pour sélectionner un fichier...")
	fmt.Scanln()
	var cryptedFilePath string
	var err error
	var attempts int = 0
	for {
		cryptedFilePath, err = dialog.File().
			Filter("Fichiers binaires encodés (.enc.bin)", "enc.bin").
			Filter("Tous les fichiers", "*").
			Title("Sélectionner un fichier").
			Load()

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

	var pwdStr string

	fmt.Print("Entrez le mot de passe : ")
	fmt.Scanf("%s", &pwdStr)

	var pwd []byte = []byte(pwdStr)

	var pwdLen int = len(pwd)

	fmt.Println("Appuyez sur [Entrée] pour sauvegarder le fichier...")
	fmt.Scanln()

	var originalFileProposition string

	if strings.HasSuffix(cryptedFilePath, ".enc.bin") {
		originalFileProposition = removeExtentions(cryptedFilePath, 2)
	} else {
		originalFileProposition = removeExtentions(cryptedFilePath, 1)
	}

	var originalFilePath string

	for {
		originalFilePath, err = dialog.File().
			Title("Sauvegardez un fichier").
			SetStartFile(filepath.Base(originalFileProposition)).
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

	fmt.Printf("Vous avez sélectionné ce fichier : %s.\n\n", originalFilePath)

	const blockSize int = 1024 * 16

	var (
		blockStart       int64 = 0
		cryptedFile      *os.File
		cryptedFileBlock []byte = make([]byte, blockSize)
		originalFile     *os.File
	)

	attempts = 0
	for {
		cryptedFile, err = os.Open(cryptedFilePath)

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

	attempts = 0
	for {
		originalFile, err = os.Create(originalFilePath)

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

	var (
		timeBefore  time.Time
		timeAfter   time.Time
		timeBetween time.Duration
	)

	for {
		attempts = 0
		for {
			_, err = cryptedFile.Seek(blockStart, io.SeekStart)
			if err == nil {
				break
			} else {
				fmt.Println("Tentative échouée (line 163) : " + err.Error())
				attempts++
			}

			if attempts == 3 {
				fmt.Println("Trop de tentatives échouées")
				return
			}
		}

		attempts = 0
		for {
			_, err = cryptedFile.Read(cryptedFileBlock)
			if err == nil {
				break
			} else if err == io.EOF {
				timeAfter = time.Now()

				timeBetween = timeAfter.Sub(timeBefore)
				fmt.Printf("Fichier décodé en %s.\n", formatDuration(timeBetween))
				return
			} else {
				fmt.Println("Tentative échouée (line 185) : " + err.Error())
				attempts++
			}

			if attempts == 3 {
				fmt.Println("Trop de tentatives échouées")
				return
			}

		}

		var originalFileBlock []byte = []byte{}

		var cryptedByte byte

		var pwdIndex int = 0

		timeBefore = time.Now()

		for _, cryptedByte = range cryptedFileBlock {
			if pwdIndex >= pwdLen {
				pwdIndex = 0
			}

			var pwdByte byte = pwd[pwdIndex]
			var originalByte byte = cryptedByte ^ pwdByte

			originalFileBlock = append(originalFileBlock, originalByte)

			pwdIndex++
		}

		attempts = 0
		for {
			_, err = originalFile.Seek(0, io.SeekEnd)

			if err == nil {
				break
			} else {
				fmt.Println("Tentative échouée (line 208): " + err.Error())
				attempts++
			}

			if attempts == 3 {
				fmt.Println("Trop de tentatives échouées")
				return
			}
		}

		attempts = 0
		for {
			_, err = originalFile.Write(originalFileBlock)

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
