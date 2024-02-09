package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func processLine(line string) int {
	firstNumber, lastNumber := -1, -1
	for i := 0; i < len(line); i++ {
		if value, err := strconv.Atoi(string(line[i])); err == nil {
			if firstNumber == -1 {
				firstNumber = value
			} else {
				lastNumber = value
			}
		}
	}

	// Number is single digit. In this case we repeat so 7 becomed 77
	if lastNumber == -1 {
		return firstNumber + firstNumber*10
	}

	return firstNumber*10 + lastNumber
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		sum += processLine(scanner.Text())
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while scanning")
	}
}
