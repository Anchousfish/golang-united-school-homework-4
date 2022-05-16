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
	var int1, int2 int64
	var err1, err2 error
	vars := make([]string, 2)

	input = strings.Replace(input, " ", "", -1)
	r_slice := []rune(input)

	if len(r_slice) > 4 {
		output = ""
		err = fmt.Errorf("%w", errorNotTwoOperands)
	}

	if input == "" {
		output = ""
		err = fmt.Errorf("%w", errorEmptyInput)
	}

	if string(r_slice[0]) == "+" || string(r_slice[0]) == "-" {
		vars[0] = string(r_slice[0:2])
		vars[1] = string(r_slice[2:4])
	} else {
		vars[0] = string(r_slice[0:1])
		vars[1] = string(r_slice[1:3])
	}
	int1, err1 = strconv.ParseInt(vars[0], 10, 0)
	err1 = fmt.Errorf("%w", err1)
	int2, err2 = strconv.ParseInt(vars[1], 10, 0)
	err2 = fmt.Errorf("%w", err2)
	if err1 == nil && err2 == nil {
		summ := int1 + int2
		output = strconv.Itoa(int(summ))
		err = nil
	} else {
		output = ""
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
	}

	return output, err
}
