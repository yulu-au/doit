package main

/*
 写出一个程序，接受一个字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）
*/

func solveA(str string) string {
	// write code here
	what := []rune(str)
	l, r := 0, len(what)-1
	for l < r {
		what[l], what[r] = what[r], what[l]
		l++
		r--
	}
	return string(what)
}
