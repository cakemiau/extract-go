package extract

import (
	"os"
	"testing"
)

func ExtractTest(t *testing.T, url string) {

	SetVerbose(true)

	dir, err := os.MkdirTemp("", "go-extract-test-*")
	if err != nil {
		t.Fatal(err)
	}
	//defer os.RemoveAll(dir)

	err = Url(url, dir, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = Url(url, dir, 1)
	if err != nil {
		t.Fatal(err)
	}

	err = Url(url, dir, 2000)
	if err != nil {
		t.Fatal(err)
	}

	err = Url(url, dir, 100000000)
	if err != nil {
		t.Fatal(err)
	}
}

func TestZipOnlyDisk(t *testing.T) {
	ExtractTest(t, "https://api.adoptopenjdk.net/v3/binary/latest/17/ga/windows/x64/jre/hotspot/normal/adoptopenjdk")
}

func TestZipTooBig(t *testing.T) {

}

func TestZipAlwaysMem(t *testing.T) {

}

func TestGzip(t *testing.T) {
	ExtractTest(t, "https://api.adoptopenjdk.net/v3/binary/latest/17/ga/linux/x64/jre/hotspot/normal/adoptopenjdk")
}

func TestZip2(t *testing.T) {
	ExtractTest(t, "https://storage.googleapis.com/chrome-for-testing-public/130.0.6723.58/win64/chrome-win64.zip")
}
