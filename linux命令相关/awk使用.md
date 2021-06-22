## 统计两个字段间包含的行数

首先定义匹配模板

```shell
OUTPUTP="<OUTPUT>"
EOUTPUTP="</OUTPUT>"
```



```shell
gawk -v eoutput="$EOUTPUTP" -v output="$OUTPUTP" -F "$" 'BEGIN{ b=0; e=0; c=0}{ if($0 ~ output) { b=NR;} else if($0 ~ eoutput){e=NR; if(e-b>2){c++}} } END{print c;print b; print e}' 77.xml
```

