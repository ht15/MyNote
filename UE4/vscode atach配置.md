```shell
	{
		"name": "ShooterGameEditor (DebugGame)",
		"request": "launch",
		"preLaunchTask": "ShooterGameEditor Mac DebugGame Build",
		"program": "/PATH/unreal_engine/UnrealEngine-4.26.2-release/Engine/Binaries/Mac/UE4Editor-Mac-DebugGame.app/Contents/MacOS/UE4Editor-Mac-DebugGame",
		"args": [
			"/PATH/UnrealProjects/ShooterGame/ShooterGame.uproject"
		],
		"cwd": "/PATH/unreal_engine/UnrealEngine-4.26.2-release",
		"type": "lldb"
	},
	{
		"name": "attach ShooterGameEditor (DebugGame)",
		"request": "attach",
		"program": "/PATH/unreal_engine/UnrealEngine-4.26.2-release/Engine/Binaries/Mac/UE4Editor-Mac-DebugGame.app/Contents/MacOS/UE4Editor-Mac-DebugGame",
		"type": "cppdbg",
		"processId": "72968",
		"MIMode": "lldb"
	},
```
