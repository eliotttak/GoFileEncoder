package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/eliotttak/GoFileEncoder/assets"
	"github.com/eliotttak/GoFileEncoder/pkg/communThings"
	"github.com/eliotttak/GoFileEncoder/pkg/decoder"
	"github.com/eliotttak/GoFileEncoder/pkg/encoder"
	"github.com/eliotttak/GoFileEncoder/pkg/translations"
)

var (
	license     string
	licenseByte []byte
)

func main() {

	communThings.Try(func() error {
		var err error
		licenseByte, err = assets.Asset("LICENSE")
		license = string(licenseByte)
		return err
	}, 3)

	fmt.Print(translations.GetTranslations().WhatWouldYouLikeToDo)
	var rep string

	for !(rep == translations.GetTranslations().EncodeInitial || rep == translations.GetTranslations().DecodeInitial || rep == translations.GetTranslations().LicenseInitial) {
		time.Sleep(300 * time.Millisecond)

		fmt.Printf(
			"(%s/%s/%s)>>>",
			translations.GetTranslations().EncodeInitial,
			translations.GetTranslations().DecodeInitial,
			translations.GetTranslations().LicenseInitial,
		)
		fmt.Scanf("%s\n", &rep)
		rep = strings.ToLower(rep)
		time.Sleep(300 * time.Millisecond)
	}

	switch rep {
	case translations.GetTranslations().EncodeInitial:
		encoder.Encoder()
	case translations.GetTranslations().DecodeInitial:
		decoder.Decoder()
	case translations.GetTranslations().LicenseInitial:
		fmt.Println(license)
		main()
	}
}
