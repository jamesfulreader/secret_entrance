package simulation

import "github.com/jamesfulreader/secret_entrance/utils"

func SimulateDial(start int, instructions []utils.Instruction) int {
	pos := start
	zeroCount := 0

	for _, inst := range instructions {
		if inst.Directon == "L" {
			pos -= inst.Clicks
		} else {
			pos += inst.Clicks
		}

		for pos < 0 {
			pos += 100
		}
		for pos > 99 {
			pos -= 100
		}

		if pos == 0 {
			zeroCount++
		}
	}

	return zeroCount
}
