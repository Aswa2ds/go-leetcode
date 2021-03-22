package leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	map1 := [26]int{}
	map2 := [26]int{}

	for i := 0; i < len(s1); i++ {
		map1[s1[i]-'a']++
		map2[s2[i]-'a']++
	}

	if map1 == map2 {
		return map1 == map2
	}

	for start, i := 0, len(s1); i < len(s2); start, i = start+1, i+1 {
		map2[s2[start]-'a']--
		map2[s2[i]-'a']++

		if map1 == map2 {
			return true
		}
	}
	return false
}
