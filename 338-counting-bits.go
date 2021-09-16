func countBits(n int) []int {
    ans := make([]int, n+1)
    ans[0] = 0
    for i := 1; i <= n; i++ {
        ans[i] = ans[i >> 1] + i % 2
    }
    
    return ans   
}
