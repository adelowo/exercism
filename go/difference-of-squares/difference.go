package diffsquares

const testVersion = 1

func SquareOfSums(num int) int {

	var total int
	for i := num; i > 0; i-- {
		total += i
	}

	return total * total

}

func SumOfSquares(num int) int {

	var total int

	for i := num; i > 0; i-- {
		total += i * i
	}

	return total
}

func Difference(num int) int {
	return SquareOfSums(num) - SumOfSquares(num)

}
