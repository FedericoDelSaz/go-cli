package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readFile(file string) []string {
	var fileLines []string
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Print(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}

func converterArray(fileLines []string) []int64 {
	intArray := make([]int64, len(fileLines))
	print(len(intArray))
	for i, str := range fileLines {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Println("Error converting string to int64:", err)
			intArray[i] = -1
		} else {
			intArray[i] = num
		}
	}
	fmt.Println(intArray)
	return intArray
}

func orderArray(intArray []int64) []int64 {
	sort.Slice(intArray, func(i, j int) bool {
		return intArray[i] > intArray[j]
	})
	return intArray
}

func removeEmptyPositions(intArray []int64) []int64 {
	result := []int64{}

	for _, value := range intArray {
		if value > 0 {
			result = append(result, value)
		}
	}
	return result
}

func writeFile(fileName string, intArray []int64) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	defer file.Close()

	for _, value := range intArray {
		_, err := fmt.Fprintf(file, "%d\n", value)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}

	fmt.Println("Array written to file successfully.")
}

func main() {
	fileLines := readFile("input.txt")
	intArray := converterArray(fileLines)
	cleanedArray := removeEmptyPositions(intArray)
	orderedArray := orderArray(cleanedArray)
	writeFile("output.txt", orderedArray)
}
