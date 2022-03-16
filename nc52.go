package main

/*
给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。
*/

//golang里面的字符类型是byte
func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, v := range []byte(s) {
		if len(stack) != 0 {
			if ok(stack[len(stack)-1], v) {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		} else {
			stack = append(stack, v)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func ok(a, b byte) bool {
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	if v, et := m[a]; et && v == b {
		return true
	}
	return false
}
