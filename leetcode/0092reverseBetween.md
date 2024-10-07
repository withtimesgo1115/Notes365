给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 

示例 1：
![](https://assets.leetcode.com/uploads/2021/02/19/rev2ex2.jpg)

输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

## 思路
1. 用一个dummyNode
2. 找到要反转的节点的前一个位置p0
3. 先开始局部反转
4. 然后注意p0.next.next也就是原来的起始元素，现在应该连接到cur上
5. p0.next应该链接到原来的尾部元素，也就是pre
6. 返回dummy.next即可

```py
class Solution:
    def reverseBetween(self, left, right):
        p0 = dummy = ListNode(next=head)
        for _ in range(left - 1):
            p0 = p0.next
        
        pre = None
        cur = p0.next
        for _ in range(right - left + 1):
            next_node = cur.next
            cur.next = pre
            pre = cur
            cur = next_node
        
        p0.next.next = cur
        p0.next = pre
        return dummy.next
```