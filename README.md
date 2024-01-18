# Go Identiface

A simple library for identifying face with golang

## Requirements

This library is based on [`go-face`](https://github.com/Kagami/go-face.git) but with support for ID data type customization. So, this library require `go-face` to be used and `go-face` require `dlib` >= `19.10` and `libjpeg` development packages installed to run.

Quoting from [`here`](https://github.com/Kagami/go-face/blob/master/README.md) for `go-face` requirement is below.

### Ubuntu 18.10+, Debian sid

Latest versions of Ubuntu and Debian provide suitable dlib package so just run:

```bash
# Ubuntu
sudo apt-get install libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg-turbo8-dev --no-install-recommends
# Debian
sudo apt-get install libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg62-turbo-dev --no-install-recommends
```

[`go-face`](https://github.com/Kagami/go-face.git) used `C` and `C++` language for some case. So, you need to install `C` and `C++` compiler to use this project as dependency (just for compile not for run), if you have no `C` and `C++` compiler you will get error `"face.Face is undefined"` if you want to compile your golang project. Makesure you have `C` and `C++` compiler installed but if you don't have any `C` and `C++` compiler. Just install it with following command below:

```bash
sudo apt-get install g++ gcc --no-install-recommends
```

### macOS

Make sure you have [Homebrew](https://brew.sh) installed.

```bash
brew install dlib
```

### Windows

Make sure you have [MSYS2](https://www.msys2.org) installed.

1. Run `MSYS2 MSYS` shell from Start menu
2. Run `pacman -Syu` and if it asks you to close the shell do that
3. Run `pacman -Syu` again
4. Run `pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-dlib`
5. Set env variables

   - If you already have Go and Git installed and available in PATH uncomment
     `set MSYS2_PATH_TYPE=inherit` line in `msys2_shell.cmd` located in MSYS2
     installation folder

   - Otherwise run `pacman -S mingw-w64-x86_64-go git`

6. Run `MSYS2 MinGW 64-bit` shell from Start menu to compile and use go-face

### Other systems

Try to install dlib/libjpeg with package manager of your distribution or
[compile from sources](http://dlib.net/compile.html). Note that go-face won't
work with old packages of dlib such as libdlib18. Alternatively create issue
with the name of your system and someone might help you with the installation
process.

### Docker

If you use [`Docker`](https://www.docker.com/). I have created a sample `Dockerfile` for you [`here`](./Dockerfile) and you can customize according you need.

## Usage

To use `go-identiface` to your Go code, use it as dependency below:

```bash
go get github.com/irdaislakhuafa/go-identiface@latest
```

Then import it on your Go project:

```go
import "github.com/irdaislakhuafa/go-identiface"
```

## Example

NOTE:

- Make sure you have installed all dependencies above and downloaded the sample assets with command `make get-assets`. This command required [`Makefile`](https://www.gnu.org/software/make) package.

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/go-identiface"
)

type user struct {
	id        string `json:"id,omitempty"`
	name      string `json:"name,omitempty"`
	imagePath string `json:"image_name,omitempty"`
}

var (
	assetsDir = "./assets"
	modelsDir = filepath.Join(assetsDir, "models")
	imagesDir = filepath.Join(assetsDir, "images")
)

// Run this code at root dir of project with `go run ./example/main.go`
func main() {
	// Imagine this is data users from database
	users := []user{
		{id: uuid.NewString(), name: "Tony Stark", imagePath: filepath.Join(imagesDir, "tony-stark.jpeg")},
		{id: uuid.NewString(), name: "Tzuyu", imagePath: filepath.Join(imagesDir, "tzuyu.jpg")},
	}

	getUserByID := map[string]user{}
	for _, user := range users {
		getUserByID[user.id] = user
	}

	// Initialize `identiface` [string] is data type of key for each image
	iFace, err := identiface.Init[string](modelsDir)
	if err != nil {
		panic(err)
	}
	defer iFace.Close()

	// Adding image datasets to `go-identiface`
	for _, user := range users {
		fileBytes, err := os.ReadFile(user.imagePath)
		if err != nil {
			panic(err)
		}

		if err := iFace.AddSingleDatasetFromBytes(user.id, fileBytes); err != nil {
			panic(err)
		}
	}

	// After adding datasets, don't forget to load them so they can be used by `go-identiface`. After this you can identify user by human face of image
	iFace.LoadDatasets()

	// Imagine this is an image file sent from client (ex. Mobile/Frontend/Etc)
	clientImg := "tony-stark3.jpeg"
	clienImgBytes, err := os.ReadFile(filepath.Join(imagesDir, clientImg))
	if err != nil {
		panic(err)
	}

	// After get image file from client then you can identify client by face of the image
	data, err := iFace.ClassifySingleFromBytes(clienImgBytes)
	if err != nil { // Will return error if face is not recognized or image contain multiple face
		panic(err)
	}

	// Imagine this is method to get detail data user from db
	user := getUserByID[data.ID]

	// Now you can identify user by face
	fmt.Printf("user: %#v\n", user)
}
```

You can see detail example usage of code [`here`](./example/main.go)

## Testing

To run testing code at `identiface_test.go` if you have installed all dependencies above depend of your Operating System you can run `go test -v .../..` directly in root folder. But if you don't want to pollute your system with temporary dependencies to test code use the command below or you can modify [`test.Dockerfile`](./test.Dockerfile) according you need.

```bash
docker compose up test
```

## Todo

- Handle identify for grey image
