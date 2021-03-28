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

