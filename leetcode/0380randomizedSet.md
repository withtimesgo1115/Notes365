实现RandomizedSet 类：

RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。


## 思路
用一个数组和一个hashmap来实现，hashmap记录val => array index

插入时，如果存在则返回false, 否则在数组末尾添加元素，记录索引

删除时可以先检查val是否存在，不存在直接返回false
否则，先获取索引id，然后更新索引id所在位置的元素值为数组尾部元素
然后更新索引记录，把尾部元素的索引改为id
然后删除尾部元素以及字典中的key


随机元素可以用choice()方法

```py
class RandomizedSet:
    def __init__(self):
        self.nums = []
        self.indices = {}
    
    def insert(self, val):
        if val in self.indices:
            return False
        self.indices[val] = len(self.nums)
        self.nums.append(val)
        return True
    
    def remove(self, val):
        if val not in self.indices:
            return False
        id = self.indices[val]
        self.nums[id] = self.nums[-1]
        self.indices[self.nums[id]] = id
        self.nums.pop()
        del self.indices[val]
        return True

    def getRandom(self):
        return choice(self.nums)
```
