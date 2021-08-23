- Upvalue 为Closure中的一个属性

- 如何调用c函数的，主要是如何传入参数的

  可以调用lua_pushcfunction注册一个c函数，关键是如何传参的

- lua调用完c function后，怎么自动清除results(result放在栈顶的位置)之外的栈的？