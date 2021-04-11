### get()获取的指针的生命周期必须小于 shared_ptr析构

```c++
Resource* p = new CResource;  
{  
    shared_ptr q(p);  
}  
p->Use() // CRASH   
```

[参考](https://www.jianshu.com/p/f1925247c14f?from=timeline&isappinstalled=0)

同时get()不会增加shared_ptr的use_count()

