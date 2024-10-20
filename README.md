# Zip (Go)

An utility package to easily extract a Zip file, omitting the shared root folder if present.

This package contains two main functions:

- `zip.ExtractAll`: Extracts all the contents of the zip.Reader into the given path, omitting the shared root folder if present.
```go
func zip.ExtractAll(zipReader *zip.Reader, outputPath string) error { ... }
```

- `ExtractFromUrl`: Makes an HTTP GET request, loads the Zip file into **memory**, and extracts it's contents into the given path,
  omitting the shared root folder if present.
```go
func zip.ExtractFromUrl(sourceUrl string, outputPath string) error { ... }
```

> [!WARNING]
> Don't use `zip.ExtractFromUrl` for big zip files, as your system may run out of memory.<br>
> Download the file to disk and use the `zip.ExtractAll` function instead.

## Usage

First, download the package
```sh
go get github.com/cakemiau/zip-go
```

Extract a file on disk
```go
import (
	zipArchive "archive/zip"
	"github.com/cakemiau/zip-go/zip"
)

func main() {
	file, err := zipArchive.OpenReader("<SOURCE PATH>")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = zip.ExtractAll(&file.Reader, "<OUTPUT PATH>")
	if err != nil {
		panic(err)
	}
}
```

Extract a file from a URL
```go
import "github.com/cakemiau/zip-go/zip"

func main() {
	err = zip.ExtractFromUrl("<INPUT URL>", "<OUTPUT PATH>")
	if err != nil {
		panic(err)
	}
}
```
