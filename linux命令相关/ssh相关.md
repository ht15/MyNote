## 生成公私钥

```shell
ssh-keygen -t rsa -C "name@mail"
```

## 免密登录

```shell
ssh-copy-id ~/.ssh/id_rsa.pub name@ip
```

可能遇到的问题 ```.ssh/authorized_keys: Permission denied```，此时可能是远端authorized_keys的文件权限不够至少需要600。chmod如果报错```chmod: changing permissions of ‘.ssh/authorized_keys’: Operation not permitted```则尝试查看文件属性```lsattr file```，然后去除多余属性```chattr -ia file```，之后再chmod，然后执行ssh-copy-id

## 配置

```shell
Host *
  TCPKeepAlive yes
  ServerAliveInterval 120
  ServerAliveCountMax 10

Host 博客(aliyun)
  HostName 101.132.40.106
  User root

```



## 杀不掉进程(显示defunct)

此时一般是父进程没有被杀掉，比如通过start.sh脚本启动了一个进程。

按下ctrl+c只是杀掉子进程，父进程start.sh还未被杀掉。此时应该找到父进程然后杀掉父进程。

查找父进程的方式为：

```shell
ps -xal  # 第4个字段即是父进程pid
```

