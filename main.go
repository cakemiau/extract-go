package main

import (
	"fmt"
	"os"

	"github.com/cakemiau/extract-go/extract"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Print("Usage: <url> <path>\n  url: Zip/Gzip/Tar to download and extract\n  path: Output directory\n")
		os.Exit(1)
	}

	err := extract.Url(os.Args[1], os.Args[2], 0)
	if err != nil {
		fmt.Println(err)
	}

}
