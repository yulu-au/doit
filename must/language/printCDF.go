/*
https://www.v2ex.com/t/555049#reply0
同步 做事的先后顺序
互斥 多线程访问资源的冲突
*/
package language

import (
	"fmt"
	"sync"
)

/*
交替打印 cat dog fish

伪代码

def echo(threadNum, Upstream, Downstream):
  for i in range(M):
    wait Upstream  // 等待上游的信号
    print(threadNum)
    signal Downstream // 给下游发送信号

*/

func printDCF() {
	var wg sync.WaitGroup
	d, c, f := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1)

	go printThing("dog", f, d, &wg)
	go printThing("cat", d, c, &wg)
	go printThing("fish", c, f, &wg)
	f <- true

	wg.Add(3)
	wg.Wait()
}

func printThing(what string, a, b chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		<-a
		fmt.Println(what)
		b <- true
	}
}

// var (
// 	N = 5
// 	M = 5
// )

// func gen(v string, times int) <-chan string {
// 	ch := make(chan string)
// 	go func() {
// 		defer close(ch)
// 		for i := 0; i < times; i++ {
// 			ch <- v
// 		}
// 	}()
// 	return ch
// }

// func fanInCDF(times int, inputs []<-chan string) <-chan string {
// 	ch := make(chan string)
// 	go func() {
// 		defer close(ch)
// 		for i := 0; i < times; i++ {
// 			for _, input := range inputs {
// 				v := <-input
// 				ch <- v
// 			}
// 		}
// 	}()
// 	return ch
// }

// func main() {
// 	times := M
// 	inputs := make([]<-chan string, 0, N)
// 	for i := 0; i < N; i++ {
// 		threadName := string('A' + i)
// 		inputs = append(inputs, gen(threadName, times))
// 	}
// 	for char := range fanInCDF(times, inputs) {
// 		fmt.Println(char)
// 	}
// }
