package leetcode

import (
	"strconv"
)

func pop(stack *[]int) int {
	ret := (*stack)[len((*stack))-1]
	*stack = (*stack)[:len((*stack))-1]
	return ret
}

func evalRPN(tokens []string) int {
	var stack []int

	for i := 0; i < len(tokens); i++ {
		if val, err := strconv.Atoi(tokens[i]); err == nil {
			stack = append(stack, val)
		} else {
			var newVal int
			val1 := pop(&stack)
			val2 := pop(&stack)
			switch tokens[i] {
			case "+":
				newVal = val2 + val1
			case "-":
				newVal = val2 - val1
			case "*":
				newVal = val2 * val1
			case "/":
				newVal = val2 / val1
			}

			stack = append(stack, newVal)
		}
	}
	retVal := pop(&stack)
	return retVal
}

func main() {
	test := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	print(evalRPN(test))
}
