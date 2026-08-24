package main

import "fmt"

// Node represents one node of the linked list
type Node struct {
	data int
	next *Node
}

// newNode creates and returns a new node
func newNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

func main() {

	// Create first node
	node1 := newNode(10)

	// Initially both head and tail point to node1
	head := node1
	tail := node1

	print(head)

	// Insert at tail
	tail = insertAtTail(tail, 12)
	print(head)

	tail = insertAtTail(tail, 15)
	print(head)

	// Current List:
	// 10 -> 12 -> 15

	// Insert 11 at position 2
	head, tail = insertAtPosition(head, tail, 2, 11)
	print(head)

	// Insert 5 at position 1
	head, tail = insertAtPosition(head, tail, 1, 5)
	print(head)

	// Insert 20 at position 6 (last)
	head, tail = insertAtPosition(head, tail, 6, 20)
	print(head)
}

// Insert at beginning
func insertAtHead(head *Node, data int) *Node {

	temp := newNode(data)

	temp.next = head

	head = temp

	return head
}

// Insert at end
func insertAtTail(tail *Node, data int) *Node {

	temp := newNode(data)

	tail.next = temp

	tail = temp

	return tail
}

// Insert at any position
func insertAtPosition(
	head *Node,
	tail *Node,
	position int,
	data int,
) (*Node, *Node) {

	// Case 1: Insert at beginning
	if position == 1 {

		head = insertAtHead(head, data)

		return head, tail
	}

	// Start from head
	temp := head

	count := 1

	// Move temp to the node BEFORE
	// the position where we want to insert
	for count < position-1 {
		temp = temp.next
		count++
	}

	// Case 2: Insert at end
	if temp.next == nil {

		tail = insertAtTail(tail, data)

		return head, tail
	}

	// Case 3: Insert somewhere in the middle

	nodeToInsert := newNode(data)

	nodeToInsert.next = temp.next

	temp.next = nodeToInsert

	return head, tail
}

// Print complete linked list
func print(head *Node) {

	if head == nil {
		fmt.Println("List is empty")
		return
	}

	temp := head

	for temp != nil {

		fmt.Print(temp.data, " ")

		temp = temp.next
	}

	fmt.Println()
}
