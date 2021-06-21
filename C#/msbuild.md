## Mac下VS可以build，msbuild不通过的情况

### 触发条件

remove整个obj文件夹

### 修复

- ```shell
    msbuild ./project.csproj -t:Clean -p:Configuration=Debug -tv:Current
    ```

- ```shell
    msbuild ./project.csproj -t:restore -p:Configuration=Debug -tv:Current
    ```

- ```shell
    msbuild ./project.csproj -t:Build -p:Configuration=Debug -tv:Current
    ```

关键在于[restore packages](https://docs.microsoft.com/en-us/nuget/consume-packages/package-restore#restore-using-msbuild)