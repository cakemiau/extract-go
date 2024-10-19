package memoryZip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sort"
)

// Finds the shared directory of two paths, assuming all directories have a trailing forward slash (/)
func CommonPath(s1 string, s2 string) string {
	max_len := min(len(s1), len(s2))

	i := 0
	end := 0
	for i < max_len && s1[i] == s2[i] {
		if s1[i] == '/' {
			end = i
		}
		i++
	}

	if end == 0 {
		return ""
	}

	return s1[:end+1]
}

// Finds the shared directory of a list of paths, assuming all directories have a trailing forward slash (/)
func CommonPathStrList(strList []string) string {
	n := len(strList)
	if n <= 0 {
		return ""
	}

	slices.Sort(strList)

	first := strList[0]
	last := strList[n-1]

	return CommonPath(first, last)
}

// Finds the shared directory of a list of zipFiles, assuming all directories have a trailing forward slash (/)
func CommonPathZipFile(zipFiles []*zip.File) string {
	n := len(zipFiles)
	if n <= 0 {
		return ""
	}

	sort.Slice(zipFiles, func(i, j int) bool {
		return zipFiles[i].Name < zipFiles[j].Name
	})

	first := zipFiles[0].Name
	last := zipFiles[n-1].Name

	return CommonPath(first, last)
}

// Extracts all the contents of the Zip file, skipping the shared root folder if possible.
func Extract(zipReader *zip.Reader, outputPath string) error {

	shared_path_len := len(CommonPathZipFile(zipReader.File))

	for _, zipFile := range zipReader.File {

		fmt.Println(zipFile.Name)

		fileName := filepath.Join(outputPath, zipFile.Name[shared_path_len:])

		if zipFile.FileInfo().IsDir() {
			os.MkdirAll(fileName, 0777)
			continue
		}

		os.MkdirAll(filepath.Dir(fileName), 0777)

		z, err := zipFile.Open()
		if err != nil {
			return err
		}
		defer z.Close()

		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, z)
		z.Close()
		f.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// Makes a GET request for a Zip file and extracts it, skipping the shared root folder if possible.
func ExtractGet(sourceUrl string, outputPath string) error {

	res, err := http.Get(sourceUrl)
	if err != nil {
		return err
	}

	var zipReader *zip.Reader
	var byteReader *bytes.Reader
	var zipBytes []byte

	defer func() {
		zipReader = nil
		byteReader = nil
		zipBytes = nil
	}()

	zipBytes, err = io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	byteReader = bytes.NewReader(zipBytes)
	zipReader, err = zip.NewReader(byteReader, int64(len(zipBytes)))
	if err != nil {
		return err
	}

	return Extract(zipReader, outputPath)
}
