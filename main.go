package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var (
	Head *Node
)

func AddElemsToList(s []int) *Node {
	var (
		pprev *Node = nil
		p     *Node = nil
	)
	for i := len(s) - 1; i >= 0; i-- {
		p = new(Node)
		p.Value = s[i]
		p.Next = pprev
		pprev = p
	}
	return p
}

func ReversList(h *Node) *Node {
	var (
		hp *Node = nil
		hn *Node = nil
	)
	for h.Next != nil {
		hn = h.Next
		h.Next = hp
		hp = h
		h = hn
	}
	h.Next = hp
	return h
}

func PrintList(h *Node) {
	c := h
	fmt.Println("List:")
	for c != nil {
		fmt.Println(c)
		c = c.Next
	}
}

func main() {
	// var p *Node
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	Head = AddElemsToList(s)
	PrintList(Head)
	Head = ReversList(Head)
	PrintList(Head)
}
