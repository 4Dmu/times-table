package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func printDigits(digits string, maxDigits int) {
	l := len(digits)
	if l >= maxDigits {
		fmt.Printf(" %s ", digits)
	}
	padding := strings.Repeat(" ", maxDigits-l)
	fmt.Printf(" %s%s ", padding, digits)
}

func loop(currentNum int, numbers []int, maxDigits int, end int) {
	printDigits(fmt.Sprintf("%d", currentNum), maxDigits)
	for _, num := range numbers {
		digits := fmt.Sprintf("%d", num*currentNum)
		printDigits(digits, maxDigits)
	}
	fmt.Print("\n\n")
	if currentNum == end {
		return
	}
	loop(currentNum+1, numbers, maxDigits, end)
}

func makeNumbers(end int) []int {
	nums := make([]int, end)

	for i := 1; i <= end; i++ {
		nums[i-1] = i
	}

	return nums
}

func main() {

	max := 10000
	args := os.Args
	var end int

	if len(args) < 2 {
		end = 18
	} else if len(args) > 3 {
		log.Fatal("To many arguments")
	} else {
		userInput := args[1]
		parsedUserInput, err := strconv.Atoi(userInput)
		if err != nil {
			log.Fatal(err)
		}

		if parsedUserInput <= 0 || parsedUserInput > max {
			log.Fatalf("length must be > 0 and < %d", max)
		}
		end = parsedUserInput
	}

	maxDigits := len(fmt.Sprintf("%d", end*end)) + 1
	nums := makeNumbers(end)
	currNum := 1

	loop(currNum, nums, maxDigits, end)
}
