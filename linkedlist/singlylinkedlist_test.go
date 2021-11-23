package linkedlist

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	ar := []int32{0, 1, 2, 3, 4}
	llist := SliceToSinglyLinkedList(ar)
	actual := SinglyLinkedListToSlice(llist)
	if !reflect.DeepEqual(ar, actual) {
		t.Errorf("expect %v but actual value is %v", ar, actual)
	}
}

func TestReverseSinglyLinkedList(t *testing.T) {
	ar := []int32{0, 1, 2, 3, 4}
	llist := SliceToSinglyLinkedList(ar)
	reversed := Reverse(llist)
	expected := []int32{4, 3, 2, 1, 0}
	actual := SinglyLinkedListToSlice(reversed)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expect %v but actual value is %v", ar, actual)
	}
}
