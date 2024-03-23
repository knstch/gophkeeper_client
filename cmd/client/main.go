package main

import (
	"github.com/knstch/gophkeeper_client/internal/common"
	"github.com/knstch/gophkeeper_client/internal/pages"
)

func main() {
	pages.MainPage(common.Pages)
	common.Pages.AddPage("Main", pages.MainMenu, true, true)

	if err := common.App.SetRoot(common.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
