## elf文件内容查看

- readelf
- objdump

## 解析C++被修饰过的符号

c++filt symbol_name

## 查看导出的符号

```dumpbin /EXPORTS XX.dll```

## 获取共享库的依赖情况

ldd  /bin/ls

## 将进程空间的内容导出

比如导出内容的到linux-gate.dsos

``` dd if=/proc/self/mem of=linux-gate.dso bs=4096 skip=1048574 count=1```



[参考](https://blog.csdn.net/zqixiao_09/article/details/50783007)

