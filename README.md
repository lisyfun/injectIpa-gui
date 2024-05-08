# README

> 使用`wails`开发的`mac`工具,简化注入dylib操作,底层调用`injectipa`;
```shell
~/workspace/go/src/injectipa-gui/frontend 
❯ injectipa 
用法: 
injectipa ipa路径 dylib路径 (如有多个dylib依次填写在末尾即可)

参数: 
-i newIcon                      New icon to change.
-n newName                      New bundle name to change.
-b newBundleID                  New bundle id to change.
-a addResource                  Add new resource files.
-f enable FixIcons              Fix app thinning icon.
-w enable weakInject            Using weak mode inject.
-p enable needPlugIns           Keep all PlugIns files.
-s enable SandboxAccess         Usually no need to open.
-t enable removePrefsKey        remove prefs Key.

关于: 
injectipa v1.3-5 © 2018 - 2023 | Developed ♥︎ by Netskao | initnil.com 
```