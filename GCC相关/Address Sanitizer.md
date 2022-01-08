## CMake 启用

```cmake
set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} -fno-omit-frame-pointer -fsanitize=address")
set(CMAKE_CXX_FLAGS_DEBUG "${CMAKE_CXX_FLAGS_DEBUG} -fno-omit-frame-pointer -fsanitize=address")
set(CMAKE_LINKER_FLAGS_DEBUG "${CMAKE_LINKER_FLAGS_DEBUG} -fno-omit-frame-pointer -fsanitize=address")
```

https://stackoverflow.com/questions/44320465/whats-the-proper-way-to-enable-addresssanitizer-in-cmake-that-works-in-xcode



## 使用记录

### 非-O0选项double free问题竟然会被修复！！！，测试时还是使用-O0





**选项**

-fsanitize=address : 打开asan内存错误检查

-fno-omit-frame-pointer : 保留函数调用的帧信息，以便分析函数调用关系

-fsanitize=undefined : 未定义行为检测

-fsanitize-recover=address : 监测到内存错误后继续执行,输出当前编译的代码文件中的所有地址错误（gcc 6.1/clang 5.0以上支持）。使用时关闭halt_on_error

ASAN_OPTIONS=halt_on_error=0 ./test

**算法**

https://github.com/google/sanitizers/wiki/AddressSanitizerAlgorithm

**gdb调试**

b __asan_report_error





asan用了再上ubsan, fuzzing interface