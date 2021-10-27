## C++标准库

libstdc++与libc++都是c++标准库。 libstdc++使用gcc编译生成，libc++使用llvm编译生成。两者之间经常不能混合使用。

## 混用时常见错误

如编译静态库时使用libstdc++ abi，后面链接时使用libc++，就会出现 connot find symbol的错误

比如function<>产生的符号就不一样(可以使用c++filt进行查看)

## 指定使用的stdlib

add_definitions(-stdlib=libc++ -lc++abi)

如何libc++在custom路径可以指定 ```-I<path_to_libcxx>/include/c++/v1 -L<path_to_libcxx>/lib```

