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
	plus                = '+'
	minus               = '-'
	operandsLen         = 2
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

func emptyCheck(input string) (output string, err error) {
	output = strings.ReplaceAll(input, " ", "")
	if output == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	return
}

func strSplit(input string, strLen int) (output []string) {
	runeInput := []rune(input)
	var word []rune
	for i := 0; i < strLen; i++ {
		if i != 0 && (runeInput[i] == plus || runeInput[i] == minus) {
			output = append(output, string(word))
			word = nil
			word = append(word, runeInput[i])
		} else {
			word = append(word, runeInput[i])
		}
	}
	output = append(output, string(word))
	return
}

func StringSum(input string) (output string, err error) {
	strSpaceCleaned, err := emptyCheck(input)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	splittedInput := strSplit(strSpaceCleaned, len([]rune(strSpaceCleaned)))
	if len(splittedInput) != operandsLen {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	} else {
		firstOperand, err := strconv.Atoi(splittedInput[0])
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}
		secondOperand, err := strconv.Atoi(splittedInput[1])
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}
		output = strconv.Itoa(firstOperand + secondOperand)
	}
	return
}
