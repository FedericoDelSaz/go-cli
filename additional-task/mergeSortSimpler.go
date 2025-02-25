package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Define flags
	n := flag.Int("--n", 1, "Number of top results.")
	inputFile := flag.String("--inputFile", "", "Path to input file.")
	outputFile := flag.String("--outputFile", "", "Path to output file.")
	chunkSize := flag.Int("--chunkSize", 1000000, "Numbers per chunk for sorting.")
	help := flag.Bool("--help", false, "Display help message.")
	h := flag.Bool("-h", false, "Display help message.")

	flag.Parse()

	// Display help message if --help or -h is provided
	if *help || *h {
		flag.Usage()
		os.Exit(0)
	}

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
	topNumbers := []uint64{}

	for _, fileName := range tempFiles {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			num, _ := strconv.ParseUint(scanner.Text(), 10, 64)
			topNumbers = insertSorted(topNumbers, num, n)
		}
		file.Close()
		os.Remove(fileName)
	}

	return topNumbers, nil
}

func insertSorted(slice []uint64, num uint64, maxLen int) []uint64 {
	index := sort.Search(len(slice), func(i int) bool { return slice[i] < num })
	slice = append(slice[:index], append([]uint64{num}, slice[index:]...)...)

	if len(slice) > maxLen {
		slice = slice[:maxLen]
	}
	return slice
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
