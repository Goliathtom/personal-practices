package main

import (
	"fmt"
	"sort"
)

func solution(cars [][]int) int {
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] < cars[j][0]
	})

	answer, position := 1, cars[0][1]

	for i := 0; i < len(cars)-1; i++ {
		if position > cars[i][1] {
			position = cars[i][1]
		}

		if position < cars[i+1][0] {
			answer++
			position = cars[i+1][1]
		}
	}

	return answer
}

func main() {
	fmt.Println(solution([][]int{{-20, 15}, {-14, -5}, {-18, -13}, {-5, -3}})) // 2
	fmt.Println(solution([][]int{{-10, 4}, {-6, 8}, {-15, -9}}))               // 2
	fmt.Println(solution([][]int{{-10, 4}, {-10, 8}, {-15, -9}}))              // 1
	fmt.Println(solution([][]int{{-10, -7}, {-6, -3}, {-1, 3}}))               // 3
}
