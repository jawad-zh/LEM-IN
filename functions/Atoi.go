package function

import "unicode"

func Atoi(str string) int {
	const (
		MaxInt = 1<<31 - 1
		MinInt = -1 << 31
	)

	i := 0
	n := len(str)
	result := 0
	sign := 1

	for i < n && str[i] == ' ' {
		i++
	}


	for i < n && unicode.IsDigit(rune(str[i])) {
		digit := int(str[i] - '0')

		// Check for overflow
		if result > (MaxInt-digit)/10 {
			if sign == 1 {
				return MaxInt
			} else {
				return MinInt
			}
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}
