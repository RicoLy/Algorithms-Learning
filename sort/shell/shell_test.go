package shell

import (
	"Algorithms-Learning/sort/utils"
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)

	Sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func TestShellSort2(t *testing.T) {
	list := utils.GetArrayOfSize(100)

	ShellSort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func benchmarkShellSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)

	for i := 0; i < b.N; i++ {
		ShellSort(list)
	}
}

func BenchmarkShellSort100(b *testing.B) { benchmarkShellSort(100, b) }
func BenchmarkShellSort1000(b *testing.B)   { benchmarkShellSort(1000, b) }
func BenchmarkShellSort10000(b *testing.B)  { benchmarkShellSort(10000, b) }
func BenchmarkShellSort100000(b *testing.B) { benchmarkShellSort(100000, b) }