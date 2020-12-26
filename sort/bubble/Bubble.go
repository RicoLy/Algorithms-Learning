package bubble

// 冒泡排序
func Sort(arr []int) {

	for first := len(arr) - 1; ; first-- {
		//终止条件，也可通过外层循环来控制
		continued := false

		for second := first - 1; second >= 0; second-- {
			continued = true
			if arr[first] < arr[second] {
				arr[second], arr[first] = arr[first], arr[second]
			}
		}
		if continued == false {
			break
		}
	}
}

func Sort2(arr []int) {
	flag := true
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				flag = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if flag {
			return
		} else {
			flag = false
		}
	}
}