package main

import (
	"fmt"
	"os"

	"github.com/cakemiau/extract-go/extract"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Print("Usage: <url> <path>\n  url: Zip to download and extract\n  path: Output directory\n")
		os.Exit(1)
	}

	err := extract.ExtractFromUrl(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
	}

}
