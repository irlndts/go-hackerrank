package linkedlist

// https://www.hackerrank.com/challenges/delete-a-node-from-a-linked-list/problem

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func deleteNode(head *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	if position == 0 {
		return head.next
	}

	start := head

	i := int32(0)
	for head.next != nil {
		if i+1 == position {
			head.next = head.next.next
			break
		}

		i++
		head = head.next
	}

	return start
}
