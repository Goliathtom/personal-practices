package main

import "fmt"

type truck struct {
	t int
	w int
}

func solution(bridgeLength int, weight int, truckWeights []int) int {
	var queue []truck

	var time, currentBridgeWeight int = 0, 0
	for len(queue) > 0 || len(truckWeights) > 0 {
		time++

		if len(queue) > 0 {
			var firstTruck truck = queue[0]
			if firstTruck.t+bridgeLength == time {
				currentBridgeWeight -= firstTruck.w
				queue = queue[1:]
			}
		}

		if len(truckWeights) > 0 {
			if weight >= currentBridgeWeight+truckWeights[0] {
				currentBridgeWeight += truckWeights[0]
				queue = append(queue, truck{t: time, w: truckWeights[0]})
				truckWeights = truckWeights[1:]
			}
		}
	}

	return time
}

type testCase struct {
	bl     int
	w      int
	trucks []int
	expect int
}

func main() {
	var testCases []testCase
	testCases = append(testCases,
		testCase{bl: 2, w: 10, trucks: []int{7, 4, 5, 6}, expect: 8},
		testCase{bl: 100, w: 100, trucks: []int{10}, expect: 101},
		testCase{bl: 100, w: 100, trucks: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, expect: 110},
	)

	for i, testCase := range testCases {
		fmt.Printf("Case %d ===================\n", i+1)
		time := solution(testCase.bl, testCase.w, testCase.trucks)
		fmt.Println("Time : ", time)
		fmt.Println("Test Case is", time == testCase.expect)
	}
}
