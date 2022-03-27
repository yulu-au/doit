package main

/*
 以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回。

数据范围：s.length,t.length≤100000s.length,t.length \le 100000s.length,t.length≤100000，字符串仅由'0'~‘9’构成
要求：时间复杂度 O(n)O(n)O(n)
*/

func solve(s string, t string) string {
	var res []byte
	//进位
	extra := 0
	slia := []byte(s)
	slib := []byte(t)
	lax := len(slia) - 1
	lbx := len(slib) - 1
	//0的值
	base := int('0')
	for lax >= 0 && lbx >= 0 {
		a, b := int(slia[lax]), int(slib[lbx])
		//每次相加得考虑到上一回的进位
		c := a - base + b - base + extra
		//进位要刷新
		extra = c / 10
		//本位的数值
		x := c % 10
		res = append(res, byte(x+base))
		lax--
		lbx--
	}

	if lax >= 0 {
		for lax >= 0 {
			a := int(slia[lax])
			x := a - base + extra
			extra = x / 10
			x = x % 10
			res = append(res, byte(x+base))
			lax--
		}
	}

	if lbx >= 0 {
		for lbx >= 0 {
			a := int(slib[lbx])
			x := a - base + extra
			extra = x / 10
			x = x % 10
			res = append(res, byte(x+base))
			lbx--
		}
	}
	//最终可能还有进位
	if extra >= 1 {
		res = append(res, byte(extra+base))
	}

	//匿名函数翻转字节数组
	func(s []byte) {
		l, r := 0, len(s)-1
		for l < r {
			s[l], s[r] = s[r], s[l]
			l++
			r--
		}
	}(res)

	// fc(res)

	return string(res)
}
