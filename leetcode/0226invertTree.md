给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

![](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)


## 思路
其实就是交换一下左右节点，然后再递归的交换左节点，右节点
根据动画图我们可以总结出递归的两个条件如下：

终止条件：当前节点为 null 时返回
交换当前节点的左右节点，再递归的交换当前节点的左节点，递归的交换当前节点的右节点
时间复杂度：每个元素都必须访问一次，所以是 O(n)
空间复杂度：最坏的情况下，需要存放 O(h) 个函数调用(h是树的高度)，所以是 O(h)
代码实现如下：


![](https://pic.leetcode-cn.com/0f91f7cbf5740de86e881eb7427c6c3993f4eca3624ca275d71e21c5e3e2c550-226_2.gif)

```py
class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root:
            return
        root.left, root.right = root.right, root.left
        self.invertTree(root.left)
        self.invertTree(root.right)
        return root
```