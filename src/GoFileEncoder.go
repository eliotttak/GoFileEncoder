package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/eliotttak/GoFileEncoder/assets"
	"github.com/eliotttak/GoFileEncoder/src/communFunctions"
	"github.com/eliotttak/GoFileEncoder/src/decoder"
	"github.com/eliotttak/GoFileEncoder/src/encoder"
)

var license string
var licenseByte []byte

func main() {

	communFunctions.Try(func() error {
		var err error
		licenseByte, err = assets.Asset("LICENSE")
		license = string(licenseByte)
		return err
	}, 3)

	fmt.Print("Que voulez-vous faire ?\n - Encoder un fichier (e)\n - Décoder un fichier (d)\n - Lire la license (l)\n")
	var rep string

	for !(rep == "e" || rep == "d" || rep == "l") {
		fmt.Print("(e/d)>>>")
		fmt.Scanf("%s\n", &rep)
		rep = strings.ToLower(rep)
		time.Sleep(300 * time.Millisecond)
	}

	switch rep {
	case "e":
		encoder.Encoder()
	case "d":
		decoder.Decoder()
	case "l":
		fmt.Println(license)
		main()
	}
}
