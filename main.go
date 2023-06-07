package main

import (
	"DebtSnowball2/backend/debt"
	"DebtSnowball2/backend/notes"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	snowball := debt.NewDebt("debt")
	snowballNotes := notes.NewNote("debt-notes", "")

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "debt-snowball2",
		Width:  1050,
		Height: 550,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			snowball,
			snowballNotes,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
