package main

type Node struct {
	data int
	next *Node
}

func newNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}
