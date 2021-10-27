## 平台

tlinux2 Centos7

## 源码编译安装clang+llvm

- 获取源码

    从[LLVM Download Page](https://releases.llvm.org/download.html#git)下载指定版本的 LLVM source code、Clang source code 、 clang-tools-extra、 compiler-rt source code

- 源码组织

    clang作为llvm的一个编译前端，需要将clang源码放置到llvm的子目录中，如如下伪代码所示：

    ```shell
    mv clang.src clang
    mv clang llvm.src/tools/
    
    mv clang-tools-extra extra
    mv extra llvm.src/tools/clang/
    
    mv compiler-rt.src compiler-rt
    mv compiler-rt llvm.src/projects/
    ```

- 编译

    ```shell
    cd llvm.src
    mkdir build && cd build
    
    cmake3 -DCMAKE_INSTALL_PREFIX=/usr/local/clang-xx -DCMAKE_BUILD_TYPE=Release -DLLVM_ENABLE_ASSERTIONS=On .. && make -j8 && make install
    ```

引用：

https://www.cnblogs.com/BinBinStory/p/7499527.html

## 编译libc++.a、libc++abi

- 获取源码

    下载指定版本的 libc++ source code 、libc++abi source code

- 构建libc++

    ```shell
    cd libc++
    mkdir build && cd build
    
    cmake3 -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr -DCMAKE_C_COMPILER=clang -DCMAKE_CXX_COMPILER=clang++ .. && make install
    ```

- 构建libc++abi

    ```shell
    cd libc++abi
    mkdir build && cd build
    
    cmake3 -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr -DCMAKE_C_COMPILER=clang -DCMAKE_CXX_COMPILER=clang++ -DLIBCXXABI_LIBCXX_INCLUDES=../../libcxx/include .. && make install
    ```

- 使用libc++abi构建libc++

    ```shell
    cd libc++
    mkdir build && cd build
    
    cmake3 -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr -DCMAKE_C_COMPILER=clang -DCMAKE_CXX_COMPILER=clang++ -DLIBCXX_CXX_ABI=libcxxabi -DLIBCXX_CXX_ABI_INCLUDE_PATHS=../../libcxxabi/include .. && make install
    ```

    

引用：

https://www.cnblogs.com/BinBinStory/p/7499648.html