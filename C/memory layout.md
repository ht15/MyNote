以struct为例

- 内部属性按申明顺序在内存中排列，且order不会改变
- 用户可以改变align的字节数，比如使用 #pragma pack(push,1)

所以struct序列化非常简单：

```c
#paragma pack(push, 1)
struct A {
	int num;
};
#paragma pack(pop)

// 发送
static char send_buff[1024];
A* ptr_a = reinterpret_cast<A*>(send_buff);
ptr_a->num = 2;
send(send_buff, sizeof(A))
    
// 接收
A* ptr_a = reinterpret_cast<A*>(recv_buff);
```

