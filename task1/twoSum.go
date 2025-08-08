func twoSum(nums []int, target int) []int {
	data := map[int]int{}
	for k, v := range nums {
		if i, ok := data[target-v]; ok {
			return []int{i, k}
		} else {
			data[v] = k
		}
	}
	return nil
}