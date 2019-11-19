/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 *
 * https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (33.32%)
 * Likes:    630
 * Dislikes: 380
 * Total Accepted:    180K
 * Total Submissions: 538.8K
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * Evaluate the value of an arithmetic expression in Reverse Polish Notation.
 * 
 * Valid operators are +, -, *, /. Each operand may be an integer or another
 * expression.
 * 
 * Note:
 * 
 * 
 * Division between two integers should truncate toward zero.
 * The given RPN expression is always valid. That means the expression would
 * always evaluate to a result and there won't be any divide by zero
 * operation.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: ["2", "1", "+", "3", "*"]
 * Output: 9
 * Explanation: ((2 + 1) * 3) = 9
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ["4", "13", "5", "/", "+"]
 * Output: 6
 * Explanation: (4 + (13 / 5)) = 6
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
 * Output: 22
 * Explanation: 
 * ⁠ ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 * 
 * 
 */
func evalRPN(tokens []string) int {
	var stack []int
	pop := func() int {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return r
	}
	push := func(v int) {
		stack = append(stack, v)
	}
	for _, token := range tokens {
		var op func(a,b int) int
		switch token {
		case "+":
			op = func(a,b int)int {
				return a+b
			}
		case "-":
			op = func(a,b int)int {
				return a-b
			}
		case "*":
			op = func(a,b int)int {
				return a*b
			}
		case "/":
			op = func(a,b int)int {
				return a/b
			}
		default:
			v, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			push(v)
			continue
		}

		a := pop()
		b := pop()
		push(op(b,a))
	}
	return pop()
}

