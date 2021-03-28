## vscode配置C++调试的说明

### 编译

自己编写MAKEFILE，然后执行make && make install，或者编写CMakeLists.txt然后cmake。[参考这里](https://oldpan.me/archives/gcc-make-cmake-clang-tell)

编译只是第一步，我们怎么能达到运行调试前一定会编译好源文件呢？这里可以通过task来实现。将这个编译的task设置为launch的preLaunchTask

```json
{
  	"version": "2.0.0",
    "tasks": [
        {
            "label": "handy_make",
            "type": "shell",
            "command": "make && make install"
        }
    ]
}
```

### 运行+调试

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "g++",
            "type": "cppdbg",
            "request": "launch",
            "program": "${fileDirname}/../10m/10m-svr", // 设置运行程序
            "args": [
                "31100",  // 运行参数
                "31110",
                "1",
                "31102"
            ],
            "stopAtEntry": false,
            "cwd": "${workspaceFolder}",
            "environment": [],
            "externalConsole": false,
            "MIMode": "gdb", // lldb 调试工具
            "preLaunchTask": "handy_make", // 运行前task，编译等工作
            "setupCommands": [
                {"text": "-gdb-set follow-fork-mode child"}, 
                //{"text": "-gdb-set attach-on-fork on"},
            ],
        }
    ],
}
```

