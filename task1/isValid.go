func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		// 如果是右括号
		if open, ok := pairs[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// 是左括号，入栈
			stack = append(stack, ch)
		}
	}
	// 最后栈必须为空
	return len(stack) == 0
}