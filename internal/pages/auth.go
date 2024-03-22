package pages

import (
	"github.com/go-resty/resty/v2"
	"github.com/knstch/gophkeeper_client/internal/common"
	"github.com/rivo/tview"
)

type credentials struct {
	email    string
	password string
}

type Message struct {
	Msg string `json:"message"`
}

var (
	AuthForm = tview.NewForm()
)

func AuthPage(pages *tview.Pages) {
	common.Flex.Clear()
	common.Flex.AddItem(authForm(), 0, 1, true)
	pages.AddPage("Auth", common.Flex, true, true)
	pages.SwitchToPage("Auth")
}

func authForm() *tview.Form {
	client := resty.New().SetBaseURL("http://localhost:8080")

	credentials := &credentials{}

	AuthForm.AddInputField("email", "", 20, nil, func(email string) {
		credentials.email = email
	})

	AuthForm.AddInputField("password", "", 20, nil, func(password string) {
		credentials.password = password
	})

	responseTextView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	var goodMsg Message
	AuthForm.AddButton("Log in", func() {
		resp, err := client.R().SetBody(prepareBody(*credentials)).SetResult(&goodMsg).Post("/auth")
		if err != nil {
			responseTextView.SetText(err.Error())
			return
		}
		responseTextView.SetText(string(resp.Body()))
	})

	AuthForm.AddFormItem(responseTextView)

	return AuthForm
}

func prepareBody(user credentials) string {
	return `{"email": "` + user.email + `","password": "` + user.password + `"}`
}
