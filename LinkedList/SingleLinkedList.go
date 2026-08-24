package main

import "fmt"

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

func main() {

	node1 := newNode(10)

	head := node1
	tail := node1

	print(head)

	insertAtTail(&tail, 12)
	print(head)

	insertAtTail(&tail, 15)
	print(head)

	// Current:
	// 10 -> 12 -> 15

	insertAtHead(&head, 5)
	print(head)

	// Current:
	// 5 -> 10 -> 12 -> 15

	insertAtPosition(&head, &tail, 3, 11)
	print(head)

	// Current:
	// 5 -> 10 -> 11 -> 12 -> 15

	//deleteNode(3, &head, &tail)
	//
	reverseLinkedList(&head)
	print(head)
}

func reverseLinkedList(head **Node) {
	if *head == nil || (*head).next == nil {
		return
	}

	var prev *Node
	curr := *head
	for curr != nil {
		forward := curr.next
		curr.next = prev
		prev = curr
		curr = forward
	}

	*head = prev //head variable stores the address of prev
}

// -----------------------------------
// INSERT AT HEAD
// -----------------------------------

func insertAtHead(head **Node, data int) {

	temp := newNode(data)

	// temp should point to the current first node
	temp.next = *head

	// change the ORIGINAL head
	*head = temp
}

// -----------------------------------
// INSERT AT TAIL
// -----------------------------------

func insertAtTail(tail **Node, data int) {

	temp := newNode(data)

	// Current last node points to new node
	(*tail).next = temp

	// Change ORIGINAL tail
	*tail = temp
}

// -----------------------------------
// INSERT AT POSITION
// -----------------------------------

func insertAtPosition(
	head **Node,
	tail **Node,
	position int,
	data int,
) {

	// Insert at beginning
	if position == 1 {
		insertAtHead(head, data)
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
		insertAtTail(tail, data)
		return
	}

	// Insert in middle
	nodeToInsert := newNode(data)

	nodeToInsert.next = temp.next

	temp.next = nodeToInsert
}

// -----------------------------------
// PRINT
// -----------------------------------

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

func deleteNode(position int, head **Node, tail **Node) {

	if position < 1 || *head == nil {
		fmt.Println("invalid position")
		return
	}

	// delete the head
	if position == 1 {
		temp := *head
		*head = temp.next
		temp.next = nil

		if *head == nil { // list is now empty
			*tail = nil
		}
		return
	}

	// walk to the node just before `position`
	curr := *head
	var prev *Node
	cnt := 1

	for cnt < position {
		if curr == nil {
			fmt.Println("position out of range")
			return
		}
		prev = curr
		curr = curr.next
		cnt++
	}

	if curr == nil {
		fmt.Println("position out of range")
		return
	}

	prev.next = curr.next

	if curr == *tail { // we deleted the last node
		*tail = prev
	}

	curr.next = nil
}
