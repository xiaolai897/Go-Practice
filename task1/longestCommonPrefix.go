func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	base := strs[0]
	fmt.Printf("base: %s\n", base)
	for i := range base {
		for _, s := range strs[1:] {
			if i >= len(s) || s[i] != base[i] {
				return base[:i]
			}
		}
	}
	return base
}