// DRACOON API
// get software version and system information
// from a DRACOON System fyne-gui version
// API ist documented in support.dracoon.com
// Autor: f.sommer@dracoon.com
// Version 0.1.0
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// defining globale variables
var request string
var apicall string
var bodyText []byte
var response string

// JSON type definition
type software_version struct {
	RestAPIVersion    string    `json:"restApiVersion"`
	SdsServerVersion  string    `json:"sdsServerVersion"`
	BuildDate         time.Time `json:"buildDate"`
	ScmRevisionNumber string    `json:"scmRevisionNumber"`
	IsDracoonCloud    bool      `json:"isDracoonCloud"`
}

type system_info struct {
	LanguageDefault       string   `json:"languageDefault"`
	HideLoginInputFields  bool     `json:"hideLoginInputFields"`
	S3Hosts               []string `json:"s3Hosts"`
	S3EnforceDirectUpload bool     `json:"s3EnforceDirectUpload"`
	UseS3Storage          bool     `json:"useS3Storage"`
	AuthMethods           []struct {
		Name      string `json:"name"`
		IsEnabled bool   `json:"isEnabled"`
		Priority  int    `json:"priority"`
	} `json:"authMethods"`
}

func main() {

	// defining a struct instance
	var Responses1 system_info
	var Responses2 software_version

	a := app.New()
	w := a.NewWindow("DRACOON Status")

	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Cloud URL eingeben")

	// set the labels as placeholders
	resultLabel := widget.NewLabel("System :\nDracoon Type:\nSprache:\nAPI Version:\nServer Version:\nUse S3 Storage:\nS3 Hosts:\nAuth. Methods:")

	// Query URL
	sendButton := widget.NewButton("Abfragen", func() {
		baseURL := urlEntry.Text
		if baseURL == "" {
			resultLabel.SetText("Bitte geben Sie eine Cloud URL ein")
			return
		}
		// check if is a valid URL
		_, err := url.ParseRequestURI(baseURL)
		if err != nil {
			resultLabel.SetText("Bitte geben Sie eine gültige URL ein")
			return
		}
		u, err := url.Parse(baseURL)
		if err != nil || u.Scheme == "" || u.Host == "" {
			resultLabel.SetText("Keine gültige URL")
		}

		// build API call for system info
		request = "/api/v4/public/system/info?is_enabled=true" // Request for system information
		apicall = baseURL + request                            // build API call
		curlcmd()

		// decoding JSON string system info
		err = json.Unmarshal(bodyText, &Responses1)
		if err != nil {
			fmt.Println("Error decoding JSON string: ", err) // if error is not nil print error
		}

		// build API call for software version
		request = "/api/v4/public/software/version" // Request for system information
		apicall = baseURL + request                 // build API call
		curlcmd()                                   // execute curl commando

		// decoding JSON string software version
		err = json.Unmarshal(bodyText, &Responses2)
		if err != nil {
			fmt.Println("Error decoding JSON string: ", err) // if error is not nil print error
		}

		// Set System Type String

		// build output from JSON strings
		response := "System :\t\t" + baseURL + "\nDracoon Type:\t"
		// Set System Type String
		if Responses2.IsDracoonCloud == true {
			response = response + "Cloud"
		} else {
			response = response + "On-Premises"
		}
		response = response + "\nSprache:\t\t" + Responses1.LanguageDefault + "\nAPI Version:\t\t" + Responses2.RestAPIVersion + "\nServer Version:\t" + Responses2.SdsServerVersion + "\nUse S3 Storage:\t"
		// Set S3 Storage String
		if Responses1.UseS3Storage == true {
			response = response + "Yes"
		} else {
			response = response + "No"
		}
		response = response + "\nS3 Hosts:\t\t"
		for i := range Responses1.S3Hosts {
			response = response + Responses1.S3Hosts[i] + " "
		}
		response = response + "\nAuth. Methods:\t"
		for i := range Responses1.AuthMethods {
			response = response + Responses1.AuthMethods[i].Name + " "
		}

		resultLabel.SetText(string(response))
	})

	// show status
	content := container.NewVBox(
		widget.NewLabel("DRACOON System Status"),
		urlEntry,
		sendButton,
		resultLabel,
	)

	w.SetContent(content)
	w.ShowAndRun()
}

func curlcmd() { // Execute the Curl Command
	client := &http.Client{}
	req, err := http.NewRequest("GET", apicall, nil)
	if err != nil {
		fmt.Println("Error execut Curl Command: ", err) // if error is not nil print error
	}
	// set headers for get request
	req.Header.Set("accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error set headers for get request: ", err) // if error is not nil print error
	}
	// closed response body
	defer resp.Body.Close()
	//Read and parse response body
	bodyText, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error closed response body: ", err) // if error is not nil print error
	}
}
