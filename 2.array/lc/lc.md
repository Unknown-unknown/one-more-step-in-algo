# Array相关常见技巧

## 1.双指针类型：追及问题和相遇问题

双指针技巧通常有两种类型，一种快慢指针，一种首尾指针，像极了小学数学应用题中的追及问题和相遇问题。🤔
repo 内搜索 `![two-pointer]` 查看所有相关题目。

- 快慢指针：
    - 一个慢指针，一个快指针，通常前者计数，后者遍历。
    - 一个典型的例子，是从有序数组中移除重复的元素，比如 [26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)。
- 首尾指针：
    - 一个指针从头开始，一个指针从尾开始，根据条件向中间回合，相遇为止。

- 典型习题：(3/57, 20190917)

    - [x][[26] Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)
    - [x][[167] Two Sum II - Input array is sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)
    - [[557] Reverse Words in a String II](https://leetcode.com/problems/reverse-words-in-a-string-ii/description/)(*Locked*)
    - [[189] Rotate Array](https://leetcode.com/problems/rotate-array/)
    - [[125] Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)
    - [[11] Container With Most Water](https://leetcode.com/problems/container-with-most-water/)
    - [[238] Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)
    - [[15] 3Sum](https://leetcode.com/problems/3sum/)

- 参考：

    - https://leetcode.com/articles/two-pointer-technique/
    - https://www.geeksforgeeks.org/two-pointers-technique/
    - https://github.com/liyin2015/Algorithms-and-Coding-Interviews/blob/master/two_pointer.pdf

## 2.哈希表

类似的功能可以用简单的数组实现，或者直接用哈希表（不同语言的类型不同，如 Java 中的 `HashMap`，C++ 中的 `unordered_map`，Golang 中的 `map` 等等）。

- 简单的数组可以当成计数器来使用，每个元素用来记录跟踪字符出现的频率。
- 哈希表的 key 记录字符，value 记录出现频率。

