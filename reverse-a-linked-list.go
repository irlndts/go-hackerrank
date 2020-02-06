package hackerrank

// https://www.hackerrank.com/challenges/reverse-a-linked-list/problem

// SinglyLinkedListNode ...
type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func reverse(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	var node, next *SinglyLinkedListNode
	for {
		next = head.next

		if node != nil {
			head.next = node
		} else {
			head.next = nil
		}

		if next == nil {
			break
		}

		node = head
		head = next
	}
	return head
}
