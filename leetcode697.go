package leetcode

import "fmt"

func findShortestSubArray(nums []int) int {
	statistic := make(map[int]int)
	for _, num := range nums {
		statistic[num]++
	}
	fmt.Println(statistic)
	return 0
}
