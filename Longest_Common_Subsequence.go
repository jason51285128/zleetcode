package zleetcode

//l[i, j]截至到 s1[i], s2[j]的LCS
//l[i, j] = l[i-1,j-1] + 1, if s1[i] == s2[j]
//l[i, j] = max(l[i-1,j], l[i,j-1])) if s1[i] != s2[j]
//i >= 1, j >= 1
func LCS(s1, s2 string) int {
	//one of zero
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	//both not zero
	if len(s1) == 1  || len(s2) == 1 {
		//one of one, or both one
		if s1[0] == s2[0] {
			return 1
		}
		return 0
	}
	//both greater one
	i := len(s1) -1
	j := len(s2) -1
	if s1[i] == s2[j] {
		return  LCS(s1[:i], s2[:j])
	}
	tmp1 := LCS(s1,s2[:j])
	tmp2 := LCS(s1[:i], s2)
	if tmp1 > tmp2  {
		return tmp1
	}
	return tmp2
}