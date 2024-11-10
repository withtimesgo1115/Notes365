给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

![](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)

## 思路
标签：深度优先遍历 DFS

终止条件与返回值：

1. 当两棵树的当前节点都为 null 时返回 true

2. 当其中一个为 null 另一个不为 null 时返回 false

3. 当两个都不为空但是值不相等时，返回 false

执行过程：当满足终止条件时进行返回，不满足时分别判断左子树和右子树是否相同，其中要注意代码中的短路效应

时间复杂度：O(n)，n 为树的节点个数


```py
class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if not p and not q:
            return True
        if not p or not q:
            return False
        if p.val != q.val:
            return False
        return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
```