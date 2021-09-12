```
cmake_minimum_required(VERSION x.x) # 如3.12

project(project_name VERSION x.x)

if(UNIX AND NOT APPLE)
    set(LINUX TRUE)
endif()

SET_PROPERTY(GLOBAL PROPERTY USE_FOLDERS ON) # 要配合 SET_PROPERTY(TARGET target_name PROPERTY FOLDER "fold_name")使用

if (NOT MY_CMAKE_BUILD_TYPE)
    set(MY_CMAKE_BUILD_TYPE Debug)
endif()

if(WIN32)
    add_compile_options(/wd4146)  # 编译选项
    add_definitions(-DDEFINE_NAME) # 宏定义
    set(CMAKE_CXX_STANDARD 17)  # 设置c++标准
elseif(APPLE)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++17") # 设置c++标准
elseif(LINUX)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++17)
else()
    message(STATUS "unknown os find!")
endif()
include(cmake/module.cmake) # 相当于引入头文件。一般用于全局函数、第三方库的定义
add_subdirectory(thridparty)
add_subdirectory(sub_dir)
```

