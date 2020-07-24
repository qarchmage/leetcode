package main

import "fmt"

func main() {
	A := [][]int{{1, 4, 2}, {2, 5, 1, 2, 5}, {1, 3, 7, 1, 7, 5}, {3, 3}}
	B := [][]int{{1, 2, 4}, {10, 5, 2, 1, 5, 2}, {1, 9, 2, 5, 1}, {3}}
	//expected : 2,3,2,1
	for i := range A {
		fmt.Println(maxUncrossedLines(A[i], B[i]))
	}
}

// work in progress
// TODO: solve with dynamic programming
func maxUncrossedLines(A, B []int) int {
	m, n, ans := len(A), len(B), 0
	if m == 0 || n == 0 {
		return 0
	}
	for i, j := 0, 0; i < m; i++ {
		if j > len(B)-1 {
			break
		}
		if A[i] == B[j] {
			ans++
			j++
		} else {
			for k := j; k < n; k++ {
				if A[i] == B[k] {
					ans++
					j = k
					break
				}
			}
		}
	}
	return ans
}

/*
class Solution {
public:
   int maxUncrossedLines(vector<int>& A, vector<int>& B) {
        int m = A.size(), n = B.size(), dp[m+1][n+1];
        memset(dp, 0, sizeof(dp));
        for (int i = 1; i <= m; ++i)
            for (int j = 1; j <= n; ++j)
                dp[i][j] = A[i - 1] == B[j - 1] ? dp[i - 1][j - 1] + 1 : max(dp[i][j - 1], dp[i - 1][j]);
        return dp[m][n];
    }
};
*/
