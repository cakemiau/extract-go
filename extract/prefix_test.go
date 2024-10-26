package extract

import "testing"

func commonPathTest(t *testing.T, expected string, input []string) {
	result := pathPrefixStrList(input)
	if result != expected {
		t.Fatalf("Expected %s got %s", expected, result)
	}
}

func TestPath1(t *testing.T) {
	commonPathTest(t, "root/", []string{
		"root/child1",
		"root/child2",
		"root/child3",
		"root/child4",
	})
}

func TestPath2(t *testing.T) {
	commonPathTest(t, "", []string{
		"root/child1",
		"root/child2",
		"file.txt",
		"root/child3",
		"root/child4",
	})
}

func TestPath3(t *testing.T) {
	commonPathTest(t, "", []string{
		"root/child1",
		"root/child2",
		"root",
		"root/child3",
		"root/child4",
	})
}
