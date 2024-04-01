package pkg

import (
	"strings"
)

func CompPath(specPath, testPath string) bool {
	specPathStr := strings.Split(specPath, "/")
	testPathStr := strings.Split(testPath, "/")

	if len(specPathStr) != len(testPathStr) {
		return false
	}
	for i, p := range specPathStr {
		if i == 0 {
			continue
		}
		if strings.Compare(p, testPathStr[i]) != 0 && !strings.Contains(p, "{") {
			return false
		}
	}
	return true
}

func InArray(str string, a []string) bool {
	for _, s := range a {
		if strings.Compare(str, s) == 0 {
			return true
		}
	}
	return false
}
