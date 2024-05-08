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
## 使用方法
### 1. 下载源码
```sh
git clone https://github.com/lisyfun/injectIpa-gui.git
```
### 2. 安装项目中的`pkg/injectipa.pkg` ,目前只支持 mac 系统
### 3. 加载前端依赖
 ```sh
 cd frontend && yarn && cd ..
 ```
### 4. 启动项目
```sh
wails dev
```
### 5. 构建项目
```sh
wails build
```

## 注意
> 使用`which injectipa`查询`injectipa`的可执行路径  
> 如果路径不是在`/usr/local/bin/injectipa`下,则需要修改源码中的路径
> 
#### `app.go`
```go
- 64	// 执行命令时必须指定全路径
- 65	cmd := exec.Command("/usr/local/bin/injectipa", args...)
```
## 项目效果
![injectIpa-gui](https://cdn.jsdelivr.net/gh/lisyfun/injectIpa-gui/doc/injectIpa-gui.gif)
