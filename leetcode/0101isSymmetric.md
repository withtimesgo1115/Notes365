给你一个二叉树的根节点 root ， 检查它是否轴对称。

![](https://pic.leetcode.cn/1698026966-JDYPDU-image.png)

输入：root = [1,2,2,3,4,4,3]
输出：true

## 思路
乍一看无从下手，但用递归其实很好解决。
根据题目的描述，镜像对称，就是左右两边相等，也就是左子树和右子树是相当的。
注意这句话，左子树和右子相等，也就是说要递归的比较左子树和右子树。
我们将根节点的左子树记做 left，右子树记做 right。比较 left 是否等于 right，不等的话直接返回就可以了。
如果相当，比较 left 的左节点和 right 的右节点，再比较 left 的右节点和 right 的左节点
比如看下面这两个子树(他们分别是根节点的左子树和右子树)，能观察到这么一个规律：
左子树 2 的左孩子 == 右子树 2 的右孩子
左子树 2 的右孩子 == 右子树 2 的左孩子

```
    2         2
   / \       / \
  3   4     4   3
 / \ / \   / \ / \
8  7 6  5 5  6 7  8
```
根据上面信息可以总结出递归函数的两个条件：
终止条件：

- left 和 right 不等，或者 left 和 right 都为空
- 递归的比较 left，left 和 right.right，递归比较 left，right 和 right.left

```py
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True
        def dfs(left, right):
            if not left and not right:
                return True
            if not left or not right:
                return False
            if left.val != right.val:
                return False
            return dfs(left.left, right.right) and dfs(left.right, right.left)
```