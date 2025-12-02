package utils

import (
	"bufio"
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

	combination := make([]Instruction, 0)

	for scanner.Scan() {
		rotation := strings.SplitAfterN(scanner.Text(), "", 2)
		direction := rotation[0]

		clicks, err := strconv.Atoi(rotation[1])
		if err != nil {
			log.Fatal(err)
		}

		instruction := Instruction{
			Directon: direction,
			Clicks:   clicks,
		}

		combination = append(combination, instruction)
	}
	return combination
}
