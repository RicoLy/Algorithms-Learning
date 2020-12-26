package bubble

import (
	"Algorithms-Learning/sort/utils"
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	Sort(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func TestSort2(t *testing.T) {
	list := utils.GetArrayOfSize(100)
	Sort2(list)
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func benchmarkBubbleSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		Sort(list)
	}
}

func benchmarkBubbleSort2(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		Sort2(list)
	}
}

func BenchmarkBubbleSort100(b *testing.B) { benchmarkBubbleSort(100, b) }
func BenchmarkBubbleSort1000(b *testing.B)   { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort100000(b *testing.B) { benchmarkBubbleSort(100000, b) }

func BenchmarkBubbleSort2_100(b *testing.B) { benchmarkBubbleSort(100, b) }
func BenchmarkBubbleSort2_1000(b *testing.B)   { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort2_10000(b *testing.B)  { benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort2_100000(b *testing.B) { benchmarkBubbleSort(100000, b) }