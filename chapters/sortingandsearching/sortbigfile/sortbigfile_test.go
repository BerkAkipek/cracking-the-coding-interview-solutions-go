package sortbigfile

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func TestExternalSort(t *testing.T) {
	os.Remove("input.txt")
	os.Remove("sorted.txt")

	data := []string{
		"orange", "apple", "banana", "pear", "kiwi",
		"grape", "melon", "avocado", "plum", "mango", "cherry",
	}
	err := os.WriteFile("input.txt", []byte(strings.Join(data, "\n")), 0644)
	if err != nil {
		t.Fatalf("failed to write input file: %v", err)
	}

	chunks, err := SplitAndSortBigFile("input.txt", 3)
	if err != nil {
		t.Fatalf("split failed: %v", err)
	}

	err = MergeChunks(chunks, "sorted.txt")
	if err != nil {
		t.Fatalf("merge failed: %v", err)
	}

	out, err := os.ReadFile("sorted.txt")
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}
	lines := strings.Fields(string(out))

	if !slices.IsSorted(lines) {
		t.Fatalf("output not sorted: %v", lines)
	}

	for _, chunk := range chunks {
		_ = os.Remove(chunk)
	}

	t.Logf("sorted output: %v", lines)
}
