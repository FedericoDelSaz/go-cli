package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const maxTopResults = 30000000

// Custom error messages
var (
	errFileNotExist    = fmt.Errorf("ERROR: input file does not exist")
	errFileNotReadable = fmt.Errorf("ERROR: input file is not readable")
	errInvalidN        = fmt.Errorf("ERROR: Number of top results must be bigger than 0")
	errNOverUpperLimit = fmt.Errorf("ERROR: The maximum number of top results must be less or equal than 30000000")
	errInvalidLine     = fmt.Errorf("WARN: Invalid line")
)

// Custom type for a slice of uint64 to implement sort.Interface
type Uint64Slice []uint64

// Implement the sort.Interface methods
func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Uint64Slice) Less(i, j int) bool { return p[i] > p[j] } // Sort descending

// checkError checks if an error occurred and exits with a message.
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// readFileLines reads the file line by line and parses unsigned 64-bit integers.
func readFileLines(filename string, n int) ([]uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errFileNotExist
		}
		return nil, errFileNotReadable
	}
	defer file.Close()

	var numbers []uint64
	scanner := bufio.NewScanner(file)

	// Reading lines and parsing unsigned 64-bit integers
	for scanner.Scan() {
		line := scanner.Text()
		// Parse the line to uint64
		num, err := parseUint64(line)
		if err != nil {
			// Warn about invalid line
			fmt.Printf("WARN: Invalid line %s\n", line)
			continue
		}

		// Append to the numbers slice
		numbers = append(numbers, num)
		if len(numbers) > n {
			// Keep only top N numbers
			sort.Sort(Uint64Slice(numbers))
			numbers = numbers[:n]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

// parseUint64 tries to parse a string into a uint64, returns an error if invalid.
func parseUint64(s string) (uint64, error) {
	// Only allow valid unsigned 64-bit integers
	if strings.Contains(s, " ") {
		return 0, fmt.Errorf("invalid")
	}
	return strconv.ParseUint(s, 10, 64)
}

// writeFile writes the sorted numbers to an output file.
func writeFile(filename string, numbers []uint64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func main() {
	// Parse arguments
	nFlag := 0
	inputFileFlag := ""
	outputFileFlag := ""

	// Read arguments from the command line
	for i := 0; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-n":
			i++
			fmt.Sscanf(os.Args[i], "%d", &nFlag)
		case "--input-file":
			i++
			inputFileFlag = os.Args[i]
		case "--output-file":
			i++
			outputFileFlag = os.Args[i]
		}
	}

	// Argument validation
	if nFlag <= 0 {
		checkError(errInvalidN)
	}
	if nFlag > maxTopResults {
		checkError(errNOverUpperLimit)
	}
	if inputFileFlag == "" {
		checkError(fmt.Errorf("ERROR: --input-file must be specified"))
	}
	if outputFileFlag == "" {
		checkError(fmt.Errorf("ERROR: --output-file must be specified"))
	}

	// Read file and get the top N largest numbers
	numbers, err := readFileLines(inputFileFlag, nFlag)
	checkError(err)

	// Sort the numbers in descending order
	sort.Sort(Uint64Slice(numbers))

	// Write to output file
	err = writeFile(outputFileFlag, numbers)
	checkError(err)

	// Print success message
	fmt.Printf("Successfully wrote %d numbers to %s\n", len(numbers), outputFileFlag)
}
