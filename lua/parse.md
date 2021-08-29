## parse过程

- 文件读取
- luaY_parser 生成token，根据token进行各种state的递归调用
- 将parse的结构Proto作为Closure的参数压入栈顶
- 从栈顶去除Closure进行相应调用

