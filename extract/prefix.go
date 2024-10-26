package extract

import (
	"archive/zip"
	"slices"
	"sort"
)

func pathPrefix(s1 string, s2 string) string {
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

func pathPrefixStrList(strList []string) string {
	n := len(strList)
	if n <= 0 {
		return ""
	}

	slices.Sort(strList)

	first := strList[0]
	last := strList[n-1]

	return pathPrefix(first, last)
}

func pathPrefixZipList(zipFiles []*zip.File) string {
	n := len(zipFiles)
	if n <= 0 {
		return ""
	}

	sort.Slice(zipFiles, func(i, j int) bool {
		return zipFiles[i].Name < zipFiles[j].Name
	})

	first := zipFiles[0].Name
	last := zipFiles[n-1].Name

	return pathPrefix(first, last)
}
