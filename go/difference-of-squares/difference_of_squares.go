package diffsquares

func SquareOfSum(number int) int {
	result := 0
	
	for i := 1; i <= number; i++ {
		result += i
	}

	return result * result
}

func SumOfSquares(number int) (result int) {
	for i := 1; i <= number; i++ {
		result = result + (i * i)
	}

	return
}

func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}
