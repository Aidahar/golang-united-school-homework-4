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
	errorNotAnumber     = errors.New("value not a number")
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
	str := strings.Split(strings.TrimSpace(input), "+")

	if str == nil {
		fmt.Errorf("error: %w", errorEmptyInput)
	} else if len(str) > 3 || len(str) < 3 {
		fmt.Errorf("error: %w", errorNotTwoOperands)
	}
	num1, err := strconv.Atoi(str[0])
	if err != nil {
		fmt.Errorf("error: %w", errorNotAnumber)
	}
	num2, err := strconv.Atoi(str[1])
	if err != nil {
		fmt.Errorf("error: %w", errorNotAnumber)
	}

	output = string(num1 + num2)
	return output, nil
}
