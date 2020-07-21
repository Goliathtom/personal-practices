package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(baseballSets [][]int) int {
	var answer []int
	for num := 123; num <= 987; num++ {
		numbers := strings.Split(strconv.Itoa(num), "")

		isValid := isValidCondition(numbers)
		if !isValid {
			continue
		}

		flag := isCorrectExpect(baseballSets, numbers)

		if flag {
			answer = append(answer, num)
		}
	}
	return len(answer)
}

func isValidCondition(numbers []string) bool {
	isValid := true

	for i := 0; i < len(numbers); i++ {
		curNumber, error := strconv.Atoi(numbers[i])
		if error != nil {
			panic("error")
		}

		if curNumber == 0 {
			isValid = false
		}

		nextIndex := (i + 1) % len(numbers)
		if numbers[nextIndex] == numbers[i] {
			isValid = false
		}
	}

	return isValid
}

func isCorrectExpect(baseballSets [][]int, numbers []string) bool {
	flag := true
	for _, baseballSet := range baseballSets {
		exNums := strings.Split(strconv.Itoa(baseballSet[0]), "")
		exSCnt := baseballSet[1]
		exBCnt := baseballSet[2]

		sCnt, bCnt := getCnts(exNums, numbers)
		if exSCnt != sCnt {
			flag = false
			break
		}

		if exBCnt != bCnt-sCnt {
			flag = false
			break
		}
	}

	return flag
}

func getCnts(exNums, numbers []string) (int, int) {
	var sCnt, bCnt int = 0, 0
	for j, exNum := range exNums {
		if numbers[j] == exNum {
			sCnt++
		}
	}

	for _, exNum := range exNums {
		foundIndex := indexOf(exNum, numbers)
		if foundIndex >= 0 {
			bCnt++
		}
	}

	return sCnt, bCnt
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func main() {
	var testCases [][]int

	testCases = append(testCases, []int{123, 1, 1}, []int{356, 1, 0}, []int{327, 2, 0}, []int{489, 0, 1})
	fmt.Println(solution(testCases))
}
