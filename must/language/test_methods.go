package language

import "fmt"

type S struct {
}

func (s *S) pointerSFunc() {
	fmt.Println("pointerSFunc")
}
func (s S) nopointerSFunc() {
	fmt.Println("nopointerSFunc")
}

type T struct {
	S
}

func (t *T) pointerTFunc() {
	fmt.Println("PointerTFunc")
}
func (t T) noPointerTFunc() {
	fmt.Println("NoPointerTFunc")
}

func TestMethods() {
	var t T
	pt := &t
	t.noPointerTFunc()
	t.pointerTFunc()
	t.nopointerSFunc()
	t.pointerSFunc()
	pt.noPointerTFunc()
	pt.nopointerSFunc()
	pt.pointerSFunc()
	pt.pointerTFunc()

	var s S
	ps := &s
	s.nopointerSFunc()
	s.pointerSFunc()
	ps.nopointerSFunc()
	ps.pointerSFunc()
}
