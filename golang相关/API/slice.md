# 插入

- 数据 array = [[1, 2], [5, 6]]，在中间插入val = [3, 4]

  ```go
  array = append(array[0:1], append([][]int{val}, array[1:]...)...)
  ```

  