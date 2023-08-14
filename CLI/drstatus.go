// DRACOON API
// get software version and system information
// from a DRACOON System
// API ist documentet in support.dracoon.com
// Autor: f.sommer@dracoon.com
// Version 0.1.0
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// defining globale variables
var request string
var apicall string
var bodyText []byte

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

	// checke args?
	if len(os.Args) == 1 {
		fmt.Println("DRACOON URL expected as a parameter")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	// check if args(1) is a valid URL
	_, err := url.ParseRequestURI(baseURL)
	if err != nil {
		panic(err) // if URL not valid print error and exit
	}
	u, err := url.Parse(baseURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		panic(err) // URL problem print error and exit
	}

	// build API call for system info
	request = "/api/v4/public/system/info?is_enabled=true" // Request for system information
	apicall = baseURL + request                            // build API call
	curlcmd()                                              // execute curl commando

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

	// printing details of JSON strings
	fmt.Printf("System :        %s\n", baseURL)
	fmt.Printf("Dracoon Cloud:  ")
	fmt.Printf("%v\n", Responses2.IsDracoonCloud) //"true or false"
	fmt.Printf("Sprache:        %s\n", Responses1.LanguageDefault)
	fmt.Printf("API Version:    %s\n", Responses2.RestAPIVersion)
	fmt.Printf("Server Version: %s\n", Responses2.SdsServerVersion)
	fmt.Printf("Use S3 Storage: ")
	fmt.Printf("%v\n", Responses1.UseS3Storage) //"true or false"
	fmt.Printf("S3 Hosts:       %s\n", Responses1.S3Hosts)
	fmt.Printf("Auth. Methods:  ")
	for i := range Responses1.AuthMethods {
		fmt.Printf("%s ", Responses1.AuthMethods[i].Name)
	}
	fmt.Printf("\n")
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
