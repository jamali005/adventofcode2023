package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func findFirstNumber(line string) (int, bool) {
	for _, char := range line {
		if unicode.IsDigit(char) {
			return int(char - '0'), true
		}
	}
	return 0, false
}

func findLastNumber(line string) (int, bool) {
	for i := len(line) - 1; i >= 0; i-- {
		char := line[i]
		if unicode.IsDigit(rune(char)) {
			return int(char - '0'), true

		}
	}

	return 0, false
}

func getCalibration(line string) int {
	firstDigit, firstDigitFound := findFirstNumber(line)
	lastDigit, lastDigitFound := findLastNumber(line)
	if firstDigitFound && lastDigitFound {
		result := firstDigit*10 + lastDigit
		fmt.Printf("%s => %d\n", line, result)
		return firstDigit*10 + lastDigit
	}

	return 0
}

func partOne(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	totalCalibration := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalCalibration += getCalibration(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	return totalCalibration
}

func main() {
	inputFilename := "Adventofcodeday1.txt"

	// Run part one
	resultPartOne := partOne(inputFilename)
	fmt.Println("Part One Result:", resultPartOne)
}
