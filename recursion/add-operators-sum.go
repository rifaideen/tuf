package main

import (
	"strconv"
	"strings"
)

func addOperators(num string, target int) []string {

	ans := []string{}

	addOperatorsRecursive(0, len(num), target, []string{}, &ans, num)

	return ans

}

func addOperatorsRecursive(index, n, target int, sequence []string, ans *[]string, num string) {

	if index == n {

		sequence = sequence[:len(sequence)-1]

		*ans = append(*ans, strings.Join(sequence, ""))

		return

	}

	if exceed(sequence, target) {
		return
	}

	operators := []string{"+", "-", "*"}

	for _, v := range operators {
		sequence = append(sequence, string(num[index]), v)

		addOperatorsRecursive(index+1, n, target, sequence, ans, num)

		sequence = sequence[:len(sequence)-2]

	}

}

func exceed(sequence []string, target int) bool {

	sum := 0

	for i := 0; i < len(sequence); i += 2 {

		value, _ := strconv.Atoi(sequence[i])

		switch sequence[i+1] {

		case "+":

			sum += value

		case "-":

			sum -= value

		case "*":

			sum *= value

		}

	}

	return sum > target

}
