func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    var ans [][]int
    pre := intervals[0]
    for i := 1; i < len(intervals); i++ {
        cur := intervals[i]
        if pre[1] < cur[0] {
            ans = append(ans, pre)
            pre = cur
        } else {
            pre[1] = max(pre[1], cur[1])
        }
    }
    ans = append(ans, pre)
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
