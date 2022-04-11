package language

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Jeffail/tunny"
)

//利用chan来限制同时运行的goroutie数量
func LimitGoroutine1() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}

//利用第三方开源代码限制同时运行的goroutine
func LimitGoroutine2() {
	pool := tunny.NewFunc(3, func(i interface{}) interface{} {
		val, _ := i.(int)
		log.Println(i)
		time.Sleep(time.Second)
		return val * val
	})
	defer pool.Close()

	for i := 0; i < 10; i++ {
		//捕获局部变量
		i := i
		go func() {
			r := pool.Process(i)
			log.Printf("we catch %v\n", r)
		}()
	}
	time.Sleep(time.Second * 4)
}

//超时控制的例子
//重点是ch缓冲区的分配和子goroutine里对于异常情况的兼容
func TimeOut01() {
	//这里一定要分配缓冲区
	ch := make(chan struct{}, 1)
	go doSomeThing(ch)

	select {
	//等待从ch里返回数据
	case <-ch:
		log.Println("usual ")
	case <-time.After(time.Millisecond * 10):
		log.Println("unusual")
	}
}

func doSomeThing(ch chan struct{}) {
	//do some thing
	time.Sleep(time.Second)
	/*
		为啥要这么写,试想ch是容量0的chan&&对面没有接收动作,那ch<-就会永远阻塞下去
		那么使用select的话,就算是这种情况,也不会阻塞
	*/
	select {
	case ch <- struct{}{}:
		log.Println("done from dosomething")
	//如果上面向chan发数据不成功,依旧可以返回
	default:
		log.Println("father goroutine is dead")
		return
	}
}

/*
超时控制例子2
任务被切分成两份,每一份都有timeout处理
*/
func Timeout02() {
	//这里使用0容量的chan是因为在子goroutine里有特殊处理
	phase1 := make(chan struct{})
	phase2 := make(chan struct{})
	go do2phases(phase1, phase2)

	select {
	case <-phase1:
		log.Println("we recv phase1 done")
	case <-time.After(time.Second * 2):
		log.Println("time out because phase1")
		return
	}

	select {
	case <-phase2:
		fmt.Println("we recv phase2 done")
	case <-time.After(time.Second * 2):
		log.Println("time out because phase2")
		return
	}

	log.Println("all things done")
}

func do2phases(ch1, ch2 chan struct{}) {
	//mock do some thing
	time.Sleep(time.Second)
	select {
	case ch1 <- struct{}{}:
		log.Println("phase1 done")
	default:
		log.Println("phase1 done but father dead")
		return
	}
	//mock
	time.Sleep(time.Second * 3)
	select {
	case ch2 <- struct{}{}:
		log.Println("phase2 done")
	default:
		log.Println("phase2 done but father dead")
		return
	}
}

/*
channel
操作 	   空值(nil) 	非空已关闭 	非空未关闭
关闭 	   panic 	    panic 	  成功关闭
发送数据 	永久阻塞 	  panic 	阻塞或成功发送
接收数据 	永久阻塞 	  永不阻塞 	 阻塞或者成功接收
*/

//关闭channel的例子
/*

   情形一：M个接收者和一个发送者，发送者通过关闭用来传输数据的通道来传递发送结束信号
   情形二：一个接收者和N个发送者，此唯一接收者通过关闭一个额外的信号通道来通知发送者不要再发送数据了
   情形三：M个接收者和N个发送者，它们中的任何协程都可以让一个中间调解协程帮忙发出停止数据传送的信号

*/
type safeChan struct {
	o  sync.Once
	ch chan int
}

func (s *safeChan) init() {
	s.ch = make(chan int, 10)
}

//保证只关闭一次
func (s *safeChan) close() {
	close(s.ch)
}

func HowSafeCloseChan() {
	var ch safeChan
	//
	ch.close()
}
