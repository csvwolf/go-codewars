package kata

func Factorial(n int) int {
    factorial := 1
    for i := 1; i <= n; i++ {
        factorial *= i
    }
    return factorial
}