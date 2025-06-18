func dfsTemplate(node *TreeNode) {
    if node == nil {
        return
    }

    // TODO: Process the current node
    // fmt.Println(node.Val)

    // Traverse left subtree
    dfsTemplate(node.Left)

    // Traverse right subtree
    dfsTemplate(node.Right)
}


DFS Template Variants:
// 1.Pre-Oreder : left -> current -> right
// 2.In-Order   : current -> left -> right
// 3.Post-Order : left -> right -> current 


// 1. Pre-order DFS 
func dfsPreOrder(node *TreeNode) {
	if node == nil {
		return
	}

	// Process the current node first
	fmt.Println(node.Val)

	// Traverse left subtree
	dfsPreOrder(node.Left)

	// Traverse right subtree
	dfsPreOrder(node.Right)
}


//2.Post-order DFS

func dfsPostorder(node *TreeNode) {
    if node == nil {
        return
    }
    dfsPostorder(node.Left)     // left
    dfsPostorder(node.Right)    // right
    fmt.Println(node.Val)       // current
}


// 3. In-order DFS
func dfsInOrder(node *TreeNode) {
	if node == nil {
		return
	}

	// Traverse left subtree
	dfsInOrder(node.Left)

	// Process the current node
	fmt.Println(node.Val)

	// Traverse right subtree
	dfsInOrder(node.Right)
}