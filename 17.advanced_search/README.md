# 高级搜索

## 初级搜索

1. 朴素搜索
2. 优化方式：不重复（Fibonacci）、剪枝（生成括号问题）
3. 搜索方向：BFS、DFS

## 高级搜索

1. 双向搜索
2. 启发式搜索/A*算法/优先级搜索：优先队列

## TODO

- LC.22 dp+剪枝的题解

## A*搜索

```lang=python
def AstarSearch(graph, start, end):
    pq = collections.priority_queue() # 优先级 -> 估价函数
    pq.append([start])
    visited.add(start)

    while pq:
        node = pq.pop() # can we add more intelligence here?
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        unvisited = [node for node in nodes if node not in visited]

        pq.push(unvisited)
```