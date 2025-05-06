package main

import "strconv"

func addOperators(num string, target int) []string {
	ans := []string{}

	addOperatorsRecursive(0, len(num), 0, 0, target, "", &ans, num)

	return ans
}

func addOperatorsRecursive(index, n, current, previous, target int, sequence string, ans *[]string, num string) {
	if index == n {
		if current == target {
			*ans = append(*ans, sequence)
		}

		return
	}

	for i := index; i < n; i++ {
		// to avoid creating sequence with leading zero (i.e 05)
		if i > index && num[index] == '0' {
			break
		}

		currentStr := num[index : i+1]
		currentNum, _ := strconv.Atoi(currentStr)

		if index == 0 {
			addOperatorsRecursive(i+1, n, currentNum, currentNum, target, currentStr, ans, num)
		} else {
			addOperatorsRecursive(i+1, n, current+currentNum, currentNum, target, sequence+"+"+currentStr, ans, num)
			addOperatorsRecursive(i+1, n, current-currentNum, -currentNum, target, sequence+"-"+currentStr, ans, num)
			addOperatorsRecursive(i+1, n, (current-previous)+(previous*currentNum), previous*currentNum, target, sequence+"*"+currentStr, ans, num)
		}
	}
}
