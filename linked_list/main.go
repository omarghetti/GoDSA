package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	lenght int
}

func (l *LinkedList) InsertLast(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		last := l.head
		for last.next != nil {
			last = last.next
		}
		last.next = newNode
	}
	l.lenght++
}

func (l *LinkedList) InsertFirst(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.lenght++
}

func (l LinkedList) Print() {
	tmp := l.head
	for tmp != nil {
		fmt.Println(tmp.value)
		tmp = tmp.next
	}
}

func (l *LinkedList) Delete(value int) error {
	if l.head == nil {
		return fmt.Errorf("cannot delete from an empty list")
	}
	if l.head.value == value {
		l.head = l.head.next
		l.lenght--
		return nil
	}
	tmp := l.head
	for tmp.next.value != value {
		if tmp.next.next == nil {
			return fmt.Errorf("value not found")
		}
		tmp = tmp.next
	}
	tmp.next = tmp.next.next
	l.lenght--
	return nil
}

func main() {
	ll := LinkedList{}
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)
	ll.InsertLast(4)
	ll.InsertFirst(0)
	ll.InsertFirst(-1)
	ll.InsertFirst(-2)
	ll.InsertFirst(-3)
	ll.InsertFirst(-4)
	err := ll.Delete(0)
	if err != nil {
		fmt.Println(err)
	}
	ll.Print()
}
