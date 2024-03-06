package firstTasks

import "fmt"

type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

func (bt *TreeNode) PreOrderTraversal() {
	if bt == nil {
		return
	}
	bt.Left.PreOrderTraversal()
	fmt.Printf("value: %v\n", bt.Value)
	bt.Right.PreOrderTraversal()
}
