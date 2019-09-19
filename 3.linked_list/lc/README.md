# 链表相关常见技巧

## 1.理解指针或者引用

## 2.警惕指针丢失和内存泄漏

在写 [19.remove nth node from end of list](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) 时，我犯过一个错误，错误代码如下：

```
1 func removeNthFromEnd(head *ListNode, n int) *ListNode {
2     slow, fast, i := head, head, 1
3	for i <= n {
4		fast = fast.Next
5		i++
6	}
7	for fast.Next != nil {
8		slow = slow.Next
9		fast = fast.Next
10	}
11	slow.Next = slow.Next.Next
12	return head
13}
```

乍一看好像没什么问题，简单测试也可以通过。但是，当linked list 为 1，n = 1 时，第 7 行的 `fast.Next` 就会成为野指针，导致程序直接 panic。而这个问题，可以通过下面“增加哨兵”的技巧来规避。

其他可能出现指针丢失和内存泄漏的场景还包括：

- 插入、删除结点时的操作顺序
- 首尾结点的特殊处理
- 等等……

## 3.利用哨兵简化实现难度

（据说）利用哨兵简化编程难度的技巧，在很多代码实现中都有用到，比如：

- 插入排序
- 归并排序
- 动态规划
- 等等……

等遇到了再回来补充吧~~ 🤓

repo 内搜索 `![sentry]` 查看所有相关题目。

## 4.重点留意边界条件的处理

备查列表：

- 链表为空时
- 链表只包含一个结点时
- 链表只包含两个结点时
- 处理头结点和尾结点时
- 其他特定边界……

## 5.一图胜千言

祝早日成为灵魂画手🍻

## 6.One more thing: Practice makes perfect

没什么好说的。试试“五毒神掌”？😎