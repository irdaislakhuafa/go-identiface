# Go Identiface

A simple library for identifying face with golang

## Requirements

This library is based on [`go-face`](https://github.com/Kagami/go-face.git) but with support for ID data type customization. So, this library require `go-face` to be used and `go-face` require `dlib` >= `19.10` and `libjpeg` development packages installed to run.

Quoting from [`here`](https://github.com/Kagami/go-face/blob/master/README.md) for `go-face` requirement is below.

### Ubuntu 18.10+, Debian sid

Latest versions of Ubuntu and Debian provide suitable dlib package so just run:

```bash
# Ubuntu
sudo apt-get install libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg-turbo8-dev
# Debian
sudo apt-get install libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg62-turbo-dev
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

If you use [`Docker`](https://www.docker.com/). I have created a sample `Dockerfile` for you [`here`](./Dockerfile).

## Usage

To use `go-identiface` to your Go code, use it as dependency below:

```bash
go get github.com/irdaislakhuafa/go-identiface.git
```

Then import it on your Go project:

```go
import "github.com/irdaislakhuafa/go-identiface"
```

## Example

NOTE:

- Make sure you have downloaded the sample assets with command `make get-assets`. This command required [`Makefile`](https://www.gnu.org/software/make) package.

You can see example usage of code [`here`](./example/main.go)

<!-- ## Testing -->
