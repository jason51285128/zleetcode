package zleetcode

//l[i, j]以 s1[i], s2[j]的LCSA
//l[i, j] = l[i-1,j-1] + 1, if s1[i] == s2[j]
//l[i, j] = 0 if s1[i] != s2[j]
//i >= 1, j >= 1
func LCSA(s1, s2 string) int {
	res := 0
	t := make([][]int, len(s1)+1)
	for i := range t {
		t[i] = make([]int, len(s2)+1)
	}
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 || j == 0 {
				t[i][j] = 0
			} else if s1[i-1] == s2[j-1] {
				t[i][j] = t[i-1][j-1] + 1
				if t[i][j] > res {
					res = t[i][j]
				}
			} else {
				t[i][j] = 0
			}
		}
	}
	return res
}
