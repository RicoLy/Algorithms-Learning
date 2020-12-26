package quick

func QuickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start + end)/2] //基准值
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			QuickSort(arr, start, j)
		}
		if end > i {
			QuickSort(arr, i, end)
		}
	}
}

func Quick(list []int, start, end int) {
	if start < end {
		i, j := start, end
		key := list[(i+j)/2]
		for i <= j {
			for list[i] < key {
				i++
			}
			for list[j] > key {
				j--
			}
			if i <= j {
				list[i], list[j] = list[j], list[i]
				i++
				j--
			}
		}
		if start < j {
			Quick(list, start, j)
		}
		if end > i {
			Quick(list, i, end)
		}

	}
}

func Sort(list []int) []int{

	if len(list) <= 1 {
		return list
	}

	//[[----递归模板区

	pivot := list[len(list)/2]

	partitionFunc :=  func (arr []int, pivot int) ([]int, []int, []int){//基点分割函数
		lessArr := make([]int, 0)
		greaterArr := make([]int, 0)
		equalArr := make([]int, 0)

		for _, value := range arr {
			switch {
			case value < pivot:
				lessArr = append(lessArr,value)
			case value > pivot:
				greaterArr = append(greaterArr,value)
			default:
				equalArr = append(equalArr,value)
			}
		}
		return lessArr, equalArr, greaterArr
	}
	less, equal, greater := partitionFunc(list, pivot)
	//--------------]]

	copy(list, Sort(less))
	copy(list[len(less):], equal)
	copy(list[(len(less)+len(equal)):], Sort(greater))

	return list
}