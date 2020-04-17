package zleetcode

//a[0]<a[1], a[n-2]>a[n-1](只是两头元素满足，并不是递推)，n>=3，
//且中间数据两两之间不相等，求一个m，使得a[m-1]<a[m]&&a[m]<a[m+1]
//首先m一定是存在的，两两不相等，且两头形成“八字形”， 所以中间不管怎么变化，都是存在的
//使用二分，若满足则返还，不满足，则看两头哪一边大，取大的那一边

func findm(a []int) int {
	start := 0
	end := len(a) - 1
	for end-start > 2 {
		mid := (start + end) / 2
		if a[mid] > a[mid-1] && a[mid] > a[mid+1] {
			return mid
		}
		if a[mid-1] > a[mid] {
			end = mid
			continue
		}
		start = mid
	}
	return start + 1
}
