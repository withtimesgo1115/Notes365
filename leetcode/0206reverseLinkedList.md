给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 

示例 1：
![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

## 思路
双指针方法
1. 定义两个双指针，一个在前pre,一个在后cur
2. 每次让pre的next指向cur，实现一次局部反转
3. 局部做完之后，pre和cur同时往前移动一个位置
4. 循环这个过程，直到pre到达链表尾部

```py
class Solution:
    def reverseList(self, head):
        pre = None
        cur = head
        while cur:
            next_node = cur.next
            cur.next = pre
            pre = cur
            cur = next_node
        return pre
```