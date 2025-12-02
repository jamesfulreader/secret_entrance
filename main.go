package main

import (
	"fmt"

	"github.com/jamesfulreader/secret_entrance/utils"
)

func main() {
	fmt.Println("Secret entrance AOC 2025")

	combination := utils.ParseInput("test.txt")

	for _, v := range combination {
		fmt.Printf("From our combination turn %s %d clicks\n", v.Directon, v.Clicks)
	}

}
