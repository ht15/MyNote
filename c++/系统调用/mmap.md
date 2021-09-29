当mmap不将地址映射到某个文件时，相当于进行了malloc操作：

```c
void* my_malloc(size_t len) {
	void* ret = mmap(0, len, PROT_READ | PROT_WRITE, MAP_PRIVATE | MAP_ANONYMOUS, 0, 0);
	if (ret == MAP_FAILED)
		return 0;
	return ret;
}
```



值得注意的是**起始地址和大小都必须是系统页的整数倍**



参考：[程序员的自我修炼 10.3.2]

