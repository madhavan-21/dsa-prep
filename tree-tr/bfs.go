func bfsTemplate(root *TreeNode) {
    if root == nil {
        return
    }

    queue := []*TreeNode{root}
    level := 0

    for len(queue) > 0 {
        levelSize := len(queue)

        // Optional: variables for level-specific processing
        // sum := 0

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]

            // TODO: Process the node
            // sum += node.Val
            // fmt.Println(node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        // Optional: logic after each level
        // fmt.Printf("Level %d sum = %d\n", level, sum)

        level++
    }
}
