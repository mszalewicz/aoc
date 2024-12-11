package main

type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
}

func (list *LinkedList[T]) Add(val T) {
	newNode := &Node[T]{Value: val}
	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		previous := list.Head
		for current.Next != nil {
			previous = current
			current = current.Next
		}
		current.Next = newNode
		current.Previous = previous
	}
}

// func (list *LinkedList[T]) Insert(T)
