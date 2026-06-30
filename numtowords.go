package numtowords

import (
	"fmt"
	"math"
)

// MinNum is the smallest number that can be converted to words.
const MinNum = -999

// MaxNum is the largest number that can be converted to words.
const MaxNum = 999

// Convert converts the specified number to words.
func Convert(number int) (string, error) {
	if number > MaxNum || number < MinNum {
		return "", fmt.Errorf("can only convert numbers between %v and %v", MinNum, MaxNum)
	}

	if number == 0 {
		return "zero", nil
	}

	units := [20]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	tens := [8]string{
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	result := ""

	if number < 0 {
		result += "minus "
		number = int(math.Abs(float64(number)))
	}

	if number > 99 {
		hundredindex := number / 100
		result += units[hundredindex] + " hundred"

		number = number % 100

		if number == 0 {
			return result, nil
		}

		result += " and "
	}

	if number > 19 {
		tensindex := number/10 - 2

		result += tens[tensindex]
		number = number % 10

		if number == 0 {
			return result, nil
		}

		result += " "
	}

	result += units[number]

	return result, nil
}
