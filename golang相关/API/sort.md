- sort.Slice

自定义排序方法，比如对 [ [7, 0], [7, 1], [6, 1], [5, 0], [5, 2], [4, 4] ]先按第一位降序排序，再按第二位升序排序，可以写成：

```go
sort.Slice(seq, func(i, j int) bool) {
  if seq[i][0] == seq[j][0] {
    return seq[i][1] < seq[j][1]
  } else {
    return seq[i][0] > seq[j][0]
  }
}
```



