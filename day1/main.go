package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func processFile(filePath string) ([]int, []int) {
	array1 := make([]int, 0)
	array2 := make([]int, 0)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return array1, array2
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 2 && len(parts[0]) == 5 && len(parts[1]) == 5 {
			num1, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("Error converting num1: %v\n", err)
				continue
			}

			num2, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Error converting num2: %v\n", err)
				continue
			}

			array1 = append(array1, num1)
			array2 = append(array2, num2)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return array1, array2
	}

	sort.Ints(array1)
	sort.Ints(array2)

	return array1, array2
}
func calculateDifference(array1 []int, array2 []int) int64 {
	var difference int64
	for i := 0; i < len(array1); i++ {
		difference += int64(math.Abs(float64(array1[i] - array2[i])))
	}

	return difference
}

// Similarity score is the sum of the duplicates in the two arrays
func similarityScore(array1 []int, array2 []int) int64 {
	var similarity int64
	for i := 0; i < len(array1); i++ {
		similarity += findDuplicateCount(array2, array1[i])
	}
	return similarity
}
func findDuplicateCount(array []int, x int) int64 {
	var count int64
	count = 0
	for i := 0; i < len(array); i++ {
		if array[i] == x {
			count++
		}
	}
	return int64(x) * count
}
func main() {
	filePath := "input.txt"
	array1, array2 := processFile(filePath)

	//	fmt.Println("Array 1:", array1)
	//	fmt.Println("Array 2:", array2)

	difference := calculateDifference(array1, array2)
	fmt.Printf("The absolute difference between the two arrays is: %d\n", difference)
	similarityscore := similarityScore(array1, array2)
	fmt.Printf("The similarity score between the two arrays is: %d\n", similarityscore)
}
