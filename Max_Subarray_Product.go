package zleetcode

//设max[i]表示以a[i]结尾的最大连续子序列积
//设min[i]表示以a[i]结尾的最小连续子序列积
// max[i] = Max(max[i-1]*a[i], min[i-1]*a[i], a[i]), i , [1,n), max[0] = a[0]
// min[i] = Min(max[i-1]*a[i], min[i-1]*a[i], a[i]), i , [1,n), min[0] = a[0]
// result = Max(max[i]), i, a[0, n)
func maxSubarrayProduct(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	curmax := a[0]
	curmin := a[0]
	res := a[0]

	for i := 1; i < len(a); i++ {
		tmp1 := curmax * a[i]
		tmp2 := curmin * a[i]

		curmax = a[i]
		if tmp1 > curmax {
			curmax = tmp1
		}
		if tmp2 > curmax {
			curmax = tmp2
		}

		curmin = a[i]
		if tmp1 < curmin {
			curmin = tmp1
		}
		if tmp2 < curmin {
			curmin = tmp2
		}

		if curmax > res {
			res = curmax
		}
	}
	return res
}
