package string_sum

//package main

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

	if len(input) == 0 {
		return output, fmt.Errorf("%w", errorEmptyInput)
	}
	n := make([]int, 2)
	var e error
	m := make([]string, 2)
	j := 0
	input = strings.ReplaceAll(input, " ", "")
	for i := 1; i < len(input); i++ {
		if string(input[i]) == "+" || string(input[i]) == "-" {
			m[0] = input[j:i]
			j = i
			break
		} else {
			return output, fmt.Errorf("%w", errorNotTwoOperands)
		}
	}
	//fmt.Println(m[0])
	//m[0] = strings.ReplaceAll(m[0], " ", "")
	n[0], e = strconv.Atoi(m[0])
	if e != nil {
		return output, fmt.Errorf("%w", e)
	}
	//fmt.Println(n[0],e)
	m[1] = input[j:len(input)]
	//m[1] = strings.ReplaceAll(m[1], " ", "")
	//fmt.Println(m[1])
	for i := 1; i < len(m[1]); i++ {
		if string(m[1][i]) == "+" || string(m[1][i]) == "-" {
			return output, fmt.Errorf("%w", errorNotTwoOperands)
		}
	}
	n[1], e = strconv.Atoi(m[1])
	if e != nil {
		return output, fmt.Errorf("%w", e)
	}
	//fmt.Println(n[1],e)
	sum := n[0] + n[1]
	output = strconv.Itoa(sum)
	return output, nil
}

//func main(){
//	s := "24"
//	a,err := StringSum(s)
//	fmt.Printf("%s,%v",a,err)
//}
