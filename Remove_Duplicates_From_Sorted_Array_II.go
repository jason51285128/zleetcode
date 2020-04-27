package zleetcode

func removeDuplicatesII(a []int) []int {
	if len(a) <= 2 {
		return a
	}
	unique := 1
	for i := 2; i < len(a); i++ {
		if a[unique-1] != a[i] {
			unique++
			a[unique] = a[i]
		}
	}
	unique++
	return a[0:unique]
}
