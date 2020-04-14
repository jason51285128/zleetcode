package zleetcode

//子数组和最接近目标，但是不能大于目标，数组元素和目标都是正整数
//头尾指针维护一个数组，如果下一个数能更近目标，则移动头指针，如果已经超过目标则移动尾指针

func closestSubarraySum(a []int, target int) []int {
	closest := 0
	start := 0
	end := 0
	sum := 0
	tail := -1
	head := 0
	v := 0
	//isexist := false
	for head, v = range a {
		if v == target {
			return []int{target, head, head}
		}
		if target-sum >= v {
			sum += v
			if tail == -1 {
				tail = head
			}
			continue
		}
		if tail != -1 {
			if sum > closest {
				closest = sum
				start = tail
				end = head - 1
			}
			if v > target {
				sum = 0
				tail = -1
				continue
			}
			for tail < head {
				sum -= a[tail]
				tail++
				if target-sum > v {
					sum += v
					break
				}
			}
		}
	}
	if sum > closest {
		closest = sum
		start = tail
		end = head
	}
	if closest != 0 {
		return []int{closest, start, end}
	}
	return []int{-1, -1, -1}
}
