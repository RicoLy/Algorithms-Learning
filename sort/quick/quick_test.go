package quick

import (
	"Algorithms-Learning/sort/utils"
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	//list := []int{1,99,6,999,5,2,}
	//Sort(list)
	//HoareSort(list, 0, len(list)-1)
	// LomutoSort(list, 0, len(list)-1)
	Quick(list, 0, len(list)-1)
	//fmt.Println(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}


func benchmarkQuickSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)

	for i := 0; i < b.N; i++ {
		Quick(list, 0, len(list)-1)
		// Sort(list)LomutoSort
		// LomutoSort(list, 0, len(list)-1)
		//HoareSort(list, 0, len(list)-1)
		//QuickSort(list, 0, len(list)-1)
	}
}

func BenchmarkQuickSort100(b *testing.B) { benchmarkQuickSort(100, b) }
func BenchmarkQuickSort1000(b *testing.B)   { benchmarkQuickSort(1000, b) }
func BenchmarkQuickSort10000(b *testing.B)  { benchmarkQuickSort(10000, b) }
func BenchmarkQuickSort100000(b *testing.B) { benchmarkQuickSort(100000, b) }
