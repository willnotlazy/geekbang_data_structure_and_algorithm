package sort

import (
	"math"
)

// CountingSort 计数排序，这里假设所有数据均为非负整数
func CountingSort(_sort []int) []int {
	if len(_sort) <= 0 {
		return _sort
	}

	var _max = math.MinInt

	for i := 0; i < len(_sort); i++ {
		if _sort[i] > _max {
			_max = _sort[i]
		}
	}
	_bucket := make([]int, _max+1)
	for _, number := range _sort {
		_bucket[number]++
	}

	for i := 1; i <= _max; i++ {
		_bucket[i] = _bucket[i] + _bucket[i-1]
	}

	_sort_res := make([]int, len(_sort))

	for i := len(_sort) - 1; i >= 0; {
		num := _sort[i]
		index := _bucket[num] - 1
		_bucket[num] = _bucket[num] - 1
		_sort_res[index] = _sort[i]
		i = i - 1
	}

	return _sort_res
}
