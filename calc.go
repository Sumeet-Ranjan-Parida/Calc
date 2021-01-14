package calc

import "fmt"

func calculate() {
	var num1, num2 int
	var operator string
	result := 0

	fmt.Printf("\nEnter the First Number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the Operator (+, -, *, /, %): ")
	fmt.Scanln(&operator)
	fmt.Print("Enter the Second Number: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		result = num1 + num2
		break
	case "-":
		result = num1 - num2
		break
	case "*":
		result = num1 * num2
		break
	case "/":
		result = num1 / num2
		break
	case "%":
		result = num1 % num2
		break
	default:
		fmt.Println("Invalid Operation")
	}

	fmt.Printf("\n%d %s %d = %d", num1, operator, num2, result)
}
