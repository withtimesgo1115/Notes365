给定一个二叉树 root ，返回其最大深度。

二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

![](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)

输入：root = [3,9,20,null,null,15,7]
输出：3

## 思路
找出终止条件：当前节点为空
找出返回值：节点为空时说明高度为 0，所以返回 0，节点不为空时则分别求左右子树的高度的最大值，同时加 1 表示当前节点的高度，返回该数值
某层的执行过程：在返回值部分基本已经描述清楚
时间复杂度：O(n)

```py
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        else:
            left = self.maxDepth(root.left)
            right = self.maxDepth(root.right)
            return max(left, right) + 1
```