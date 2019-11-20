# 并查集

## 常见场景

组团、配对问题

## 代码模板

### Java 代码模板

```lang=java
class UnionFind {
    private int count = 0;
    private int[] parent;
    public UnionFind(int n) {
        count = n;
        parent = new int[n];
        for (int i = 0; i < n; i++) {
            parent[i] = i;  // 初始化，根节点的特点
        }
    }
    public int find(int p) {
        while (p != parent[p]) {
            parent[p] = parent[parent[p]];  // 注意这两行的顺序
            p = parent[p];
        }
        return p;
    }
    public void union(int p, int q) {
        int rootP = find(p);
        int rootQ = find(q);
        if (rootP == rootQ) return;
        parent[rootP] = rootQ;
        count--;
    }
}
```

### Python 代码模板

```lang=python
def init(p):
    # for i = 0 .. n: p[i] = i;
    p = [i for i in range(n)]

def union(self, p, i, j):
    p1 = self.parent(p, i)
    p2 = self.parent(p, j)
    p[p1] = p2

def parent(self, p, i):
    root = i
    while p[root] != root:
        root = p[root]
    while p[i] != i: # 路径压缩（optional）：如果 a = parent[b], b = parent[c], c = parent[d], 则可以压缩成 a = parent[b, c, d]
        x = i; i = p[i]; p[x] = root
    return root
```