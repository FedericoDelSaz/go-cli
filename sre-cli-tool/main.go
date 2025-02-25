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

// Error messages
const (
	errorInputFileNotExist    = "ERROR: input file does not exist."
	errorInputFileNotReadable = "ERROR: input file is not readable."
	errorInvalidNValue        = "ERROR: Number of top results must be bigger than 0."
	errorNValueTooHigh        = "ERROR: The maximum number of top results must be less or equal than 30000000."
	warnInvalidLine           = "WARN: Invalid line"
)

func main() {
	// Define flags
	n := flag.Int("n", 1, "The value N for the number of top results.")
	inputFile := flag.String("input-file", "", "Path to the input file.")
	outputFile := flag.String("output-file", "", "Path to the output file.")

	// Parse the flags
	flag.Parse()

	// Validate inputs
	if err := validateFlags(*n, *inputFile, *outputFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read numbers from the file
	numbers, err := readNumbersFromFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Sort numbers in descending order
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	// Get top N largest numbers
	if len(numbers) > *n {
		numbers = numbers[:*n]
	}

	// Write output to file
	if err := writeNumbersToFile(numbers, *outputFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Processing complete. Results written to:", *outputFile)
}

// validateFlags validates the input flags
func validateFlags(n int, inputFile, outputFile string) error {
	if n <= 0 {
		return fmt.Errorf(errorInvalidNValue)
	}
	if n > 30000000 {
		return fmt.Errorf(errorNValueTooHigh)
	}
	if inputFile == "" {
		return fmt.Errorf("ERROR: Input file path is required.")
	}
	if outputFile == "" {
		return fmt.Errorf("ERROR: Output file path is required.")
	}
	return nil
}

// readNumbersFromFile reads numbers from the input file and returns them as a slice of uint64
func readNumbersFromFile(inputFile string) ([]uint64, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf(errorInputFileNotExist)
		}
		return nil, fmt.Errorf(errorInputFileNotReadable)
	}
	defer file.Close()

	var numbers []uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Fields(line)
		for _, numStr := range numStrs {
			num, err := strconv.ParseUint(numStr, 10, 64)
			if err != nil {
				// If the line is invalid, print a warning and skip
				fmt.Printf("%s %s\n", warnInvalidLine, line)
				continue
			}
			numbers = append(numbers, num)
		}
	}

	// Handle error from scanner
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ERROR: Error reading the input file: %v", err)
	}
	return numbers, nil
}

// writeNumbersToFile writes the given numbers to the specified output file
func writeNumbersToFile(numbers []uint64, outputFile string) error {
	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("ERROR: Could not create output file: %v", err)
	}
	defer output.Close()

	for _, num := range numbers {
		_, err := fmt.Fprintln(output, num)
		if err != nil {
			return fmt.Errorf("ERROR: Failed to write to output file: %v", err)
		}
	}

	return nil
}
