## 为什么函数可以返回unique_ptr

- 编译器具有返回值优化，直接在接收方的对unique_ptr进行构造
- unique_ptr支持move语义

