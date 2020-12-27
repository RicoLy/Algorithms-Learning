package merge

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := MergeSort(arr[0:i])
	right := MergeSort(arr[i:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] < right[n] {
			result = append(result, left[m])
			m++
			continue
		}
		result = append(result, right[n])
		n++
	}
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}

func MergeSort1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := MergeSort1(arr[0:i])
	right := MergeSort1(arr[i:])
	return merge1(left, right)
}

func merge1(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] < right[n] {
			result = append(result, left[m])
			m++
			continue
		}
		result = append(result, right[n])
		n++
	}
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}