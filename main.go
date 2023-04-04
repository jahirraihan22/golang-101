package main

import "fmt"

type Node struct {
	data interface {}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) Add(data interface{}){
	// newNode := &Node{data, list.head}
    // list.head = newNode
    // list.size++

	newNode := &Node{data: data}

    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }

    list.size++
}

func (list *LinkedList) Traverse() {
    currentNode := list.head
    // fmt.Println(list.size)
    for currentNode != nil {
        fmt.Print(currentNode.data)
        fmt.Print(" -> ")
        currentNode = currentNode.next
    }
}

func (list *LinkedList) Remove(data interface{}) bool {
	if list.head == nil {
		return false
	}

	if list.head.data == data {
        list.head = list.head.next
        list.size--
        return true
    } 

	prevNode := list.head
    currentNode := list.head.next
    for currentNode != nil {
        if currentNode.data == data {
            prevNode.next = currentNode.next
            list.size--
            return true
        }
        prevNode = currentNode
        currentNode = currentNode.next
    }

    return false
}

func (list *LinkedList) Update(data interface{},updatedData interface{}) bool {
	if list.head == nil {
		return false
	}

	if list.head.data == data {
        list.head.data = updatedData
        return true
    } 

    currentNode := list.head.next
    for currentNode != nil {
        if currentNode.data == data {
            currentNode.data = updatedData
            return true
        }
        currentNode = currentNode.next
    }

    return false
}

func main() {
	myList := LinkedList{}
	myList.Add("apple")
	myList.Add("apple 1")
	myList.Add("apple 2")
	myList.Add("apple 3")
	myList.Traverse()
	// myList.Remove("apple")
	// myList.Remove("apple 3")
	myList.Update("apple 1","apple 4")
	fmt.Println(" ")
	myList.Traverse()
}