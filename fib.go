package fib

func FibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

func FibRecursiveCache(n int) int {
	cache := make([]int, n+1, n+1)
	fibRecursiveCache(n, &cache)
	return cache[n]
}

func fibRecursiveCache(n int, cache *[]int) {
	if n < 2 {
		(*cache)[0] = 0
		(*cache)[1] = 1
		return
	}
	fibRecursiveCache(n-1, cache)

	(*cache)[n] = (*cache)[n-1] + (*cache)[n-2]
}
