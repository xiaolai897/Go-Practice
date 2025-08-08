func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// i 是慢指针，j 是快指针
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1 // 长度是索引 + 1
}