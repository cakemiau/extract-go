package main

import (
	"fmt"
	"os"

	"github.com/cakemiau/zip-go/zip"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Print("Usage: <url> <path>\n  url: Zip to download and extract\n  path: Output directory\n")
		os.Exit(1)
	}

	err := zip.ExtractFromUrl(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
	}

}
