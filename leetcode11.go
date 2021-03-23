package leetcode

func maxArea(height []int) int {
	//var min = func(x, y int) int {
	//	if x < y {
	//		return x
	//	}
	//	return y
	//}
	start, end := 0, len(height)-1
	var area, ret int
	for start <= end {
		var high, wide int
		wide = end - start
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		area = high * wide
		if area > ret {
			ret = area
		}
	}
	return ret
}

func maxArea1(height []int) int {
	var min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	if len(height) == 0 {
		return 0
	}
	var ret int
	for i := 0; i < len(height); i++ {
		for j := i+1; j < len(height); j++ {
			area := (j-i) * min(height[i], height[j])
			if area > ret {
				ret = area
			}
		}
	}
	return ret
}