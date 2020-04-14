package zleetcode

//s[i], 以a[i]结尾的最大子数组和
//s[i] = max(s[i-1] + a[i], a[i]), i >= 1

func maxSubarraySum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	max := a[0]
	cur := a[0]
	for i := 1; i < len(a); i++ {
		if cur+a[i] > cur {
			cur = cur + a[i]
		} else {
			cur = a[i]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
