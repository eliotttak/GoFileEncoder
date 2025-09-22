package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/eliotttak/GoFileEncoder/assets"
	"github.com/eliotttak/GoFileEncoder/pkg/commonThings"
	"github.com/eliotttak/GoFileEncoder/pkg/decoder"
	"github.com/eliotttak/GoFileEncoder/pkg/encoder"
	"github.com/eliotttak/GoFileEncoder/pkg/translate"
)

var (
	license     string
	licenseByte []byte
)

func main() {
	translations := translate.GetTranslations()

	commonThings.Try(func() error {
		var err error
		licenseByte, err = assets.Asset("LICENSE")
		license = string(licenseByte)
		return err
	}, 3)

	fmt.Print(translations.Intro.WhatWouldYouLikeToDo)
	var rep string

	for !(rep == translations.Intro.EncodeInitial || rep == translations.Intro.DecodeInitial || rep == translations.Intro.LicenseInitial) {
		time.Sleep(300 * time.Millisecond)

		fmt.Printf(
			"(%s/%s/%s)>>>",
			translations.Intro.EncodeInitial,
			translations.Intro.DecodeInitial,
			translations.Intro.LicenseInitial,
		)
		fmt.Scanf("%s\n", &rep)
		rep = strings.ToLower(rep)
		time.Sleep(300 * time.Millisecond)
	}

	switch rep {
	case translations.Intro.EncodeInitial:
		encoder.Encoder()
	case translations.Intro.DecodeInitial:
		decoder.Decoder()
	case translations.Intro.LicenseInitial:
		fmt.Println(license)
		main()
	}
}
