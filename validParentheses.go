package main

func validParentheses(s string) bool {
	pointer := -1
	pair := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	stack := []string{}

	for _, i := range s {
		if string(i) == "{" || string(i) == "[" || string(i) == "(" {
			stack = append(stack, string(i))
			pointer += 1
		} else {
			if pointer == -1 {
				return false
			}
			if stack[pointer] != pair[string(i)] {
				return false
			}
			stack = stack[:pointer]
			pointer -= 1
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true

}
