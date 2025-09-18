# GoFileEncoder

> &#x1F1EC;&#x1F1E7; Do you speak English ? Open [README.md](README.md) in English.

## Introduction

GoFileEncoder est un petit encodeur de fichiers qui utilise l'[encryption XOR <sup>(EN)</sup>](https://en.wikipedia.org/wiki/XOR_cipher). Il est codé en Golang, et est compilé pour plusieurs OS (Windows&reg;, Linux&reg;,et bientôt macOS&reg;).

Je ne suis pas un pro en Go, donc si vous trouvez un bug, ou simplement voulez faire une suggestion, je suis totalement ouvert aux [_issues_ <sup>(EN)</sup>](https://github.com/eliotttak/GoFileEncoder/issues), [_pull requests_ <sup>(EN)</sup>](https://github.com/eliotttak/GoFileEncoder/pulls) et à la [discussion <sup>(EN)</sup>](https://github.com/eliotttak/GoFileEncoder/discussions).

---

## Sommaire
- [GoFileEncoder](#gofileencoder)
  - [Introduction](#introduction)
  - [Sommaire](#sommaire)
  - [Construction](#construction)
    - [Clonage](#clonage)
    - [Arbre des dépendances du _package_](#arbre-des-dépendances-du-package)
    - [Dépendances de construction](#dépendances-de-construction)
    - [Installation des dépendances](#installation-des-dépendances)
    - [Compilation et création du _package_ d'assets](#compilation-et-création-du-package-dassets)
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


## Construction

### Clonage

```bash
git clone https://github.com/eliotttak/GoFileEncoder
cd GoFileEncoder
```

### Arbre des dépendances du _package_

```plaintext
GoFileEncoder (ce projet)
|
+-- github.com/sqweek/dialog (pour les popups des fichiers)
|   |
|    \_ github.com/TheTitanrain/w32 (indirect)
|
+-- golang.org/x/term (pour la demande de mot de passe)
|   |
|    \_ golang.org/x/sys (indirect)
|
 \_ github.com/christianhujer/isheadless (pour vérifier s'il y a une GUI)
```

Se réferer au fichier [go.mod](./go.mod) pour plus de détails.
> &#x1F6C8; Dans [go.mod](./go.mod), il peut y avoir une directive d'import pour [github.com/abdfnx/gosh](https://github.com/abdfnx/gosh). Elle a été utilisée uniquement pour mon assistant d'installation, et vous ne devriez pas en avoir besoin. Elle va disparaître si vous exécutez `go mod tidy`.

### Dépendances de construction

- `github.com/go-bindata/go-bindata/go-bindata/...` (pour créer le fichier d'_assets_)

### Installation des dépendances
```bash
go mod tidy
go install github.com/go-bindata/go-bindata/go-bindata/...
go get
```
### Compilation et création du _package_ d'assets

#### En Bash
```bash
# Vous allez peut-être devoir exécuter 'chmod 744 build.sh'
./build.sh # Vous pouvez ajouter une valeur de GOOS et une de GOARCH, par ex. './build.sh linux amd64'
```

#### En Batch
```batch
rem Vous pouvez ajouter une valeur de GOOS et une de GOARCH, par ex. 'build.bat windows amd64'
build.bat
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

> &#x1F6C8; Si votre configuration ne comprends pas d'interface graphique, vous allez devoir entrer manuellement le chemin absolu vers vorte fichier.

### 3. Entrez le mot de passe

Après cela, un mot de passe vous sera demandé. Entrez-le puis validez avec <kbd>Entrée</kbd>.
> &#x26A0; Si vous avez choisi `Décoder un fichier`, faîtes attention à bien écrire le mot de passe. Si vous vous trompez ne serait-ce que d'un caractère, le fichier sera corrompu et ne sera plus utilisable.

> &#x1F6C8; Pour garantir la confidentialité de votre mot de passe, celui-ci ne s'affichera pas sur votre écran.

### 4. Choisissez le fichier de destination

Ensuite, le programme va vous demander d'appuyez sur <kbd>Entrée</kbd> pour choisir le fichier de destination. Faites-le, et un popup va apparaître. Sélectionnez le fichier à créer, puis validez.
> &#x1F6C8; Si vous annulez, le popup va réapparaître 2 fois, puis à la 3<sup>ème</sup>, le message `"Trop de tentatives échouées"` va s'afficher puis le programme va se fermer.

> &#x1F6C8; Si votre configuration ne comprends pas d'interface graphique, vous allez devoir entrer manuellement le chemin absolu vers vorte fichier.


### 5. Attendez...

Le fichier est en train d'être encodé. Ne fermez pas le programme.

### 6. Terminé !
Au bout de quelques secondes, le fichier est encodé ou décodé. Le programme se ferme automatiquement.

## License
Ce logiciel est distribué sous la license GNU GENERAL PUBLIC LICENSE version 3 (GNU GPL v3).

[Voir la license <sup>(EN)</sup>](LICENSE)

