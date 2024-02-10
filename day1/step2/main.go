package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func stringToInt(input string) int {
	switch input {
	case "1":
		fallthrough
	case "one":
		return 1
	case "2":
		fallthrough
	case "two":
		return 2
	case "3":
		fallthrough
	case "three":
		return 3
	case "4":
		fallthrough
	case "four":
		return 4
	case "5":
		fallthrough
	case "five":
		return 5
	case "6":
		fallthrough
	case "six":
		return 6
	case "7":
		fallthrough
	case "seven":
		return 7
	case "8":
		fallthrough
	case "eight":
		return 8
	default:
		return 9
	}
}

func processLine(line string) int {
	r, _ := regexp.Compile(`one|two|three|four|five|six|seven|eight|nine|\d`)
	var firstNumber, lastNumber int

	firstNumber = stringToInt(r.FindString(line))
	var match string
	for i := len(line) - 1; i >=0; i-- {
		if match = r.FindString(line[i:len(line)]); len(match) == 0 {
			continue
		}
		lastNumber = stringToInt(match)
		break
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
		fmt.Println(scanner.Text(), "=", processLine(scanner.Text()))
		sum += processLine(scanner.Text())
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while scanning")
	}
}
