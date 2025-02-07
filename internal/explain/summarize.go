package explain

import (
	"github.com/jmendick/echodiff/internal/diff"
	"fmt"
	"strings"
)

// GenerateSummary generates a human-readable explanation of file differences.
func GenerateSummary(file1, file2 string) (string, error) {
	diffResult, err := diff.CompareFiles(file1, file2)
	if err != nil {
		return "", err
	}

	summary := summarizeDiff(diffResult)
	return summary, nil
}

// summarizeDiff converts a raw diff into human-friendly descriptions.
func summarizeDiff(diffResult string) string {
	lines := strings.Split(diffResult, "\n")
	var explanation []string

	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			explanation = append(explanation, fmt.Sprintf("Removed: %s", line[1:]))
		} else if strings.HasPrefix(line, "+") {
			explanation = append(explanation, fmt.Sprintf("Added: %s", line[1:]))
		}
	}

	if len(explanation) == 0 {
		return "No significant changes detected."
	}

	return strings.Join(explanation, "\n")
}
