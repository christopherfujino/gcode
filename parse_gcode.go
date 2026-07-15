package main

import (
	"fmt"
	"math"
)

func ParseGCode(line string) Instruction {
	var runes = []rune(line)
	if runes[0] != 'G' {
		panic(0)
	}
	runes = runes[1:]
	var postDot = false
	var preDotDigits []float64
	var postDotDigits []float64

	for _, thisRune := range runes {
		if thisRune == '.' {
			postDot = true
			continue
		} else if thisRune >= '0' && thisRune <= '9' {
			if postDot {
				postDotDigits = append(postDotDigits, float64(thisRune-'0'))
			} else {
				preDotDigits = append(preDotDigits, float64(thisRune-'0'))
			}
		} else {
			break
		}
	}

	if len(preDotDigits) == 0 {
		panic(fmt.Sprintf("Parsing G with no numbers: %v, %v, %s", preDotDigits, postDotDigits, line))
	}

	var code float64

	// e.g. 420
	// index   0   1  2
	// len-i-1 2   1  0
	// pow     2   1  0
	//        [4,  2, 0]
	for i, digit := range preDotDigits {
		var zeroes = len(preDotDigits) - i - 1
		code += digit * math.Pow10(zeroes)
	}

	// e.g. 0.103
	// index   0   1    2
	// -1 - i  -1  -2   -3
	// pow10   -1  -2   -3
	//        [1,  0,   3]
	for i, digit := range postDotDigits {
		var zeroes = -1 - i
		code += digit * math.Pow10(zeroes)
	}

	fmt.Printf("%s -> G%f\n", line, code)

	return Instruction{
		Type:      'g',
		Code:      code,
		DebugLine: line,
	}
}
