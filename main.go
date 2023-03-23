package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Email Alias")

	content := createUI(myWindow)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()
}
