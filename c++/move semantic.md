## rvalue

- 所有参数都是lvalue，只不过他的类型可以是 rvalue reference type



## universal reference

- 作为template parameter，格式为 T&&（加个const都不行）
- 和auto关键字一起使用
- template T&& param不一定就是universal引用，因为它有可能是作为模板类的一个成员函数，模板类实例化时就确定了T的类型，没有**type deduction**
- 使用universal reference使得代码更好维护、更好扩展也更有效率。使用左值和右值的重载来替代universal reference有时候会增加runtime的消耗
- avoid overloading on universal reference  ->可使用dispatch的技术

## move|forward semantic

- move|forward会modify the object，所以参数如果为const一定不会调用移动构造函数。同时正因为会修改object所以一般move|forward操作放在最后。
- rvalue与move配合，universal reference与forward配合



## 不要对local variable进行move操作

任何情况下都不用对局部变量使用move或者forward。编译器要么使用copy elision，要么默认当rvalue来使用

