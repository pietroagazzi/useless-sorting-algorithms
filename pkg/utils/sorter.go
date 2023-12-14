package utils

type Sorter func([]int) []int

func AdaptSlowsort(f func([]int, int, int) []int) Sorter {
	return func(arr []int) []int {
		return f(arr, 0, len(arr)-1)
	}
}

// AdaptSleepsort adapts a function that works on uint slices to work on int slices
func AdaptSleepsort(f func([]uint) []uint) Sorter {
	return func(arr []int) []int {
		uintArr := make([]uint, len(arr))
		for i, v := range arr {
			uintArr[i] = uint(v)
		}

		result := f(uintArr)

		intArr := make([]int, len(result))
		for i, v := range result {
			intArr[i] = int(v)
		}

		return intArr
	}
}
