fibonacciSize :: (Integral a) => a -> a -> a
fibonacciSize size n
	| n <= size = 1
	| otherwise = sum [fibonacciSize size x | x <- [(n-size)..(n-1)]]
