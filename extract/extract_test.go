package extract

import (
	"math"
	"os"
	"testing"
)

// no bytes are copied to memory, the zip archive is always written to disk.
func TestZipOnlyDisk(t *testing.T) {
	SetVerbose(true)

	dir, err := os.MkdirTemp("", "go-extract-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	if err := Url("https://api.adoptopenjdk.net/v3/binary/latest/17/ga/windows/x64/jre/hotspot/normal/adoptopenjdk", dir, 0); err != nil {
		t.Fatal(err)
	}
}

// the zip archive is initially copied to memory, but it does not fit and is moved to disk.
func TestZipTooBig(t *testing.T) {
	SetVerbose(true)

	dir, err := os.MkdirTemp("", "go-extract-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	if err := Url("https://api.adoptopenjdk.net/v3/binary/latest/17/ga/windows/x64/jre/hotspot/normal/adoptopenjdk", dir, 5000); err != nil {
		t.Fatal(err)
	}
}

// the zip archive is always fully copied into memory.
func TestZipAlwaysMem(t *testing.T) {
	SetVerbose(true)

	dir, err := os.MkdirTemp("", "go-extract-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	if err := Url("https://api.adoptopenjdk.net/v3/binary/latest/17/ga/windows/x64/jre/hotspot/normal/adoptopenjdk", dir, math.MaxInt32); err != nil {
		t.Fatal(err)
	}
}

// no bytes are copied to memory, the tar archive is always written to disk.
func TestGZipOnlyDisk(t *testing.T) {
	SetVerbose(true)

	dir, err := os.MkdirTemp("", "go-extract-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	if err := Url("https://api.adoptopenjdk.net/v3/binary/latest/17/ga/linux/x64/jre/hotspot/normal/adoptopenjdk", dir, 0); err != nil {
		t.Fatal(err)
	}
}

// the tar archive is initially copied to memory, but it does not fit and is moved to disk.
func TestGZipTooBig(t *testing.T) {
	SetVerbose(true)

	dir, err := os.MkdirTemp("", "go-extract-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	if err := Url("https://api.adoptopenjdk.net/v3/binary/latest/17/ga/linux/x64/jre/hotspot/normal/adoptopenjdk", dir, 5000); err != nil {
		t.Fatal(err)
	}
}

// the tar archive is always fully copied into memory.
func TestGZipAlwaysMem(t *testing.T) {
	SetVerbose(true)

	dir, err := os.MkdirTemp("", "go-extract-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	if err := Url("https://api.adoptopenjdk.net/v3/binary/latest/17/ga/linux/x64/jre/hotspot/normal/adoptopenjdk", dir, math.MaxInt32); err != nil {
		t.Fatal(err)
	}
}

// alternative test for detecting archive formats
func TestZipAlt(t *testing.T) {
	SetVerbose(true)

	dir, err := os.MkdirTemp("", "go-extract-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	if err := Url("https://storage.googleapis.com/chrome-for-testing-public/130.0.6723.58/win64/chrome-win64.zip", dir, math.MaxInt32); err != nil {
		t.Fatal(err)
	}
}
