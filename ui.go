package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	CpanelAPIKey            string   `json:"cpanelAPIKey"`
	CpanelHost              string   `json:"cpanelHost"`
	DefaultForwardersDomain string   `json:"defaultForwardersDomain"`
	YourRealEmail           string   `json:yourRealEmail`
	Domains                 []string `json:"domains"`
}

func loadConfig() (Config, error) {
	config := Config{}

	if config.CpanelAPIKey == "" || config.CpanelHost == "" || config.DefaultForwardersDomain == "" || len(config.Domains) == 0 || config.YourRealEmail == "" {
		file, err := ioutil.ReadFile("config/config.json")
		if err != nil {
			return config, err
		}

		err = json.Unmarshal(file, &config)
		if err != nil {
			return config, err
		}
	}

	return config, nil
}

func createUI(myWindow fyne.Window) *fyne.Container {

	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	cpanelAPIKey := config.CpanelAPIKey
	cpanelHost := config.CpanelHost
	defaultForwardersDomain := config.DefaultForwardersDomain
	yourRealEmail := config.YourRealEmail
	domains := config.Domains
	alias := ""

	label1 := widget.NewLabel("alias:")
	input := widget.NewEntry()

	grid := container.New(layout.NewFormLayout(), label1, input)

	defaultForwardersDomain = domains[0]
	defaultEmail := yourRealEmail

	domainSelect := widget.NewSelect(domains, func(selectedDomain string) {
		defaultForwardersDomain = selectedDomain
	})
	domainSelect.SetSelected(defaultForwardersDomain)

	content := container.NewVBox(
		grid,
		container.NewHBox(widget.NewLabel("Select Domain:"), domainSelect),
		widget.NewButton("Create!", func() {
			returnvalue := AddAlias(strings.ToLower(input.Text), cpanelHost, defaultForwardersDomain, cpanelAPIKey, defaultEmail)
			alias = returnvalue
			runPopUp(myWindow, fmt.Sprintf("%v", returnvalue), "Copy and Close", func() {
				clipboard := myWindow.Clipboard()
				clipboard.SetContent(alias)
			})
			log.Println("Content was:", input.Text)
		}),
		widget.NewButton("Search!", func() {
			returnvalue := FindAlias(strings.ToLower(input.Text), cpanelHost, cpanelAPIKey, domains[domainSelect.SelectedIndex()])
			runPopUp(myWindow, fmt.Sprintf("%v", returnvalue), "Close", func() {})
			log.Println("Content was:", input.Text)
		}),
		widget.NewButton("Delete!", func() {
			returnvalue := TrashAlias(strings.ToLower(input.Text), cpanelHost, defaultForwardersDomain, cpanelAPIKey, yourRealEmail)

			if returnvalue == 200 {
				runPopUp(myWindow, "Success!", "Close", func() {})
			} else {
				runPopUp(myWindow, "Fail!", "Close", func() {})
			}

			log.Println("Content was:", input.Text)
		}),
	)

	return content
}
