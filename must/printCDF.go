package must

import (
	"fmt"
	"sync"
)

//交替打印 cat dog fish

// 	"fmt"

// func PrintCat(fishCH, catCH chan bool) {
// 	defer waitgroup.Done()
// 	defer close(catCH)
// 	for i := 0; i < 100; i++ {
// 		<-fishCH
// 		fmt.Println("cat ...")
// 		catCH <- true
// 	}
// }

// func PrintDog(catCH, dogCH chan bool) {
// 	defer waitgroup.Done()
// 	defer close(dogCH)
// 	for i := 0; i < 100; i++ {
// 		<-catCH
// 		fmt.Println("dog ...")
// 		dogCH <- true
// 	}
// }

// func PrintFish(dogCH, fishCH chan bool) {
// 	defer waitgroup.Done()
// 	defer close(fishCH)
// 	for i := 0; i < 100; i++ {
// 		<-dogCH
// 		fmt.Println("fish ...")
// 		fishCH <- true
// 	}
// }

// var waitgroup sync.WaitGroup

// func Bug02() {

// 	catCH := make(chan bool, 1)
// 	dogCH := make(chan bool, 1)
// 	fishCH := make(chan bool, 1)
// 	fishCH <- true

// 	go PrintFish(dogCH, fishCH)
// 	go PrintDog(catCH, dogCH)
// 	go PrintCat(fishCH, catCH)

// 	waitgroup.Add(3)
// 	waitgroup.Wait()

// }
func PrintDCF() {
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

	for i := 0; i < 100; i++ {
		<-a
		fmt.Println(what)
		b <- true
	}
}
