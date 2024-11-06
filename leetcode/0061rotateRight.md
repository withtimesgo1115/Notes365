给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置

![](https://assets.leetcode.com/uploads/2020/11/13/rotate1.jpg)

```py
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        # 特殊处理
        if not head or not head.next or k == 0:
            return head
        # 记录长度，形成环
        n = 1
        tail = head
        while tail.next:
            tail = tail.next
            n += 1
        tail.next = head
        # 找到要断开的位置
        new_tail = head
        for i in range(n - k % n - 1):
            new_tail = new_tail.next
        
        new_head = new_tail.next # 新的头结点
        new_tail.next = None  # 断开环
        return new_head 
```