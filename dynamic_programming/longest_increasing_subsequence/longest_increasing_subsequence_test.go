package longestincreasingsubsequence

import (
	"fmt"
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	arr := []int{
		1, 4, 2, 3,
	}
	ans := LongestIncreasingSubsequence(arr)

	fmt.Println(ans)
}

func TestLongestIncreasingSubsequence2(t *testing.T) {
	arr := []int{
		1, 2, 3, -1, -2, 7, 9,
	}
	ans := LongestIncreasingSubsequence2(arr)
	fmt.Println(ans)
}

func TestF(t *testing.T) {
	f := func(vec []int, d int) int {
		left, right := 0, len(vec)-1
		for (left < right) {
			mid := left + (right - left) / 2
			if vec[mid] > d {
				right = mid
			} else if vec[mid] == d {
				left = mid + 1
			} else {
				left = mid + 1
			}
		}

		return left
	}

	i := f([]int{3, 5, 6}, 2)
	fmt.Println(i)
}