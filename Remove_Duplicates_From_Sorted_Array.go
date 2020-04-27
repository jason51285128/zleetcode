package zleetcode

func removeDuplicates (a []int) []int {
	if len(a) == 0{
		return nil
	}
	unique := 0
	for i := 1; i < len(a); i++ {
		if a[unique] != a[i] {
			unique++
			a[unique] = a[i] 
		}
	}
	unique++
	return a[0: unique]
}