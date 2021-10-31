- -Dmacro

    定义一个宏。如-DISMACRO定义了一个ISMACRO宏，代码中可以：

    ```c++
    #ifdef ISMACRO
    ...
    #endif
    ```

- -I，-L, -WI,-rpath

    指定头文件搜索路径 -I

    指定库搜索路径 -L. 为啥要配合-Wl,-rpath?(我想rpath指定的是运行时依赖库的查找路径，而-L则是链接时)；另一个方法时[修改/etc/ld.so.conf](https://blog.csdn.net/mybelief321/article/details/9099659)

    [Cmake设置rpath的方法](https://love.junzimu.com/archives/2758)：set_target_properties(myexe PROPERTIES INSTALL_RPATH "$ORIGIN;$ORIGIN/../lib")

    其中[$ORIGIN](https://blog.csdn.net/chengyq116/article/details/104552007)为可执行文件的路径

