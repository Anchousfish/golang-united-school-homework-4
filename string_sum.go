package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	input = strings.Trim(input, "\t\n\v\f\r")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.TrimSpace(input)

	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	} else {
		num := strings.Split(input, "+")

		if len(num) == 1 || len(num) > 2 {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		} else {
			x1, err1 := strconv.Atoi(num[0])
			x2, err2 := strconv.Atoi(num[1])
			if err1 != nil || err2 != nil {
				if err1 != nil {
					err = err1
				} else {
					err = err2
				}
				return "", fmt.Errorf("%w", err)
			} else {

				output = strconv.Itoa(x1 + x2)
				return output, err
			}
		}

	}

}
