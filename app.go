package main

import (
	"bytes"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	"path"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved, so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// 设置窗口启动位置
	runtime.WindowSetPosition(a.ctx, 1151, 118)
}

// OpenFileDlgAndGenerateFile 打开上传文件窗口,并生成文件
func (a *App) OpenFileDlgAndGenerateFile() {
	opts := runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Ipa (*.ipa)",
				Pattern:     "*.ipa",
			},
			{
				DisplayName: "Dylib (*.Dylib)",
				Pattern:     "*.dylib",
			},
		},
	}

	filePaths, err := runtime.OpenMultipleFilesDialog(a.ctx, opts)
	if err != nil {
		runtime.LogErrorf(context.Background(), "error: %v", err.Error())
	}
	var ipaFilePath string
	var dylibFilePaths []string

	for _, filePath := range filePaths {
		fileSuffix := strings.ToLower(path.Ext(filePath))
		switch fileSuffix {
		case ".ipa":
			ipaFilePath = filePath
		case ".dylib":
			dylibFilePaths = append(dylibFilePaths, filePath)
		}
	}

	args := []string{ipaFilePath}
	args = append(args, dylibFilePaths...)

	// 执行命令时必须指定全路径
	cmd := exec.Command("/usr/local/bin/injectipa", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		runtime.LogErrorf(a.ctx, "%s\n", args)
	}
}
