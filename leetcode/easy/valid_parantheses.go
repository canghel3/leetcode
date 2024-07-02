package easy

func IsValid(s string) bool {
	//trivial case
	if len(s)%2 == 1 {
		return false
	}

	var link = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	//trivial case
	//if the first element is any type of closing bracket, return false
	if _, ok := link[rune(s[0])]; !ok {
		return false
	}

	var stack []rune
	for _, r := range s {
		_, isOpener := link[r]
		if isOpener {
			stack = append(stack, r)
		} else {
			//if stack is empty and current element is a closer
			if len(stack) == 0 {
				return false
			}

			//if the current closing bracket does not correspond to the last opening bracket
			if link[stack[len(stack)-1]] != r {
				return false
			} else {
				//otherwise just remove both (pop)
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{}

	modifyHead := true
	for list1.Next == nil || list2.Next == nil {
		if list1.Next != nil && list2.Next != nil {
			if list1.Val < list2.Val {
				if modifyHead {
					merged.Val = list1.Val
					merged.Next = nil
				} else {
					current := merged.Next
					for current != nil {
						current = merged.Next
					}
				}

			}
		}
	}

	return merged
}
