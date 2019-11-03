# DFS & BFS

## DFS 代码模板

1. 递归写法

```
visted = set()

def dfs(node, visited):
    if node in visited: // terminator
        return

    visited.add(node)

    # process current node here.
    ...
    for next_node in node.children():
        if not next_node in visited:
            dfs(next_node, visited)            
```

2. 非递归写法

```
def DFS(self, tree):
    if tree.root is None:
        return []

    visited, stack = [], [tree.root]

    while stack:
        node = stack.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        stack.push(nodes)

    # other processing work
    ...
```

## BFS 代码模板

类似树的按层遍历，使用一个 queue 来维护节点：

- 先将根节点放入 queue 中；
- 循环该 queue，pop 出一个节点，将其标记为已访问，
- 处理该节点，并将该节点的所有子节点 push 到 queue 中
- 直到 queue 为空，退出循环，遍历结束

由于 queue 是先进先出的顺序，所以节点的访问顺序也是按层来访问的，即广度优先遍历。

```
def BFS(graph, start, end):
    queue = []
    queue.append([start])
    visited.add(start)

    while queue:
        node = queue.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        queue.push(nodes)

    # other processing work
    ...
```