package helper

import (
	"sort"
)

func BinaryFindString(arr []string, target string) bool {
	l := 0
	h := len(arr) - 1
	sort.Strings(arr)

	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return false
}
