### incomplete type

- The void type comprises an empty set of values; it is an incomplete type that cannot be completed.
- An array type of unknown size is an incomplete type.
- A structure or union type of unknown content (as described in 6.7.2.3) is an incomplete type.

char[0]就是一个incomplete type

### c99之后运行struct包含一个incomplete type

所以c风格的代码是可以使用长度为0的数组来处理变长情况的，c++则不能。

一般用法：

```c
typedef struct A{
  char tag;
  char alz[0];
}Astruct;

int main(void){

  Astruct *x=(struct A*)malloc(sizeof(Astruct)+sizeof(char)*10);
  free(x);

  return 0;
}
```

注意 ```sizeof(A)```的结果为1

### 参考

[Array of length zero](https://imzlp.com/posts/21095/)