package leetcode

import (
	"path/filepath"
	"strings"
)

// api
func simplifyPath2(path string) string {
	return filepath.Clean(path)
}

// 正常解法
func simplifyPath1(path string) string {
	var stack []string
	split := strings.Split(path, "/")
	for _, s := range split {
		switch s {
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case ".":
		case "":
		default:
			stack = append(stack, s)
		}
	}
	var ret string
	if len(stack) == 0 {
		ret = "/"
	}
	ret = "/" + strings.Join(stack, "/")
	return ret
}