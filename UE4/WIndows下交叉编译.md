1.修改[YourProject]\Config\DefaultEngine.ini，添加：[/Script/LinuxTargetPlatform.LinuxTargetSettings]

​	TargetArchitecture=X86_64UnknownLinuxGnu

2.右键[YourProject]\[YourProjectName].uproject执行GenerateVisiualStudioProjectFiles重新生成sln文件。

​	这时打开sln文件后，在解决方案平台复选框里就有Linux选项了。

[参考](https://zhuanlan.zhihu.com/p/74057757)



## 设置编译选项

Engine\Source\Programs\UnrealBuildTool\Platform\Linux\LinuxToolChain.cs

## uint64_t与UE::uint64不兼容

最好用stdint.h uint_least64_t