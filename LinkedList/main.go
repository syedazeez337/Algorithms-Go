package main

import "fmt"

// LinkedListNode represents a node in the linked list
type LinkedListNode struct {
    Value interface{}
    Next  *LinkedListNode
}

// LinkedListInsert inserts a new node at the given index in the linked list
func LinkedListInsert(head *LinkedListNode, index int, value interface{}) (*LinkedListNode, error) {
    // Special case: inserting a new head node
    if index == 0 {
        newHead := &LinkedListNode{Value: value, Next: head}
        return newHead, nil
    }

    current := head
    var previous *LinkedListNode
    count := 0

    // Traverse the list to find the position
    for count < index && current != nil {
        previous = current
        current = current.Next
        count++
    }

    // Check if the index is invalid
    if count < index {
        return nil, fmt.Errorf("invalid index: %d", index)
    }

    // Insert the new node
    newNode := &LinkedListNode{Value: value}
    if previous != nil {
        newNode.Next = previous.Next
        previous.Next = newNode
    }

    return head, nil
}

// Helper function to print the linked list
func PrintLinkedList(head *LinkedListNode) {
    current := head
    for current != nil {
        fmt.Printf("%v -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}

// Example usage
func main() {
    // Create a linked list: 1 -> 2 -> 3 -> nil
    head := &LinkedListNode{Value: 1}
    head.Next = &LinkedListNode{Value: 2}
    head.Next.Next = &LinkedListNode{Value: 3}

    fmt.Println("Original list:")
    PrintLinkedList(head)

    // Insert at index 1
    head, err := LinkedListInsert(head, 1, 10)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("After inserting 10 at index 1:")
    PrintLinkedList(head)

    // Insert at index 0
    head, err = LinkedListInsert(head, 0, 0)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("After inserting 0 at index 0:")
    PrintLinkedList(head)

    // Attempt to insert at an invalid index
    _, err = LinkedListInsert(head, 10, 20)
    if err != nil {
        fmt.Println(err)
    }
}
