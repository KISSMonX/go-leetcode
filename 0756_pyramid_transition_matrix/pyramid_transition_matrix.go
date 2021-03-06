package leetcode0756

import "strings"

func pyramidTransition(bottom string, allowed []string) bool {
	blocks := "#" + strings.Join(allowed, "#")

	var dfs func(string, int, int) bool
	dfs = func(bottom string, cur, length int) bool {
		if cur+2 > length {
			bottom = bottom[length:]
			length = len(bottom)
			if length == 1 {
				return true
			}
			return dfs(bottom, 0, length)
		}

		b := "#" + bottom[cur:cur+2]
		beg := 0
		for beg < len(blocks) {
			index := strings.Index(blocks[beg:], b) + beg
			if index < beg {
				break
			}
			beg = index + 4
			color := blocks[index+3 : index+4]
			if dfs(bottom+color, cur+1, length) {
				return true
			}
		}

		return false
	}

	return dfs(bottom, 0, len(bottom))
}
