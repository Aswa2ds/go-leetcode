package leetcode

func maxSubArray(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	ret := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}
