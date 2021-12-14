package app

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func AppInit() {
	icon, _ := walk.NewIconFromFile("icon.ico")
	var inTE *walk.TextEdit

	MainWindow{
		Title:   "Genesis DAT - Database admin tool",
		MinSize: Size{Width: 600, Height: 400},
		Icon:    icon,
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
				},
			},
		},
	}.Run()
}
