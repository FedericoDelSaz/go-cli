package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MinHeap []uint64

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(uint64))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// Define flags
	n := flag.Int("n", 1, "Number of top results.")
	inputFile := flag.String("input-file", "", "Path to input file.")
	outputFile := flag.String("output-file", "", "Path to output file.")
	chunkSize := flag.Int("chunk-size", 1000000, "Numbers per chunk for sorting.")
	flag.Parse()

	// Validate flags
	if *inputFile == "" || *outputFile == "" {
		fmt.Println("ERROR: Input and output file paths are required.")
		os.Exit(1)
	}
	if *n <= 0 {
		fmt.Println("ERROR: Number of top results must be greater than 0.")
		os.Exit(1)
	}

	// Step 1: Read file in chunks, sort, and write to temp files
	tempFiles, err := sortChunks(*inputFile, *chunkSize)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	// Step 2: Merge sorted chunks and extract top N numbers
	topNumbers, err := mergeChunks(tempFiles, *n)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	// Step 3: Write top numbers to output file
	if err := writeNumbersToFile(topNumbers, *outputFile); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println("Processing complete. Results written to:", *outputFile)
}

func sortChunks(inputFile string, chunkSize int) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []uint64
	var tempFiles []string
	chunkIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		for _, numStr := range strings.Fields(line) {
			num, err := strconv.ParseUint(numStr, 10, 64)
			if err == nil {
				numbers = append(numbers, num)
			}
		}

		if len(numbers) >= chunkSize {
			tempFile, err := writeSortedChunk(numbers, chunkIndex)
			if err != nil {
				return nil, err
			}
			tempFiles = append(tempFiles, tempFile)
			numbers = nil
			chunkIndex++
		}
	}

	if len(numbers) > 0 {
		tempFile, err := writeSortedChunk(numbers, chunkIndex)
		if err != nil {
			return nil, err
		}
		tempFiles = append(tempFiles, tempFile)
	}

	return tempFiles, nil
}

func writeSortedChunk(numbers []uint64, chunkIndex int) (string, error) {
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] > numbers[j] })
	tempFileName := fmt.Sprintf("chunk_%d.tmp", chunkIndex)
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	for _, num := range numbers {
		fmt.Fprintln(tempFile, num)
	}

	return tempFileName, nil
}

func mergeChunks(tempFiles []string, n int) ([]uint64, error) {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	fileScanners := make([]*bufio.Scanner, len(tempFiles))
	files := make([]*os.File, len(tempFiles))

	for i, fileName := range tempFiles {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		files[i] = file
		fileScanners[i] = bufio.NewScanner(file)

		if fileScanners[i].Scan() {
			num, _ := strconv.ParseUint(fileScanners[i].Text(), 10, 64)
			heap.Push(minHeap, num)
		}
	}

	var topNumbers []uint64
	for minHeap.Len() > 0 && len(topNumbers) < n {
		top := heap.Pop(minHeap).(uint64)
		topNumbers = append(topNumbers, top)

		for i, scanner := range fileScanners {
			if scanner.Scan() {
				num, _ := strconv.ParseUint(scanner.Text(), 10, 64)
				heap.Push(minHeap, num)
			} else {
				files[i].Close()
				os.Remove(tempFiles[i])
				fileScanners[i] = nil
			}
		}
	}

	for _, file := range files {
		if file != nil {
			file.Close()
		}
	}

	return topNumbers, nil
}

func writeNumbersToFile(numbers []uint64, outputFile string) error {
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	for _, num := range numbers {
		fmt.Fprintln(output, num)
	}
	return nil
}
