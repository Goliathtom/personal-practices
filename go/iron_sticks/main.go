package main

import (
	"fmt"
	"strings"
)

func solution(arrangement string) int {
	var stack []string
	answer := 0
	arrangements := strings.Split(arrangement, "")

	for i, s := range arrangements {
		if s == "(" {
			stack = append(stack, s)
		} else {
			stack = stack[:len(stack)-1]
			if arrangements[i-1] == "(" {
				answer += len(stack)
			} else {
				answer++
			}
		}

	}

	return answer
}

func main() {
	arrangement := "()(((()())(())()))(())"

	fmt.Println(solution(arrangement))
}
