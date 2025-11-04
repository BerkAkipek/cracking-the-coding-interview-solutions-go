package sortbigfile

/*
Sort Big File: Imagine you have a 20 GB file with one string per line. Explain how you would sort
the file.
*/

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

// entry represents a current line from one chunk file
type entry struct {
	line string
	src  int
}

// MinHeap implements a priority queue for entries
type MinHeap []entry

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].line < h[j].line }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(entry)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// SplitAndSortBigFile splits the input file into sorted chunks.
// chunkSize controls how many lines per chunk fit into memory.
func SplitAndSortBigFile(inputPath string, chunkSize int) ([]string, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0, chunkSize)
	var chunkFiles []string
	chunkIndex := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) >= chunkSize {
			fileName := fmt.Sprintf("chunk_%03d.tmp", chunkIndex)
			if err := writeSortedChunk(fileName, lines); err != nil {
				return nil, err
			}
			chunkFiles = append(chunkFiles, fileName)
			lines = lines[:0]
			chunkIndex++
		}
	}
	if len(lines) > 0 {
		fileName := fmt.Sprintf("chunk_%03d.tmp", chunkIndex)
		if err := writeSortedChunk(fileName, lines); err != nil {
			return nil, err
		}
		chunkFiles = append(chunkFiles, fileName)
	}
	return chunkFiles, nil
}

// helper: sorts lines in-memory and writes to disk
func writeSortedChunk(filename string, lines []string) error {
	slices.Sort(lines)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	for _, l := range lines {
		_, _ = writer.WriteString(l + "\n")
	}
	return writer.Flush()
}

// MergeChunks merges sorted temporary chunk files into one sorted output file.
func MergeChunks(chunkFiles []string, outputPath string) error {
	readers := make([]*bufio.Reader, len(chunkFiles))
	files := make([]*os.File, len(chunkFiles))

	for i, path := range chunkFiles {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		files[i] = f
		readers[i] = bufio.NewReader(f)
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()
	writer := bufio.NewWriter(out)

	h := &MinHeap{}
	heap.Init(h)

	// Initialize heap with the first line from each file
	for i, r := range readers {
		line, err := r.ReadString('\n')
		if err == nil {
			heap.Push(h, entry{line: line, src: i})
		}
	}

	for h.Len() > 0 {
		smallest := heap.Pop(h).(entry)
		writer.WriteString(smallest.line)

		next, err := readers[smallest.src].ReadString('\n')
		if err == nil {
			heap.Push(h, entry{line: next, src: smallest.src})
		}
	}

	writer.Flush()
	for _, f := range files {
		f.Close()
	}
	return nil
}
