func singleNumber(nums []int) int {
	owq := make(map[int]int)
	for _, v := range nums {
		owq[v] ++
	}
	for k, v := range owq {
		if v == 1 {
			return k
		}
	}
	return 1
}