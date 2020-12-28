package _004FindMedianSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1) + len(nums2) + 1)>>2, 0, 0

	for low <= high {
		nums1Mid = low + (high-low) >>1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			high = nums1Mid -1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			low = nums2Mid + 1
		} else {
			break
		}
	}
	midleft, midRigh := 0, 0
	if nums1Mid == 0 {
		midleft = nums2[nums2Mid]
	} else if nums2Mid == 0 {
		midleft = nums1[nums1Mid-1]
	} else {
		midleft = max(nums1[nums1Mid-1], nums2[nums2Mid -1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midleft)
	}
	if nums1Mid == len(nums1) {
		midRigh = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRigh = nums1[nums1Mid]
	} else {
		midRigh = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midleft+midRigh) / 2
}


func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}