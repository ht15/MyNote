````
function(fun_name VARS... OUTPUT)
    set(${OUTPUT} ...)
endfunction()

# 组织第三方库
add_library(lib_3rd STATIC IMPORTED GLOBAL)
set_target_properties(lib_3rd PROPERTIES
    IMPORTED_LOCATION
  	*.lib
    INTERFACE_INCLUDE_DIRECTORIES
    */include
)
````

