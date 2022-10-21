package fizzBuzz

import (
	"strconv"
)

// GetFizzBuzz Returns a slice of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2
// are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
func GetFizzBuzz(int1 int, int2 int, limit int, str1 string, str2 string) ([]string, error) {
	var result []string
	// Iterate from 1 to limit and build the response accordingly to the rules
	for i := 1; i < limit; i++ {
		var iterationStr string
		// Check if a condition is met
		if i%int1 == 0 || i%int2 == 0 {
			if i%int1 == 0 {
				iterationStr += str1
			}
			if i%int2 == 0 {
				iterationStr += str2
			}
		} else { // If no condition is met, add the default value to the result
			iterationStr += strconv.Itoa(i)
		}
		result = append(result, iterationStr)
	}

	return result, nil // Error is always nil in this fizz-buzz example, so this could be removed.
}
