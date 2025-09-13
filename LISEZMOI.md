# GoFileEncoder

> &#x1F1EC;&#x1F1E7; Are you English ? Open [README.md](README.md) in English.


## Introduction

GoFileEncoder est un petit encodeur de fichiers qui utilise l'[encryption XOR <sup>(EN)</sup>](https://en.wikipedia.org/wiki/XOR_cipher). Il est codé en Golang, et est compilé pour plusieurs OS (Windows&reg;, Linux&reg;,et bientôt macOS&reg;).

Je ne suis pas un pro en Go, donc si vous trouvez un bug, ou simplement voulez faire une suggestion, je suis totalement ouvert aux [_issues_](https://github.com/eliotttak/GoFileEncoder/issues), [_pull requests_](https://github.com/eliotttak/GoFileEncoder/pulls) et à la [discussion](https://github.com/eliotttak/GoFileEncoder/discussions).

---

## Sommaire
- [GoFileEncoder](#gofileencoder)
  - [Introduction](#introduction)
  - [Sommaire](#sommaire)
  - [Construction](#construction)
    - [Clonage](#clonage)
    - [Arbre des dépendances](#arbre-des-dépendances)
    - [Installation des dépendances](#installation-des-dépendances)
    - [Création du _package_ d'_assets_](#création-du-package-dassets)
    - [Compilation](#compilation)
      - [En Bash](#en-bash)
      - [En Batch](#en-batch)
  - [Utilisation](#utilisation)
    - [1. Encoder ou décoder ?](#1-encoder-ou-décoder-)
    - [2. Sélectionnez un fichier](#2-sélectionnez-un-fichier)
    - [3. Entrez le mot de passe](#3-entrez-le-mot-de-passe)
    - [4. Choisissez le fichier de destination](#4-choisissez-le-fichier-de-destination)
    - [5. Attendez...](#5-attendez)
    - [6. Terminé !](#6-terminé-)
  - [License](#license)
  - [Introduction](#introduction-1)
  - [Sommaire](#sommaire-1)
  - [Construction](#construction-1)
    - [Clonage](#clonage-1)
    - [Arbre des dépendances](#arbre-des-dépendances-1)
    - [Installation des dépendances](#installation-des-dépendances-1)
    - [Compilation](#compilation-1)
      - [En Bash](#en-bash-1)
      - [En Batch](#en-batch-1)
  - [Utilisation](#utilisation-1)
    - [1. Encoder ou décoder ?](#1-encoder-ou-décoder--1)
    - [2. Sélectionnez un fichier](#2-sélectionnez-un-fichier-1)
    - [3. Entrez le mot de passe](#3-entrez-le-mot-de-passe-1)
    - [4. Choisissez le fichier de destination](#4-choisissez-le-fichier-de-destination-1)
    - [5. Attendez...](#5-attendez-1)
    - [6. Terminé !](#6-terminé--1)
  - [License](#license-1)


## Construction

### Clonage

```bash
git clone https://github.com/eliotttak/GoFileEncoder
cd GoFileEncoder
```

### Arbre des dépendances

```plaintext
GoFileEncoder (ce projet)
|
+-- github.com/sqweek/dialog (pour les popups des fichiers)
|   |
|    \_ github.com/TheTitanrain/w32 (indirect)
|
 \_ golang.org/x/term (pour la demande de mot de passe)
    |
     \_ golang.org/x/sys (indirect)
```

Se réferer au fichier [go.mod](./go.mod) pour plus de détails.
> &#x1F6C8; Dans [go.mod](./go.mod), il y a une directive d'import pour [github.com/abdfnx/gosh](https://github.com/abdfnx/gosh). Il a été utilisé seulement pour mon installateur, et vous ne devriez pas en avoir besoin. Il devrait disparaître si vous exécutez `go mod tidy`.

### Installation des dépendances
```bash
go mod tidy
go get
```

### Création du _package_ d'_assets_
```bash
go-bindata -pkg assets -o assets/bindata.go LICENSE
```

### Compilation

#### En Bash
```bash
go build -o ./bin/GoFileEncoder ./src/
```

#### En Batch
```batch
go build -o .\bin\GoFileEncoder .\src\
```

---

## Utilisation

### 1. Encoder ou décoder ?

En premier, lancez le programme. Celui-ci va vous demander si vous souhaitez encoder ou décoder un fichier :

```plaintext
Que voulez-vous faire ?
 - Encoder un fichier (e)
 - Décoder un fichier (d)
(e/d)>>>
```
Si vous voulez encoder un fichier, entrez <kbd>e</kbd>, sinon, entrez <kbd>d</kbd>. Dans tous les cas, validez avec <kbd>Entrée</kbd>.

### 2. Sélectionnez un fichier

Ensuite, il va vous demander d'appuyez sur <kbd>Entrée</kbd> pour sélectionner un fichier. Faites-le, et un popup va apparaître. Sélectionnez votre fichier, puis validez.
> &#x1F6C8; Si vous annulez, le popup va réapparaître 2 fois, puis à la 3<sup>ème</sup>, le message `"Trop de tentatives échouées"` va s'afficher puis le programme va se fermer.

### 3. Entrez le mot de passe

Après cela, un mot de passe vous sera demandé. Entrez-le puis validez avec <kbd>Entrée</kbd>.
> &#x26A0; Si vous avez choisi `Décoder un fichier`, faîtes attention à bien écrire le mot de passe. Si vous vous trompez ne serait-ce que d'un caractère, le fichier sera corrompu et ne sera plus utilisable.

> &#x1F6C8; Pour garantir la confidentialité de votre mot de passe, celui-ci ne s'affichera pas sur votre écran.

### 4. Choisissez le fichier de destination

Ensuite, le programme va vous demander d'appuyez sur <kbd>Entrée</kbd> pour choisir le fichier de destination. Faites-le, et un popup va apparaître. Sélectionnez le fichier à créer, puis validez.
> &#x1F6C8; Si vous annulez, le popup va réapparaître 2 fois, puis à la 3<sup>ème</sup>, le message `"Trop de tentatives échouées"` va s'afficher puis le programme va se fermer.

### 5. Attendez...

Le fichier est en train d'être encodé. Ne fermez pas le programme.

### 6. Terminé !
Au bout de quelques secondes, le fichier est encodé ou décodé. Le programme se ferme automatiquement.

## License
Ce logiciel est distribué sous la license GNU GENERAL PUBLIC LICENSE version 3 (GNU GPL v3).
[Voir la license](..\# <small>(FR)</small> GoFileEncoder

## Introduction

GoFileEncoder est un petit encodeur de fichiers qui utilise l'[encryption XOR <sup>(EN)</sup>](https://en.wikipedia.org/wiki/XOR_cipher) (une autre ). Il est codé en Golang, et est compilé pour plusieurs OS (Windows&reg;, Linux&reg;,et bientôt macOS&reg;).

Je ne suis pas un pro en Go, donc si vous trouvez un bug, ou simplement voulez faire une suggestion, je suis totalement ouvert aux [_issues_ <sup>(EN)</sup>](https://github.com/eliotttak/GoFileEncoder/issues), [_pull requests_ <sup>(EN)</sup>](https://github.com/eliotttak/GoFileEncoder/pulls) et à la [discussion <sup>(EN)</sup>](https://github.com/eliotttak/GoFileEncoder/discussions).

---

## Sommaire
- [GoFileEncoder](#gofileencoder)
  - [Introduction](#introduction)
  - [Sommaire](#sommaire)
  - [Construction](#construction)
    - [Clonage](#clonage)
    - [Arbre des dépendances](#arbre-des-dépendances)
    - [Installation des dépendances](#installation-des-dépendances)
    - [Création du _package_ d'_assets_](#création-du-package-dassets)
    - [Compilation](#compilation)
      - [En Bash](#en-bash)
      - [En Batch](#en-batch)
  - [Utilisation](#utilisation)
    - [1. Encoder ou décoder ?](#1-encoder-ou-décoder-)
    - [2. Sélectionnez un fichier](#2-sélectionnez-un-fichier)
    - [3. Entrez le mot de passe](#3-entrez-le-mot-de-passe)
    - [4. Choisissez le fichier de destination](#4-choisissez-le-fichier-de-destination)
    - [5. Attendez...](#5-attendez)
    - [6. Terminé !](#6-terminé-)
  - [License](#license)
  - [Introduction](#introduction-1)
  - [Sommaire](#sommaire-1)
  - [Construction](#construction-1)
    - [Clonage](#clonage-1)
    - [Arbre des dépendances](#arbre-des-dépendances-1)
    - [Installation des dépendances](#installation-des-dépendances-1)
    - [Compilation](#compilation-1)
      - [En Bash](#en-bash-1)
      - [En Batch](#en-batch-1)
  - [Utilisation](#utilisation-1)
    - [1. Encoder ou décoder ?](#1-encoder-ou-décoder--1)
    - [2. Sélectionnez un fichier](#2-sélectionnez-un-fichier-1)
    - [3. Entrez le mot de passe](#3-entrez-le-mot-de-passe-1)
    - [4. Choisissez le fichier de destination](#4-choisissez-le-fichier-de-destination-1)
    - [5. Attendez...](#5-attendez-1)
    - [6. Terminé !](#6-terminé--1)
  - [License](#license-1)


## Construction

### Clonage

```bash
git clone https://github.com/eliotttak/GoFileEncoder
cd GoFileEncoder
```

### Arbre des dépendances

```plaintext
GoFileEncoder (ce projet)
|
+-- github.com/sqweek/dialog (pour les popups des fichiers)
|   |
|    \_ github.com/TheTitanrain/w32 (indirect)
|
 \_ golang.org/x/term (pour la demande de mot de passe)
    |
     \_ golang.org/x/sys (indirect)
```

Se réferer au fichier [go.mod](./go.mod) pour plus de détails.

### Installation des dépendances
```bash
go mod tidy
```

### Compilation

#### En Bash
```bash
go build -o ./bin/GoFileEncoder ./src/
```

#### En Batch
```batch
go build -o .\bin\GoFileEncoder .\src\
```

---

## Utilisation

### 1. Encoder ou décoder ?

En premier, lancez le programme. Celui-ci va vous demander si vous souhaitez encoder ou décoder un fichier :
```plaintext
Que voulez-vous faire ?
 - Encoder un fichier (e)
 - Décoder un fichier (d)
(e/d)>>>
```
Si vous voulez encoder un fichier, entrez <kbd>e</kbd>, sinon, entrez <kbd>d</kbd>. Dans tous les cas, validez avec <kbd>Entrée</kbd>.

### 2. Sélectionnez un fichier

Ensuite, il va vous demander d'appuyez sur <kbd>Entrée</kbd> pour sélectionner un fichier. Faites-le, et un popup va apparaître. Sélectionnez votre fichier, puis validez.
> &#x1F6C8; Si vous annulez, le popup va réapparaître 2 fois, puis à la 3<sup>ème</sup>, le message `"Trop de tentatives échouées"` va s'afficher puis le programme va se fermer.

### 3. Entrez le mot de passe

Après cela, un mot de passe vous sera demandé. Entrez-le puis validez avec <kbd>Entrée</kbd>.
> &#x26A0; Si vous avez choisi `Décoder un fichier`, faîtes attention à bien écrire le mot de passe. Si vous vous trompez ne serait-ce que d'un caractère, le fichier sera corrompu et ne sera plus utilisable.

> &#x1F6C8; Pour garantir la confidentialité de votre mot de passe, celui-ci ne s'affichera pas sur votre écran.

### 4. Choisissez le fichier de destination

Ensuite, le programme va vous demander d'appuyez sur <kbd>Entrée</kbd> pour choisir le fichier de destination. Faites-le, et un popup va apparaître. Sélectionnez le fichier à créer, puis validez.
> &#x1F6C8; Si vous annulez, le popup va réapparaître 2 fois, puis à la 3<sup>ème</sup>, le message `"Trop de tentatives échouées"` va s'afficher puis le programme va se fermer.

### 5. Attendez...

Le fichier est en train d'être encodé. Ne fermez pas le programme.

### 6. Terminé !
Au bout de quelques secondes, le fichier est encodé ou décodé. Le programme se ferme automatiquement.

## License
Ce logiciel est distribué sous la license GNU GENERAL PUBLIC LICENSE version 3 (GNU GPL v3).
[Voir la license <sup>(EN)</sup>](LICENSE))