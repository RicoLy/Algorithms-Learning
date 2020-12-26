package shell

func insertion(arr []int, start, gap int) {
	length := len(arr)
	for traverseVal := start + gap; traverseVal < length; traverseVal += gap {
		backup := arr[traverseVal]
		trackVal := traverseVal - gap
		for trackVal >= 0 && backup < arr[trackVal] {
			arr[trackVal + gap] = arr[trackVal]
			trackVal -= gap
		}
		arr[trackVal + gap] = backup
	}
}

func Sort(arr []int) {
	gap := len(arr) / 2
	for gap > 0 {
		for pos := 0; pos < gap; pos++ {
			insertion(arr, pos, gap)
		}
		gap /= 2
	}
}

func ShellSort(arr []int) {
	//根据增量（len(arr)/2）每次变成上一次的1/2
	for inc := len(arr) / 2; inc > 0; inc /= 2 {

		for i := inc; i < len(arr); i++ {
			temp := arr[i]

			//根据增量从数据到0进行比较
			for j := i - inc; j >= 0; j -= inc {
				//满足条件交换数据
				if temp < arr[j] {
					arr[j], arr[j+inc] = arr[j+inc], arr[j]
				} else {
					break
				}
			}
		}
	}
}
