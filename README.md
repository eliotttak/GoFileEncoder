# GoFileEncoder

> &#x1F1EB;&#x1F1F7; Vous êtes français ? Ouvrez [LISEZMOI.md](LISEZMOI.md) en français

## Introduction

GoFileEncoder is a little file encoder that uses the [XOR encryption](https://en.wikipedia.org/wiki/XOR_cipher). It is coded in Golang, and is compiled for several OSs (Windows&reg;, Linux&reg;, and soon macOS&reg;).

I am not a pro in Go, so if you find a bug, or simply want to make a suggestion, I am totally open to [_issues_](https://github.com/eliotttak/GoFileEncoder/issues), [_pull requests_](https://github.com/eliotttak/GoFileEncoder/pulls) and to the [discussion](https://github.com/eliotttak/GoFileEncoder/discussions).

---

## Summary
- [GoFileEncoder](#gofileencoder)
  - [Introduction](#introduction)
  - [Summary](#summary)
  - [Building](#building)
    - [Cloning](#cloning)
    - [Dependencies tree](#dependencies-tree)
    - [Dependencies installing](#dependencies-installing)
    - [Creating assets package](#creating-assets-package)
    - [Compiling](#compiling)
      - [In Bash](#in-bash)
      - [In Batch](#in-batch)
  - [Usage](#usage)
    - [1. Encode or decode?](#1-encode-or-decode)
    - [2. Select a file](#2-select-a-file)
    - [3. Enter the password](#3-enter-the-password)
    - [4. Choose the destination file](#4-choose-the-destination-file)
    - [5. Wait...](#5-wait)
    - [6. All's done!](#6-alls-done)
  - [License](#license)


## Building

### Cloning

```bash
git clone https://github.com/eliotttak/GoFileEncoder
cd GoFileEncoder
```

### Dependencies tree

```plaintext
github.com/eliotttak/GoFileEncoder (this project)
|
+-- github.com/sqweek/dialog (for the file popups)
|   |
|    \_ github.com/TheTitanrain/w32 (indirect)
|
 \_ golang.org/x/term (for the password asking)
    |
     \_ golang.org/x/sys (indirect)
```

Refer to the file [go.mod](./go.mod) for more details.
> &#x1F6C8; In [go.mod](./go.mod), there is an import for [github.com/abdfnx/gosh](https://github.com/abdfnx/gosh). It was used only for my setup, and you shouldn't need it. It will disappear if you run `go mod tidy`.

### Dependencies installing
```bash
go mod tidy
go get
```

### Creating assets package
```bash
go-bindata -pkg assets -o assets/bindata.go LICENSE
```

### Compiling

#### In Bash
```bash
go build -o ./bin/GoFileEncoder ./src/
```

#### In Batch
```batch
go build -o .\bin\GoFileEncoder .\src\
```

---

## Usage

### 1. Encode or decode?

At first, run the program. It will ask you if you want to decode or encode a file :
```plaintext
Que voulez-vous faire ?
 - Encoder un fichier (e)
 - Décoder un fichier (d)
(e/d)>>>
```

If you want to encode a file, enter <kbd>e</kbd>, else, enter <kbd>d</kbd>. In all the cases, confirm with <kbd>Enter</kbd>.

### 2. Select a file

Then, it will ask you to press <kbd>Enter</kbd> to select a file. Do it, and a popup will aappear. Select your file, then confirm.
> &#x1F6C8; If you cancel, the popup will appear back twice, then at the 3<sup>rd</sup>, the message `"Trop de tentatives échouées"` (`Too many failed attempts`) will be displayed and the program will be closed.

### 3. Enter the password

After it, you will be asked for a password. Enter it then confirm by pressing <kbd>Enter</kbd>.
> &#x26A0; If you choosed `Decode a file`, be careful to write the password correctly. If you misspell it by even one character, the file will be corrupted and no longer usable.

> &#x1F6C8; To ensure that your password remains confidential, it will not be displayed on your screen.

### 4. Choose the destination file

Then, the program will ask you to press <kbd>Entrée</kbd> to choose the destination file. Do it, and a popup will aappear. Select the file to create, then confirm.
> &#x1F6C8; If you cancel, the popup will appear back twice, then at the 3<sup>rd</sup>, the message `"Trop de tentatives échouées"` (`Too many failed attempts`) will be displayed and the program will be closed.

### 5. Wait...

The file is being encoded. Do not close the program.

### 6. All's done!
After a few seconds, the file is encoded or decoded. The programme closes automatically.

## License
This software is distributed under the GNU GENERAL PUBLIC LICENSE version 3 (GNU GPL v3).
[See the license](LICENSE)