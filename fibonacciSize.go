func fibonacciSize(size int, n int) int {
	if n<=size {
		return 1
	}
	sum := 0
	for i:=n-size; i<n; i++ {
	  	sum += fibonacciSize(size, i)
	}
	return sum
}
