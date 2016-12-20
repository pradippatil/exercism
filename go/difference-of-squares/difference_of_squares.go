package diffsquares

// SquareOfSums returns square of sum of first n numbers
// Formula for sum of first n numbers : n(n+1)/2
func SquareOfSums(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)

}

// SumOfSquares returns sum of squares first n numbers
// Formula: n(n+1)(2n+1)/6
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns difference between SquareOfSums and SumOfSquares for given number n
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)

}
