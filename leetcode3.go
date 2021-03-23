package leetcode

func lengthOfLongestSubstring(s string) int {
	var i, j, lenth = 0, 0, 0
	var ret = 0
	var statistic = make(map[uint8]int)
	for ; j < len(s); j++ {
		lenth++
		if index, ok := statistic[s[j]]; ok {
			// 当前字串中已经有j 对应的字母了
			for ; i <= index; i++ {
				delete(statistic, s[i])
				lenth--
			}
		}
		statistic[s[j]] = j
		if lenth > ret {
			ret = lenth
		}
		//fmt.Println(i, " ", j, " ", ret)
	}
	return ret
}
