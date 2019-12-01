# 位运算

## 异或（XOR）

异或：相同为 0，不同为 1.也可用“不进位加法”来理解。

异或操作的一些特点：

```
x ^ 0 = x
x ^ 1s = ~x // 注意 1s = ~0
x ^ (~x) = 1s
x ^ x = 0
c = a ^ b => a ^ c = b, b ^ c = a   // 交换两个数
a ^ b ^ c = a ^ (b ^ c) = (a ^ b) ^ c   // associative
```

## 指定位置的位运算

1. 将 x 最右边的 n 位清零：`x&(~0<<n)`
2. 获取 x 的第 n 位值（0或者1）：`(x>>n)&1`
3. 获取 x 的第 n 位的幂值：`x&(1<<(n-1))`
4. 仅将第 n 位置为1：`x|(1<<n)`
5. 仅将第 n 位置为0：`x&(~(1<<n))`
6. 将 x 最高位至第 n 位（含）清零：`x&((1<<n)-1)`
7. 将第 n 位至第0位（含）清零：`x&(~((1<<(n+1))-1))`

## 实战位运算要点

- 判断奇偶：
```
x%2==1 -> (x&1)==1
x%2==0 -> (x&1)==0
```

- x>>1 -> x/2，即：`x=x/2 -> x=x>>1`，`mid=(left+right)/2 -> mid=(left+right)>>1`
- `x=x&(x-1)`，清零最低位的1
- `x&-x`，得到最低位的1
- `x&~x`，0

## N 皇后位运算代码示例

```lang=python
def totalNQueens(self, n):
    if n < 1: return []
    self.count = 0
    self.DFS(n, 0, 0, 0, 0)
    return self.count

def DFS(self, n, row, cols, pie, na):
    if row >= n:
        self.count += 1
        return

    bits = (~(cols | pie | na)) & ((1<<n) - 1)  # 得到当前所有的空位

    while bits:
        p = bits & -bits # 取到最低位的1
        bits = bits & (bits - 1) # 表示在p位置上放入皇后
        self.DFS(n, row+1, cols | p, (pie | p) << 1, (na | p) >> 1)
        # 不需要revert cols，pie，na的状态
```

## Golang的一点差别

- Golang 中不支持 `~` 这个运算符。
- ^(XOR) ，如果是作为一元运算符出现，意为按位取反；如果是作为二维运算符出现，则为异或运算。