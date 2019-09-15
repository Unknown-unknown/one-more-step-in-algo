# 数组

**数组**是一种线性表数据结构，用连续的内存空间存储相同类型的数据。

## 数组操作
### 随机访问

根据下标随机访问的时间复杂度是O(1)。

### 插入

**对于有序数据**

* 最好时间复杂度：数组末尾插入数据，无需移动数据，复杂度为O(1)；
* 最坏时间复杂度：数组开头插入数据，所有数据需依次移动一位，复杂度为O(n)；
* 平均时间复杂度：各个位置插入元素的概率相同，复杂度为(1+2+...n)/n = O(n)。

**对于无序数据**

直接将目标index上的数据移至数组末尾，将新元素放入目标index。这样时间复杂度直接降为O(1)。

### 删除

类似插入的时间复杂度。

为了避免每次删除都移动数据，可以先将已经删除的数据记录下来，当数组没有更多存储空间时再真正执行一次删除操作，以此减少删除操作导致的数据搬移。

**应用举例：**

* JVM标记清除垃圾回收算法

## 容器 vs 数组

很多语言提供了容器类，如Java的ArrayList、C++ STL中的vector，其优势在于：
* 封装数组操作的细节
* 支持动态扩容：但扩容涉及内存申请和数据搬移，耗时，所以如果事先能确定数据大小就先指定数据大小

数组的优势在于：
* 更关注性能时可以使用
* 数据简单操作时可以使用

如何选择：
* 更注重性能优化的底层开发，首选数组
* 更注重开发效率的业务开发，首选容器，省时省力

---

## 动手小例子

测试数组访问越界之后的运行结果，如`out_of_range.c`：
```lang=c
$ gcc -o out_of_range out_of_range.c
$ ./out_of_range
> 
hello world
hello world
hello world
hello world
[1]    23919 abort      ./out_of_range
```

并没有如原文中讲到的会无限循环。看起来需要查找 **函数调用的栈桢结构** 的细节。

## 对应LeetCode题目
[x]15.Three Sum（求三数之和）：https://leetcode.com/problems/3sum/

[x]169.Majority Element（求众数）：https://leetcode.com/problems/majority-element/

[]41.Missing Positive（求缺失的第一个正数）：https://leetcode.com/problems/first-missing-positive/


