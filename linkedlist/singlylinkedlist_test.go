package linkedlist

import (
	"reflect"
	"testing"
)

func TestSinglylist(t *testing.T) {
	ar := []int32{0, 1, 2, 3, 4}
	llist := SliceToSinglyLinkedList(ar)
	actual := SinglyLinkedListToSlice(llist)
	if !reflect.DeepEqual(ar, actual) {
		t.Errorf("expect %v but actual value is %v", ar, actual)
	}
}
