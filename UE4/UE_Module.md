## 如何新建一个新的Module

- IMPLEMENT_MODULE

  实现module时的一个boilerplate。作用就是export初始化module的接口给外部调用，以及一些module公用的一些初始化操作(比如设置new delete malloc的实现)。

  ```c++
  	#define IMPLEMENT_MODULE( ModuleImplClass, ModuleName ) \
  		\
  		/**/ \
  		/* InitializeModule function, called by module manager after this module's DLL has been loaded */ \
  		/**/ \
  		/* @return	Returns an instance of this module */ \
  		/**/ \
  		extern "C" DLLEXPORT IModuleInterface* InitializeModule() \
  		{ \
  			return new ModuleImplClass(); /* export出去的接口 */ \
  		} \
  		/* Forced reference to this function is added by the linker to check that each module uses IMPLEMENT_MODULE */ \
  		extern "C" void IMPLEMENT_MODULE_##ModuleName() { } \
  		PER_MODULE_BOILERPLATE /* 设置new delete malloc等的实现 */ \
  		PER_MODULE_BOILERPLATE_ANYLINK(ModuleImplClass, ModuleName)
  ```

- 编写Build.cs

  定义DLL的依赖关系(依赖哪些DLL，使用哪些头文件)、以及添加编译用到的一些宏。最终由UBT来处理具体的编译处理。

## UE如何load一个module的

- load Module

  ```IModuleInterface* FModuleManager::LoadModuleWithFailureReason(const FName InModuleName, EModuleLoadResult& OutFailureReason``` 

  比如Sockets模块的加载

  ```	FModuleManager::Get().LoadModuleWithFailureReason(TEXT("Sockets"), LoadResult);```

  调用链如下：

  - ProjectManager

    (LaunchEngineLoop->) LoadModulesForProject -> LoadModulesForPhase -> LoadModuleWithFailureReason

  - PluginManager

    TryLoadModulesForPlugin -> LoadModulesForPhase -> LoadModuleWithFailureReason

  LoadingPhase的配置：

  - *.uproject

    ```
    {
    	"FileVersion": 3,
    	"EngineAssociation": "",
    	"Category": "",
    	"Description": "",
    	"Modules": [
    		{
    			"Name": "PlatformerGame",
    			"Type": "Runtime",
    			"LoadingPhase": "Default"
    		}
    	]
    }
    ```

  - *.uplugin

- 初始化Module

  调用IMPLEMENT_MODULE export出来的接口。即(ModuleManager.cpp):

  ```c++
  				FInitializeModuleFunctionPtr InitializeModuleFunctionPtr =
  					(FInitializeModuleFunctionPtr)FPlatformProcess::GetDllExport(ModuleInfo->Handle, TEXT("InitializeModule"));
  ```

  

