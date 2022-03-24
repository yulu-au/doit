package bug

//交替打印 cat dog fish
import (
	"fmt"
	"sync"
)

func PrintCat(fishCH, catCH chan bool) {
	defer waitgroup.Done()
	defer close(catCH)
	for i := 0; i < 100; i++ {
		<-fishCH
		fmt.Println("cat ...")
		catCH <- true
	}
}

func PrintDog(catCH, dogCH chan bool) {
	defer waitgroup.Done()
	defer close(dogCH)
	for i := 0; i < 100; i++ {
		<-catCH
		fmt.Println("dog ...")
		dogCH <- true
	}
}

func PrintFish(dogCH, fishCH chan bool) {
	defer waitgroup.Done()
	defer close(fishCH)
	for i := 0; i < 100; i++ {
		<-dogCH
		fmt.Println("fish ...")
		fishCH <- true
	}
}

var waitgroup sync.WaitGroup

func main() {

	catCH := make(chan bool, 1)
	dogCH := make(chan bool, 1)
	fishCH := make(chan bool, 1)
	fishCH <- true

	go PrintFish(dogCH, fishCH)
	go PrintDog(catCH, dogCH)
	go PrintCat(fishCH, catCH)

	waitgroup.Add(3)
	waitgroup.Wait()

}
