package insertion

func Sort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}

	for i := 1; i < length-1; i++ {
		backup := arr[i]
		j:=i-1
		for j >= 0 && backup < arr[j]  {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = backup
	}
}

func InsertSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		//如果当前数据小于有序数据
		if arr[i] < arr[i-1] {
			j := i - 1
			//获取有效数据
			temp := arr[i]
			//一次比较有序数据
			for j >= 0 && arr[j] > temp {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = temp
		}
	}
}

func InsertSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i-1
		for ; j >= 0 && arr[j] < temp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}