package leetcode

func longestOnes(A []int, K int) (ans int) {
	var left, lsum, rsum int
	for right, v := range A {
		rsum += 1 - v
		for rsum-lsum > K {
			lsum += 1 - A[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
