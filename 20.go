package main

func isValid(s string) bool {
	mp := make(map[rune]rune, 3)
	mp['}'] = '{'
	mp[')'] = '('
	mp[']'] = '['
	stack := make([]rune, 0)
	for _, x := range s {
		if x == '{' || x == '[' || x == '(' {
			stack = append(stack, x)
		} else {
			tmp := mp[x]
			for len(stack) > 0 && stack[len(stack)-1] != tmp {
        return false
			}
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
  if len(stack) != 0 {
    return false
  }
  return true
}
