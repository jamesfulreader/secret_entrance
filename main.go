package main

import (
	"fmt"

	"github.com/jamesfulreader/secret_entrance/simulation"
	"github.com/jamesfulreader/secret_entrance/utils"
)

func main() {
	fmt.Println("Secret entrance AOC 2025")

	combination := utils.ParseInput("input.txt")

	count := simulation.SimulateDial(50, combination)
	fmt.Println(count)
}
