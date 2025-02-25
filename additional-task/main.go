package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"os"
	"strconv"
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

func validateInputs(n int, inputFile, outputFile string) {
	if n <= 0 {
		fmt.Println("ERROR: Number of top results must be greater than 0.")
		os.Exit(1)
	}
	if inputFile == "" || outputFile == "" {
		fmt.Println("ERROR: Input and output file paths are required.")
		os.Exit(1)
	}
}

func findTopN(filename string, N int) ([]uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	h := &MinHeap{}
	heap.Init(h)

	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			continue
		}
		if h.Len() < N {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}
	return *h, scanner.Err()
}

func writeToFile(outputFile string, result []uint64) {
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("ERROR: Could not create output file:", err)
		os.Exit(1)
	}
	defer output.Close()

	for _, num := range result {
		_, err := fmt.Fprintln(output, num)
		if err != nil {
			fmt.Println("ERROR: Failed to write to output file:", err)
			os.Exit(1)
		}
	}
}

func main() {
	n := flag.Int("n", 10, "Number of top results to retrieve.")
	inputFile := flag.String("input-file", "", "Path to the input file.")
	outputFile := flag.String("output-file", "", "Path to the output file.")
	flag.Parse()

	validateInputs(*n, *inputFile, *outputFile)
	result, err := findTopN(*inputFile, *n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	writeToFile(*outputFile, result)
	fmt.Println("Processing complete. Results written to:", *outputFile)
}
