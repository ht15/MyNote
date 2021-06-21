```shell
#!/bin/sh
name=$1
pass=$2
echo "you are setting username : ${name}"
echo "you are setting password : $pass for ${name}"
sudo useradd $name

if [ $? -eq 0 ];then
   echo "user ${name} is created successfully!!!"
else
   echo "user ${name} is created failly!!!"
   exit 1
fi

echo $pass | sudo passwd $name --stdin  &>/dev/null
if [ $? -eq 0 ];then
   echo "${name}'s password is set successfully"
else
   echo "${name}'s password is set failly!!!"
fi
```

