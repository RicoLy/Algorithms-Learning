package _014FourSum

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	nums []int
	target int
}

type ans1 struct {
	one [][]int
}


func Test_Problem1(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{3, 2, 4, 55, 3, 4, 55, 22, 44, 22}, 6},
			ans1{[][]int{}},
		},

		{
			para1{[]int{3, 2, 4, 55, 3, 4, 55, 22, 44, 22}, 5},
			ans1{[][]int{}},
		},

		{
			para1{[]int{3, 2, 4, 55, 3, 4, 55, 1, 22, 44, 22}, 10},
			ans1{[][]int{{1,2,3,4}}},
		},

		{
			para1{[]int{3, 2, 4, 55, 3, 4, 55, 22, 44, 22}, 1},
			ans1{[][]int{}},
		},

		{
			para1{[]int{3, 2, 4, 55, 3, 4, 55, 22, 44, 22}, 5},
			ans1{[][]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf("【input】:%v       【output】:%v\n", p, FourSum(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}

