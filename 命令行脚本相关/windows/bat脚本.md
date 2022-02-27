## if else

```powershell
if "%BuildType%"=="Debug" (
	echo "build with debug"
) ^
else if "%BuildType%"=="Release" (
	echo "build with release"
) ^
else (
	echo "error"
)
```

## parse command line

```powershell
:parse_loop
IF NOT "%1"=="" (
    IF "%1"=="-arg1" (
    	echo "parse arg1"
      SHIFT
    )
    IF "%1"=="-arg2" (
    	echo "parse arg2"
    )
    SHIFT
    GOTO parse_loop
)
```

## string replace

```powershell
set DST_FILE="dst.txt"
set SRC_FILE="src.txt"
type NUL > %DST_FILE%
setlocal disabledelayedexpansion
for /f "delims=" %%A in ('findstr /N "^" %SRC_FILE%') do (
    set line=%%A
    setlocal enabledelayedexpansion
    set "line=!line:%replace_str%=%dst_str%!"
>>%DST_FILE% echo(!line:*:=!
    endlocal
)

e.g.2
for /f "delims=" %%i in ('git status -s ') do set b=%%i
if "%b%"=="" (
	echo "working tree clean"
) ^
else (
	echo "working tree dirty"
)
```

