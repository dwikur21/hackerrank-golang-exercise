package main

import (
	"fmt"
	"sync"

	"golang.org/x/tour/tree"
)

type stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []tree.Tree
}

func newStack() *stack {
	return &stack{sync.Mutex{}, make([]tree.Tree, 0)}
}

func (s *stack) Push(v tree.Tree) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Pop() tree.Tree {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	// if l == 0 {
	// 	return nil, errors.New("Empty Stack")
	// }

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	st := newStack()
	st.Push(*t)

	for len(st.s) > 0 {
		node := st.Pop()
		ch <- node.Value

		if node.Right != nil {
			st.Push(*node.Right)
		}
		if node.Left != nil {
			st.Push(*node.Left)
		}
	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	sum1 := 0
	sum2 := 0

	go Walk(t1, c1)
	go Walk(t2, c2)

	for value := range c1 {
		sum1 += value
	}

	for value := range c2 {
		sum2 += value
	}

	return (sum1 == sum2)
}

func main() {
	//start := time.Now()

	t1 := tree.New(121)
	t2 := tree.New(121)

	//elapsed := time.Since(start)
	fmt.Println(Same(t1, t2))
}
