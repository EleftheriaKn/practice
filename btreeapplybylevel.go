package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	theQueue := []*TreeNode{root}
	for len(theQueue) > 0 {
		theNode := theQueue[0]
		theQueue = theQueue[1:]
		_, err := f(theNode.Data)
		if err != nil {
			return
		}
		if theNode.Left != nil {
			theQueue = append(theQueue, theNode.Left)
		}
		if theNode.Right != nil {
			theQueue = append(theQueue, theNode.Right)
		}
	}
}
