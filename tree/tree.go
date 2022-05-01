package tree

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Less(val interface{}) {

}
