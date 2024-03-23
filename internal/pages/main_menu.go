package pages

import (
	"github.com/knstch/gophkeeper_client/internal/common"
	"github.com/rivo/tview"
)

var (
	MainMenu = tview.NewModal()
)

func MainPage(pages *tview.Pages) {
	common.Flex.Clear()
	common.Flex.AddItem(mainMenu(), 0, 1, true)
	AuthPage(common.Pages)
	pages.SwitchToPage("Main")
}

func mainMenu() *tview.Modal {
	MainMenu.SetBorder(true).SetTitle("Main menu")
	MainMenu.AddButtons([]string{"Auth", "Quit"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonIndex == 0 {
			common.Pages.SwitchToPage("Auth")
		} else {
			common.App.Stop()
		}
	})
	return MainMenu
}
