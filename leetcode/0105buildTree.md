给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

![](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]


## 思路
![](https://pic.leetcode-cn.com/beff309937462b352940c1925de8ff50c22b65bada872cf286b0228a45054ea2-2.jpg)

递归函数实现如下：

终止条件:前序和中序数组为空
根据前序数组第一个元素，拼出根节点，再将前序数组和中序数组分成两半，递归的处理前序数组左边和中序数组左边，递归的处理前序数组右边和中序数组右边。

```py
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if not preorder and not inorder:
            return
        root = TreeNode(preorder[0])
        mid_idx = inorder.index(preorder[0])

        root.left = self.buildTree(preorder[1:mid_idx+1], inorder[:mid_idx])
        root.right = self.buildTree(preorder[mid_idx+1:], inorder[mid_idx+1:])
        return root
```