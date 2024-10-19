package zip

import (
	"os"
	"testing"
)

func CommonPathTest(t *testing.T, expected string, input []string) {
	result := CommonPathStrList(input)
	if result != expected {
		t.Fatalf("Expected %s got %s", expected, result)
	}
}

func TestPath1(t *testing.T) {
	CommonPathTest(t, "root/", []string{
		"root/child1",
		"root/child2",
		"root/child3",
		"root/child4",
	})
}

func TestPath2(t *testing.T) {
	CommonPathTest(t, "", []string{
		"root/child1",
		"root/child2",
		"file.txt",
		"root/child3",
		"root/child4",
	})
}

func TestPath3(t *testing.T) {
	CommonPathTest(t, "", []string{
		"root/child1",
		"root/child2",
		"root",
		"root/child3",
		"root/child4",
	})
}

func ZipTest(t *testing.T, url string) {
	dir, err := os.MkdirTemp("", "zip-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	err = ExtractFromUrl(url, dir)
	if err != nil {
		t.Fatal(err)
	}
}

func TestZip1(t *testing.T) {
	ZipTest(t, "https://api.adoptopenjdk.net/v3/binary/latest/17/ga/windows/x64/jre/hotspot/normal/adoptopenjdk")
}

func TestZip2(t *testing.T) {
	ZipTest(t, "https://storage.googleapis.com/chrome-for-testing-public/130.0.6723.58/win64/chrome-win64.zip")
}
