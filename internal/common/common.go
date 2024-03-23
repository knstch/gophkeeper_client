package common

import (
	"net/http"

	"github.com/rivo/tview"
)

var (
	Pages        = tview.NewPages()
	App          = tview.NewApplication()
	Flex         = tview.NewFlex()
	ContactText  = tview.NewTextView()
	UserCookies  []*http.Cookie
	ReturnButton = tview.NewButton("Return to Main Menu")
)

func SetupEventHandling() {
	ReturnButton.SetSelectedFunc(func() {
		Pages.SwitchToPage("Main")
	})
}
