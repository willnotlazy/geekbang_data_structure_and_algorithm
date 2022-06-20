package sort

// QuickSort 快速排序
func QuickSort(target []int, start, end int) {
	if start >= end {
		return
	}

	mid := QuickSortMid(target, start, end)
	QuickSort(target, start, mid-1)
	QuickSort(target, mid+1, end)
}

func QuickSortMid(target []int, start, end int) int {
	mid_value := target[end]
	i, j := start-1, start

	for j < end {
		if target[j] <= mid_value {
			i++
			target[i], target[j] = target[j], target[i]
		}
		j++
	}

	target[i+1], target[end] = target[end], target[i+1]

	return i + 1
}
