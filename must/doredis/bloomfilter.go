package doredis

import (
	"fmt"

	"github.com/willf/bloom"
)

func BloomTest() {
	filter := bloom.NewWithEstimates(1000000, 0.01)
	filter.AddString("love1")
	filter.AddString("love2")
	filter.AddString("love3")
	filter.AddString("love4")
	filter.AddString("love5")
	filter.AddString("love6")
	filter.AddString("love7")
	filter.AddString("love8")
	filter.AddString("love9")

	for i := 10; i < 100; i++ {
		yes := filter.TestString(fmt.Sprintf("love%v", i))
		if yes {
			fmt.Println("this bloom not ok")
		}
	}
}
