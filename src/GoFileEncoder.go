package main

import (
	"GoFileEncoder/src/decoder"
	"GoFileEncoder/src/encoder"
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Print("Que voulez-vous faire ?\n - Encoder un fichier (e)\n - Décoder un fichier (d)\n")
	var rep string = ""

	for !(rep == "e" || rep == "d") {
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
	}
}
