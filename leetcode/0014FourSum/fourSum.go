package _014FourSum

import "sort"

func FourSum(nums []int, target int) [][]int  {
	if len(nums) < 4 {
		return nil
	}

	// 生序排序
	sort.Ints(nums)
	// 遍历、收集等于targe的组合
	var equals [][]int
	for i := 0; i < len(nums)-3; i++ {
		// 第一个元素去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			// 第二个元素去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k, l := j+1, len(nums)-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				distance := sum - target
				if distance == 0 {
					equals = append(equals, []int{nums[i], nums[j], nums[k], nums[l]})
				}
				if distance > 0 {
					// 第四个元素去重
					for cur := l; l > k && nums[l] == nums[cur]; l-- {
					}
				} else {
					// 第三个元素去重
					for cur := k; k < l && nums[k] == nums[cur]; k++ {
					}
				}
			}
		}
	}

	return equals
}