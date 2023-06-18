package lcs

import "fmt"

/*
最长公共子序列
longest common sequence
https://oi-wiki.org/dp/basic/#%E6%9C%80%E9%95%BF%E5%85%AC%E5%85%B1%E5%AD%90%E5%BA%8F%E5%88%97
*/

func LCS(a, b string) int {

	dp := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, len(b))
	}

	for i, ach := range a {
		for j, bch := range b {
			if ach == bch {
				var tmp int
				if i > 0 && j > 0 {
					tmp = dp[i-1][j-1]
				} else {
					tmp = 0
				}
				dp[i][j] = tmp + 1
			} else {
				var tmp1 int
				if i > 0 {
					tmp1 = dp[i-1][j]
				} else {
					tmp1 = 0
				}

				var tmp2 int
				if j > 0 {
					tmp2 = dp[i][j-1]
				} else {
					tmp2 = 0
				}

				dp[i][j] = func() int {
					if tmp1 > tmp2 {
						return tmp1
					}
					return tmp2
				}()
			}
		}
	}

	fmt.Printf("%+v\n", dp)

	return dp[len(a)-1][len(b)-1]
}
