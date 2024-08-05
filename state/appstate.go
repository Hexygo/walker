package state

import (
	"context"
	"slices"

	"github.com/abenz1267/walker/config"
	"github.com/abenz1267/walker/modules"
	"github.com/abenz1267/walker/modules/clipboard"
	"github.com/junegunn/fzf/src/algo"
)

type AppState struct {
	Clipboard           modules.Workable
	IsDmenu             bool
	Dmenu               *modules.Dmenu
	DmenuSeparator      string
	DmenuLabelColumn    int
	ExplicitConfig      string
	ExplicitModules     []string
	ExplicitPlaceholder string
	ExplicitTheme       string
	ForcePrint          bool
	HasUI               bool
	IsRunning           bool
	IsService           bool
	KeepSort            bool
	Password            bool
	Benchmark           bool
	IsSingle            bool
	Labels              []string
	LabelsF             []string
	UsedLabels          []string
	InitialQuery        string
}

func Get() *AppState {
	algo.Init("default")

	return &AppState{
		IsService:      false,
		IsRunning:      false,
		HasUI:          false,
		ExplicitConfig: "config.json",
	}
}

func (app *AppState) StartServiceableModules(cfg *config.Config) {
	cfg.IsService = true

	app.Clipboard = &clipboard.Clipboard{}
	app.Dmenu = &modules.Dmenu{}

	app.Clipboard.Setup(cfg)
	app.Dmenu.Setup(cfg)

	if !slices.Contains(cfg.Disabled, app.Clipboard.General().Name) {
		app.Clipboard.SetupData(cfg, context.Background())
	}

	if !slices.Contains(cfg.Disabled, app.Dmenu.General().Name) {
		app.Dmenu.SetupData(cfg, context.Background())
	}
}
