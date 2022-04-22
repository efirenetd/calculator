package calculator

var logMessage = "[LOG]"

// Version of the calculator
// The Version variable can reach from anywhere (uppercase V in the name of a variable equivalent to "public", in other language)
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

func Sum(number1, number2 int) int {
	return number1 + number2
}
