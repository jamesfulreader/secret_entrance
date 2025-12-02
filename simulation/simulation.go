package simulation

import "github.com/jamesfulreader/secret_entrance/utils"

func SimulateDial(start int, instructions []utils.Instruction) int {
	pos := start
	zeroCount := 0

	for _, inst := range instructions {
		for i := 0; i < inst.Clicks; i++ {
			if inst.Directon == "L" {
				pos--
			} else {
				pos++
			}

			if pos < 0 {
				pos += 100

			}
			if pos > 99 {
				pos -= 100
			}

			if pos == 0 {
				zeroCount++
			}
		}
	}

	return zeroCount
}
