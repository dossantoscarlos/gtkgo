package views

import (
	"fmt"
	"gtkgo/core/adapters/controllers"
	"gtkgo/core/domain/entities"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	registerScreen = "Register User"
)

type Register struct {
	app           fyne.App
	window        fyne.Window
	mainContainer *fyne.Container
	userName      *widget.Entry
	emailEntry    *widget.Entry
	passwordEntry *widget.Entry
	submit        *widget.Button
}

func NewRegisterWindow(app fyne.App) *Register {

	fmt.Println("NewRegisterWindow")

	screen := app.NewWindow(registerScreen)
	screen.CenterOnScreen()

	screen.Resize(fyne.NewSize(400, 300))

	return &Register{
		app:    app,
		window: screen,
	}
}

func (rg *Register) BuildAndShow() {
	rg.setMainContainer()
	rg.window.SetContent(rg.mainContainer)
	rg.window.Show()
}

func (rg *Register) setMainContainer() {
	var objs []fyne.CanvasObject
	rg.setUserName()
	rg.setEmailEntry()
	rg.setPasswordEntry()
	rg.setSubmitButton()
	objs = append(
		objs,
		rg.userName,
		rg.emailEntry,
		rg.passwordEntry,
		rg.submit,
	)
	rg.mainContainer = container.NewVBox(objs...)
}

func (rg *Register) setUserName() {
	rg.userName = widget.NewEntry()
	rg.userName.SetPlaceHolder("Nome")
}

func (rg *Register) setEmailEntry() {
	rg.emailEntry = widget.NewEntry()
	rg.emailEntry.SetPlaceHolder("E-mail")
}

func (rg *Register) setPasswordEntry() {
	rg.passwordEntry = widget.NewPasswordEntry()
	rg.passwordEntry.SetPlaceHolder("Senha")
}

func (rg *Register) setSubmitButton() {
	rg.submit = widget.NewButton("Entrar", func() {
		register(rg)
	})
}

func register(rg *Register) {
	var err error
	var user *entities.User

	ctrl := controllers.NewUserController()
	userName := rg.userName.Text
	email := rg.emailEntry.Text
	password := rg.passwordEntry.Text

	if userName == "" || email == "" || password == "" {
		dialog.ShowError(fmt.Errorf("Preencha todos os campos"), rg.window)
		return
	}

	user, err = ctrl.HandleCreateUser(userName, email, password)
	if err != nil {
		dialog.NewError(fmt.Errorf("Erro ao criar usuário: %v", err), rg.window)
	}

	fmt.Printf("Usuário criado com sucesso: %v\n", user)

	fmt.Println(email, password)

	limpaCampos(rg)
}

func limpaCampos(rg *Register) {
	rg.userName.SetText("")
	rg.emailEntry.SetText("")
	rg.passwordEntry.SetText("")
}
