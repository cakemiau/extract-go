package extract

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Zip(zipReader *zip.Reader, outputPath string) error {

	prefixPathLen := len(pathPrefixZipList(zipReader.File))

	for _, zipFile := range zipReader.File {

		verbose(zipFile.Name)

		fileName := filepath.Join(outputPath, zipFile.Name[prefixPathLen:])

		if zipFile.FileInfo().IsDir() {
			if err := os.MkdirAll(fileName, 0755); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
			return err
		}

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

func Tar(tarReader *tar.Reader, outputPath string) error {

	// Tar files are read sequentally, so we have to extract all files to
	// a temporary directory and then move them to remove the path prefix.

	tempDir, err := os.MkdirTemp("", "extract-go-output-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	fileNames := []string{}

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		verbose(header.Name)

		fileName := filepath.Join(tempDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(fileName, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			f, err := os.Create(fileName)
			if err != nil {
				return err
			}
			defer f.Close()
			if _, err := io.Copy(f, tarReader); err != nil {
				return err
			}
			f.Close()
			fileNames = append(fileNames, header.Name)
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, fileName); err != nil {
				return err
			}
			fileNames = append(fileNames, header.Name)
		default:
			return fmt.Errorf("unsupported type %c for %s", header.Typeflag, header.Name)
		}
	}

	prefixPathLen := len(pathPrefixStrList(fileNames))

	for _, fileName := range fileNames {
		outputFilePath := filepath.Join(outputPath, fileName[prefixPathLen:])
		if err := os.MkdirAll(filepath.Dir(outputFilePath), 0755); err != nil {
			return err
		}
		if err := os.Rename(filepath.Join(tempDir, fileName), outputFilePath); err != nil {
			return err
		}
	}

	return nil
}

func Url(sourceUrl string, outputPath string, maxMemorySize int) error {

	res, err := http.Get(sourceUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http response code %d", res.StatusCode)
	}

	mimeType := res.Header.Get("content-type")
	fileName := filepath.Base(res.Request.URL.Path)
	if _, params, err := mime.ParseMediaType(res.Header.Get("content-disposition")); err == nil {
		value, exists := params["filename"]
		if exists {
			fileName = value
		}
	}

	// Create a buffer to hold the initial part of the response
	buff := make([]byte, maxMemorySize)
	bytesRead, err := io.ReadFull(res.Body, buff)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return err
	}

	var rawStream io.ReaderAt
	var rawSize int64

	if bytesRead < maxMemorySize {
		// file fits in the defined max memory size
		rawStream = bytes.NewReader(buff)
		rawSize = int64(bytesRead)

	} else {
		// file does not fit in the defined max memory size

		file, err := os.CreateTemp("", "extract-go-input-")
		if err != nil {
			return err
		}
		defer file.Close()

		// write the data in the buffer to the file
		bytesReadBuffer, err := file.Write(buff)
		if err != nil {
			return err
		}

		// copy the remaining data from the response body to the file
		bytesReadRemaining, err := io.Copy(file, res.Body)
		if err != nil {
			return err
		}
		res.Body.Close()

		rawStream = file
		rawSize = int64(bytesReadBuffer) + bytesReadRemaining
	}

	if mimeType == "application/zip" || mimeType == "application/x-zip-compressed" || strings.HasSuffix(fileName, ".zip") {

		zipReader, err := zip.NewReader(rawStream, rawSize)
		if err != nil {
			return err
		}
		return Zip(zipReader, outputPath)

	} else if mimeType == "application/gzip" || mimeType == "application/x-gzip" || strings.HasSuffix(fileName, ".tar.gz") {

		ungzippedStream, err := gzip.NewReader(io.NewSectionReader(rawStream, 0, rawSize))
		if err != nil {
			return err
		}
		tarReader := tar.NewReader(ungzippedStream)
		return Tar(tarReader, outputPath)

	} else if mimeType == "application/x-tar" || strings.HasSuffix(fileName, ".tar") {

		tarReader := tar.NewReader(io.NewSectionReader(rawStream, 0, rawSize))
		return Tar(tarReader, outputPath)

	}

	return fmt.Errorf("unsupported or undefined file format (Type: %s | Name: %s)", mimeType, fileName)
}
