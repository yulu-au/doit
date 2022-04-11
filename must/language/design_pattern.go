package language

import "sync"

//单例模式

// type singleton struct{}

// var instance *singleton

// func GetInstance() *Singleton {
// 	if instance == nil {
// 		instance = &Singleton{} // 不是并发安全的
// 	}
// 	return instance
// }
/*

type Once struct {
	done uint32
	m    Mutex
}


func (o *Once) Do(f func()) {
	//判断这个单例是不是第一次运行,不是的话直接返回
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	//加锁而不是CAS的原因是保证Do函数返回的时候f函数一定已经执行完成了
	//CAS的话,没有抢到原子操作的goroutine会直接返回,可抢到的goroutine还没执行完f函数
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		//为什么要defer执行,思考一下f中panic的情况
		//就算f报panic也算一次执行f
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

*/
type Singleton struct {
	Str string
}

var Instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		Instance = &Singleton{Str: "panic"}
		panic("i am panic")
	})
	return Instance
}
