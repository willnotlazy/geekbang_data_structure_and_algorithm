package sort

// MergeSort 归并排序
func MergeSort(_sort []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)>>1
	MergeSort(_sort, start, mid)
	MergeSort(_sort, mid+1, end)
	Merge(_sort, start, mid, end)
}

func Merge(_sort []int, start, mid, end int) {
	p1 := start
	p2 := mid + 1
	i := 0
	tmp := make([]int, end-start+1)

	for p1 <= mid && p2 <= end {
		if _sort[p1] > _sort[p2] {
			tmp[i] = _sort[p2]
			p2++
		} else {
			tmp[i] = _sort[p1]
			p1++
		}
		i++
	}

	for p1 <= mid {
		tmp[i] = _sort[p1]
		i++
		p1++
	}
	for p2 <= end {
		tmp[i] = _sort[p2]
		i++
		p2++
	}

	for i := 0; i < end-start+1; i++ {
		_sort[start+i] = tmp[i]
	}
}
