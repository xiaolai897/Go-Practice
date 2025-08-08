func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	data := [][]int{intervals[0]}
	for i := 1; i < n; i++ {
		last := data[len(data)-1]
		current := intervals[i]
		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			data = append(data, current)
		}
	}
	return data
}