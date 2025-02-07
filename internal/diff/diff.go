package diff

import (
	"os"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// CompareFiles compares two files and returns the diff result.
func CompareFiles(file1, file2 string) (string, error) {
	content1, err := readFile(file1)
	if err != nil {
		return "", err
	}

	content2, err := readFile(file2)
	if err != nil {
		return "", err
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(content1), string(content2), false)
	return dmp.DiffPrettyText(diffs), nil
}

// readFile reads a file's content and returns it as a byte slice.
func readFile(filepath string) ([]byte, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
