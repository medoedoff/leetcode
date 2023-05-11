package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	arr := []*ListNode{head}
	for arr[len(arr)-1].Next != nil {
		arr = append(arr, arr[len(arr)-1].Next)
	}
	return arr[len(arr)/2]
}
