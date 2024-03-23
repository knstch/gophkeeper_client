package pages

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/knstch/gophkeeper_client/internal/common"
	"github.com/knstch/gophkeeper_client/internal/config"
	"github.com/rivo/tview"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ServerResponse struct {
	Msg   string `json:"message"`
	Error string `json:"error"`
	Code  int    `json:"code"`
}

var (
	AuthForm       = tview.NewForm()
	serverResponse ServerResponse
)

func AuthPage(pages *tview.Pages) {
	common.Flex.Clear()
	common.Flex.AddItem(authForm(), 0, 1, true).SetBorder(true).SetTitle("Auth form")
	pages.AddPage("Auth", common.Flex, true, true)
	pages.SwitchToPage("Auth")
}

func authForm() *tview.Form {
	credentials := &credentials{}

	AuthForm.AddInputField("Email", "", 36, nil, func(email string) {
		credentials.Email = email
	})

	AuthForm.AddInputField("Password", "", 36, nil, func(password string) {
		credentials.Password = password
	})

	AuthForm.AddButton("Log in", func() {
		requestBody, err := json.Marshal(credentials)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post(config.BaseURL+config.AuthURL, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&serverResponse); err != nil {
			log.Fatal(err)
		}

		responseReceiver(AuthForm, resp.StatusCode, false)
	})
	AuthForm.AddButton("Register", func() {
		requestBody, err := json.Marshal(credentials)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post(config.BaseURL+config.RegisterURL, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&serverResponse); err != nil {
			log.Fatal(err)
		}

		responseReceiver(AuthForm, resp.StatusCode, true)
	})
	AuthForm.AddButton("Back to main menu", func() {
		common.Pages.SwitchToPage("Main")
	})
	return AuthForm
}

func responseReceiver(authForm *tview.Form, code int, isNew bool) {
	idx := authForm.GetFormItemIndex("Response")
	if idx != -1 {
		authForm.RemoveFormItem(idx)
	}

	switch code {
	case 200:
		if isNew {
			authForm.AddTextArea("Response", "Вы успешно зарегестрировались", 36, 2, 30, nil)
			return
		}
		authForm.AddTextArea("Response", "Вы успешно залогинились", 36, 2, 30, nil)
	case 400:
		authForm.AddTextArea("Response", serverResponse.Error, 36, 7, 30, nil)
	case 404:
		authForm.AddTextArea("Response", "Неверный логин или пароль", 36, 2, 30, nil)
	case 409:
		authForm.AddTextArea("Response", "Эта почта уже занята", 36, 2, 30, nil)
	default:
		authForm.AddTextArea("Response", "Упс, произошла непредвиденная ошибка", 36, 2, 30, nil)
	}
}
