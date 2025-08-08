func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始向前处理
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits // 不需要进位，直接返回
		}
		// 当前位是9，加一变成0，需要进位
		digits[i] = 0
	}

	// 如果全部都是9，比如 999 -> 1000
	// 需要在最前面插入1，其他位默认是0
	return append([]int{1}, digits...)
}
