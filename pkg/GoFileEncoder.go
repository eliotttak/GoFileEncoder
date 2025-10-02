// Copyright 2025 Eliott Takvorian
//
// This file is part of GoFileEncoder.
//
// GoFileEncoder is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.
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
