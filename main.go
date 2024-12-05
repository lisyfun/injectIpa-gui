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
		BackgroundColour: &options.RGBA{R: 28, G: 28, B: 28, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		DisableResize: true,
		AlwaysOnTop:   true,
		MinWidth:      230,
		MinHeight:     200,
		MaxWidth:      230,
		MaxHeight:     200,
	})

	if err != nil {
		runtime.LogErrorf(app.ctx, "Error:%v", err.Error())
	}
}
