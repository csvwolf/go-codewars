package kata

func SumCubes(n int) int {
    result := 0
    for i := 1; i <= n; i++ {
        result += i * i * i
    }
    return result
}
