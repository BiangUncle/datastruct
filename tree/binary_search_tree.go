package tree

type BSTree interface {
	GetRoot() *TreeNode
	Less(val interface{}) bool
}

func Insert(t BSTree, Val interface{}) {
	t.GetRoot().Less(Val)
}
