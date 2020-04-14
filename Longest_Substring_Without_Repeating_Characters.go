package zleetcode

//最长不重复子序列

//用一个map记录已出现的字符，变量保存当前最长，当出现重复字符时，起始位置设置为
//前一个重复字符前面
func lswrcGreedy(s string) int {
	start := 0
	lastpos := make(map[byte]int)
	maxlength := 0
	for i := 0; i < len(s); i++ {
		index, ok := lastpos[byte(s[i])]
		if ok {
			if i-start > maxlength {
				maxlength = i - start
			}
			start = index + 1
		}
		lastpos[byte(s[i])] = i
	}
	if len(s)-start > maxlength {
		return len(s) - start
	}
	return maxlength

}

//l[i],以s[i]结尾的最长不重复子串的长度
//l[i] = l[i - 1] + 1, if s[i] not repeat
//l[i] = index{s[i]} - index{s[i] last repeat}
func lswrcDP(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	maxlength := 1
	curlen := 1
	lastpos := make(map[byte]int)
	lastpos[byte(s[0])] = 0
	for i := 1; i < len(s); i++ {
		index, ok := lastpos[byte(s[i])]
		if !ok {
			curlen++
			if curlen > maxlength {
				maxlength = curlen
			}
		} else {
			curlen = i - index
			if curlen > maxlength {
				maxlength = curlen
			}
		}
		lastpos[byte(s[i])] = i
	}
	return maxlength
}
