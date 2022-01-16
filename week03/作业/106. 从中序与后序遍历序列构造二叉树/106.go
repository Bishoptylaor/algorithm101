package main

/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2022/1/15 -- 8:42 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 106.go
*/

// TreeNode /**
type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    idxMap := map[int]int{}
	// 因为题目假定了没有重复元素
    for i, v := range inorder {
        idxMap[v] = i
    }
    var build func(int, int) *TreeNode
    build = func(inorderLeft, inorderRight int) *TreeNode {
        // 无剩余节点
        if inorderLeft > inorderRight {
            return nil
        }

        // 后序遍历的末尾元素即为当前子树的根节点
        val := postorder[len(postorder)-1]
		// 使用过后减掉根节点
        postorder = postorder[:len(postorder)-1]
        root := &TreeNode{Val: val}

        // 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
        // 根据后序性质，距离根节点最近的是右子树，所以优先构造右子树
        inorderRootIndex := idxMap[val]
        root.Right = build(inorderRootIndex+1, inorderRight)
        root.Left = build(inorderLeft, inorderRootIndex-1)
        return root
    }
    return build(0, len(inorder)-1)
}

