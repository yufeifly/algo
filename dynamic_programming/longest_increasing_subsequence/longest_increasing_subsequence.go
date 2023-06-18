package longestincreasingsubsequence

import (
	"fmt"
	"math"

)

/*
最长不下降子序列
longest_increasing_subsequence
https://oi-wiki.org/dp/basic/#%E6%9C%80%E9%95%BF%E4%B8%8D%E4%B8%8B%E9%99%8D%E5%AD%90%E5%BA%8F%E5%88%97
*/

// 复杂度n^2
func LongestIncreasingSubsequence(arr []int) int {
	dp := make([]int, len(arr))

	for i := range dp{
		dp[i] = 1
	}

	max := func(a, b int)int {
		if a > b {
			return a
		}
		return b
	}

	var ans = 1

	for j:=1;j<len(arr);j++ {
		for i:=0;i<j;i++ {
			if arr[j] >= arr[i] {
				dp[j] = max(dp[j], dp[i] + 1)
				ans= max(dp[j], ans)
			}
		}
	}

	return ans
}

func LongestIncreasingSubsequence2(arr []int) int {
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = arr[0]

	// 找第一个比d大的数字的索引
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

	for i:=1;i<len(arr);i++ {
		idx := f(dp, arr[i])
		dp[idx] = arr[i]
	}

	ans := 0
	for i := range dp {
		if dp[i] != math.MaxInt {
			ans ++
		}
	}

	fmt.Println(dp)

	return ans
}
