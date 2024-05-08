package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title: "injectIpa-gui",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		DisableResize: true,
		AlwaysOnTop:   true,
		MinWidth:      306,
		MinHeight:     241,
		MaxWidth:      306,
		MaxHeight:     241,
	})

	if err != nil {
		runtime.LogErrorf(app.ctx, "Error:%v", err.Error())
	}
}
