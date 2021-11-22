package singlylist

type SinglyLinkedListNode struct{
	data int32
	next *SinglyLinkedListNode
}

func (llist *SinglyLinkedListNode) Reverse() *SinglyLinkedListNode {
    var curr, tail *SinglyLinkedListNode
    curr = llist
    for curr != nil {
        tmp := curr.next
        curr.next = tail
        tail = curr
        curr = tmp
    }
    return tail
}
