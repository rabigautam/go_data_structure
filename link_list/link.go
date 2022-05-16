package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkList struct {
	head   *node
	length int
}

func (l *linkList) prepend(value int) {

	newNode := node{data: value}

	if l.head == nil {
		l.head = &newNode
		l.length++
		return
	}
	newNode.next = l.head
	l.head = &newNode
	l.length++
	return

}

func (l *linkList) printList() {
	if l.head == nil {
		fmt.Println("List empty to print")
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next

	}

}
func (l *linkList) deleteWithValue(value int) {
	if l.length == 0 {
		fmt.Println("list empty")
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		//end of list
		//check next node from current traversal
		if previousToDelete.next.next == nil {
			return
		}
		// previousToDelete.next is actually
		//so now previous Node for the next traversal will be current node
		previousToDelete = previousToDelete.next
	}
	//removing the deleted node from list and targeting to next node
	previousToDelete.next = previousToDelete.next.next

	l.length--

}
func main() {

	newList := linkList{}
	newList.prepend(2)
	newList.prepend(3)
	newList.prepend(5)
	newList.prepend(8)
	newList.prepend(13)
	newList.prepend(21)
	newList.prepend(34)

	newList.printList()
	fmt.Println("deleting 5 from list")

	newList.deleteWithValue(5)

	newList.printList()

}
