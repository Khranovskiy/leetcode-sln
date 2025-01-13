func eraseOverlapIntervals(intervals [][]int) int {
	// [[1,2],[2,3],[3,4],[1,3]]
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
    [[1,2],[2,3],[1,3],[3,4]]
	res := 0
	k := intervals[0][1] // 2
	for idx := 1; idx < len(intervals); idx++ {
        // [1,3]
		end := intervals[idx][1]
		start := intervals[idx][0]

        //  1 >= 2
		if start >= k {
			k = end
		} else {
			res++ // 1
		}
	}
	return res
}
