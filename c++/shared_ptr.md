## get()获取的指针的生命周期必须小于 shared_ptr析构

```c++
Resource* p = new CResource;  
{  
    shared_ptr q(p);  
}  
p->Use() // CRASH   
```

[参考](https://www.jianshu.com/p/f1925247c14f?from=timeline&isappinstalled=0)

同时get()不会增加shared_ptr的use_count()

## control block的创建时机

- 使用 make_shared
- 通过 unique_ptr或者auto_ptr进行构造
- 通过raw pointer进行构造

一个原始指针不能对应两个control block，要不然第二次析构时会产生**undefined behavior**。使用建议是 直接传递 new o给构造函数

### CRTP(curiously recurring template pattern)

多个control block的问题对this指针的引用也同样存在，为此出现了**enable_shared_from_this**。

enable_shared_from_this的原理很简单：如果需要返回指向this的shared_ptr时，调用shared_from_this即可，在调用shared_from_this之前必须要存在指向this的shared_ptr，所以一般derived from enable_shared_from_this的类的构造函数都写成私有的，通过public static method来返回shared_ptr。同时在shared_ptr的构造函数中有一个友方法，当raw pointer的类型是enable_shared_from_this时，会产生一个weak_ptr。shared_from_this就是weak_ptr来构造新的shared_ptr的，这样一来this就对应了唯一的一个control_block



## weak_ptr可以解决循环引用问题

