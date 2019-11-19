# 递归

## 思维要点

1. 不要人肉进行递归（**最大误区**）
2. 找到最近最简方法，将其拆解成可重复解决的问题（重复子问题）
3. 数学归纳法思维

## 递归模板

1. 找到终止条件
2. 处理当前层的逻辑
3. 递归到下一层中
4. 如有必要恢复当前层被改变的状态


### Python 代码模板

```lang=python
def recursion(level, param1, param2, ...): 
    # recursion terminator 
    if level > MAX_LEVEL: 
	   process_result 
	   return 

    # process logic in current level 
    process(level, data...) 

    # drill down 
    self.recursion(level + 1, p1, ...) 

    # reverse the current level status if needed
```

### Java 代码模板
```lang=java
public void recur(int level, int param) { 
  // terminator 
  if (level > MAX_LEVEL) { 
    // process result 
    return; 
  } 

  // process current logic 
  process(level, param); 

  // drill down 
  recur( level: level + 1, newParam); 

  // restore current status 
 
}
```