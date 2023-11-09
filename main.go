package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	mainTabs *container.AppTabs
)

func main() {
	a := app.New()
	mainWindow := a.NewWindow("Hello World")
	mainTabs = container.NewAppTabs()
	mainMenu := fyne.NewMainMenu()

	mainMenu.Items = append(mainMenu.Items, fyne.NewMenu("File",
		fyne.NewMenuItem("New Tab", appendNewTab),
	))

	mainWindow.SetMainMenu(mainMenu)
	mainWindow.SetContent(container.New(layout.NewStackLayout(), mainTabs))
	mainWindow.Resize(fyne.NewSize(1800, 900))
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}

func appendNewTab() {
	tab := container.NewTabItem(
		fmt.Sprintf("Tab %d", len(mainTabs.Items)+1),
		widget.NewLabel("new content"),
	)

	mainTabs.Append(tab)
}
