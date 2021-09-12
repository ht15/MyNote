```
cmake_minimum_required(VERSION x.x)
source_group(// FILES ${src})

add_library(libname MODULE ${src})  # MODULE类型只能dlopen动态加载，不能link
target_include_directories(libname PUBLIC include PRIVATE src)

target_link_libraries(libname ...)
target_compile_definitions(libname PUBLIC DEFINE)

install(DIRECTORY include/ DESTINATION ${INCLUDE_INSTALL_DIR})
install(TARGETS libname LIBRARY DESTINATION ${RUNTIME_INSTALL_DIR})

```

