package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type aliases struct {
	Number int `json:"number"`
}

type CpanelResponse struct {
	Metadata struct {
		Command string `json:"command"`
		Reason  string `json:"reason"`
		Version int    `json:"version"`
	} `json:"metadata"`
	Status int           `json:"status"`
	Errors []CpanelError `json:"errors"`
}

type CpanelError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func runPopUp(w fyne.Window, m string, btnText string, callback func()) (modal *widget.PopUp) {
	modal = widget.NewModalPopUp(
		container.NewVBox(
			widget.NewLabel(m),
			widget.NewButton(btnText, func() {
				callback()
				modal.Hide()
			}),
		),
		w.Canvas(),
	)
	modal.Show()
	return modal
}
func FindAlias(alias string, cpanelHost string, cpanelAPIKey string, selectedDomain string) string {
	url := "https://" + cpanelHost + "/execute/Email/list_forwarders?domain=" + selectedDomain
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Set("Authorization", "cpanel "+cpanelAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "No Body to Search"
	}

	var forwarders map[string]interface{}
	err = json.Unmarshal(body, &forwarders)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return "No Aliases Found"
	}

	forwardersFound := false
	for _, v := range forwarders["data"].([]interface{}) {
		forwarder := v.(map[string]interface{})
		if forwarder["dest"] == fmt.Sprintf("%s@%s", alias, selectedDomain) {
			fmt.Println("Found forwarder:", forwarder)
			forwardersFound = true
			return fmt.Sprintf("Found: %s@%s", alias, selectedDomain)
		}
	}

	if !forwardersFound {
		return "Forwarder not found"
	}

	return "Error occurred"
}

func TrashAlias(alias string, cpanelHost string, forwardersDomain string, cpanelAPIKey string, realEmail string) int {

	url1 := "https://" + cpanelHost + "/execute/Email/delete_forwarder?address="
	url2 := "%40" + forwardersDomain + "&forwarder=" + realEmail
	result := url1 + alias + url2
	client := &http.Client{}
	req, err := http.NewRequest("GET", result, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Set("Authorization", "cpanel "+cpanelAPIKey)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Printf("%+v\n", resp.StatusCode, resp.Body)
	log.Println("URL:", req.URL)
	return resp.StatusCode
}

func AddAlias(alias string, cpanelHost string, forwardersDomain string, cpanelAPIKey string, defaultEmail string) string {
	url := "https://" + cpanelHost + "/execute/Email/add_forwarder?domain=" + forwardersDomain + "&email=" + alias + "&fwdemail=" + url.QueryEscape(defaultEmail) + "&fwdopt=fwd"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
		return "Creation Failed"
	}

	req.Header.Set("Authorization", "cpanel "+cpanelAPIKey)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Printf("%+v\n", resp.StatusCode)
	log.Println("URL:", req.URL)
	return alias + "@" + forwardersDomain
}
