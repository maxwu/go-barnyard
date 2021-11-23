package linkedlist

import (
	"log"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

// Reverse func reverses the singly linked list and return the new header
func Reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	var curr, prev *SinglyLinkedListNode
	curr = llist
	for curr != nil {
		tmp := curr.next
		curr.next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func NewNode() *SinglyLinkedListNode {
	return &SinglyLinkedListNode{}
}

func SliceToSinglyLinkedList(ar []int32) *SinglyLinkedListNode {
	var hdr, p *SinglyLinkedListNode
	for _, a := range ar {
		log.Printf("Insert data %d\n", a)
		if p == nil {
			p = NewNode()
			hdr = p
			p.data = a
		} else {
			p.next = NewNode()
			p = p.next
			p.data = a
		}
	}
	return hdr
}

func SinglyLinkedListToSlice(hdr *SinglyLinkedListNode) []int32 {
	var ar []int32 = make([]int32, 0)
	p := hdr
	for p != nil {
		ar = append(ar, p.data)
		p = p.next
	}
	return ar
}
