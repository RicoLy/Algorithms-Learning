package insertion

import (
	"Algorithms-Learning/sort/utils"
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	list := utils.GetArrayOfSize(1000000)

	Sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func benchmarkInsertionSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)

	for i := 0; i < b.N; i++ {
		//InsertSort(list)
		Sort(list)
	}
}

func BenchmarkInsertionSort100(b *testing.B) { benchmarkInsertionSort(100, b) }
func BenchmarkInsertionSort1000(b *testing.B)   { benchmarkInsertionSort(1000, b) }
func BenchmarkInsertionSort10000(b *testing.B)  { benchmarkInsertionSort(10000, b) }
func BenchmarkInsertionSort100000(b *testing.B) { benchmarkInsertionSort(100000, b) }