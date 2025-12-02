package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Directon string
	Clicks   int
}

func ParseInput(inputPath string) []Instruction {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	start := 50
	end := 0
	countZero := 0

	combination := make([]Instruction, 0)

	for scanner.Scan() {
		rotation := strings.SplitAfterN(scanner.Text(), "", 2)
		direction := rotation[0]
		fmt.Printf("Turn %s ", direction)
		clicks, err := strconv.Atoi(rotation[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d clicks\n", clicks)

		instruction := Instruction{
			Directon: direction,
			Clicks:   clicks,
		}

		combination = append(combination, instruction)

		if direction == "L" {
			end = start - clicks
		} else {
			end = start + clicks
		}
		start = end

		if start < 0 {
			start += 100
		} else if start > 99 {
			start -= 100
		}

		if start == 0 {
			countZero++
		}

		fmt.Println("The dial points at", start)
	}
	fmt.Printf("In the end the dial pointed at 0 location %d times\n", countZero)
	return combination
}
