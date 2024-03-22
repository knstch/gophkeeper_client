package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/knstch/gophkeeper_client/internal/common"
	"github.com/knstch/gophkeeper_client/internal/pages"
	"github.com/rivo/tview"
)

var contactText = tview.NewTextView()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) to add a new contactq \n(q) to quit")

func main() {
	common.Flex.SetDirection(tview.FlexRow).AddItem(text, 0, 1, false)

	common.Flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if common.Pages.HasPage("Auth") {
			if event.Rune() == 'a' {
				return event
			}
		} else {
			if event.Rune() == 'q' {
				common.App.Stop()
			} else if event.Rune() == 'a' {
				pages.AuthPage(common.Pages)
			}
		}
		return event
	})

	common.Pages.AddPage("Menu", common.Flex, true, true)

	if err := common.App.SetRoot(common.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
