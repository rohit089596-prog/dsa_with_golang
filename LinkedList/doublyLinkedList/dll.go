package main

import "fmt"

func main() {

	node := newNode(10)
	head := node
	tail := node
	insertAtHead(12, &head)
	insertAtHead(15, &head)
	print(head)
	insertAtTail(25, &tail)
	print(head)
	insertAtPosition(&head, &tail, 3, 11)
	print(head)

	fmt.Println(getLength(head))
}

func newNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
		prev: nil,
	}
}

func insertAtTail(data int, tail **Node) {
	temp := newNode(data)
	(*tail).next = temp
	temp.prev = *tail
	*tail = temp
}
func insertAtHead(data int, head **Node) {
	node := newNode(data)
	node.next = *head
	(*head).prev = node
	*head = node

}

// Node represents one node of the linked list
type Node struct {
	data int
	next *Node
	prev *Node
}

func print(head *Node) {

	if head == nil {
		fmt.Println("List is empty ")
	}

	temp := head

	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}

func getLength(head *Node) int {
	len := 0
	if head == nil {
		fmt.Println("List is empty")
	}
	temp := head
	for temp != nil {
		len++
		temp = temp.next
	}

	return len
}

func insertAtPosition(
	head **Node,
	tail **Node,
	position int,
	data int,
) {

	// Insert at beginning
	if position == 1 {
		insertAtHead(data, head)
		return
	}

	// Get the actual node pointed to by head
	temp := *head

	count := 1

	// Reach node just BEFORE insertion position
	for count < position-1 {
		temp = temp.next
		count++
	}

	// Insert at end
	if temp.next == nil {
		insertAtTail(data, tail)
		return
	}

	// Insert in middle
	nodeToInsert := newNode(data)

	nodeToInsert.next = temp.next
	temp.next.prev = nodeToInsert
	temp.next = nodeToInsert
	nodeToInsert.prev = temp
}
