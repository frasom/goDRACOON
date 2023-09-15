// DRACOON API
// retrieve the information of a DRACOON downlodlink and
// download the file "dreams of the future"
// from a DRACOON System
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
	"os"
	"strings"
)

// defining globale variables
var request string
var apicall string
var bodyText []byte

// JSON type definition
type Download_Share_information struct {
	Protected     bool   `json:"isProtected"`
	Filename      string `json:"fileName"`
	Filesize      int    `json:"size"`
	Limit         bool   `json:"limitReached"`
	CreatorName   string `json:"creatorName"`
	CreatedAt     string `json:"createdAt"`
	Downloadlimit bool   `json:"hasDownloadLimit"`
	Filetype      string `json:"mediaType"`
	Freigabename  string `json:"name"`
	Ablauf        string `json:"expireAt"`
	Notes         string `json:"notes"`
	Encrypted     bool   `json:"isEncrypted"`
}

func main() {

	// defining a struct instance
	var Responses1 Download_Share_information

	// checke args?
	if len(os.Args) == 1 {
		fmt.Println("DRACOON URL expected as a parameter")
		os.Exit(1)
	}

	// URL Parser

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

	cuttingField := strings.FieldsFunc(u.Path, func(r rune) bool {
		if r == '/' {
			return true
		}
		return false
	})

	// Link Key extraction
	cutLength := len(cuttingField)         // number of fields
	accessKey := cuttingField[cutLength-1] // last field is LinkKey

	// build API call for system info
	// aufbereiten linkURL

	linkURL := "https://" + u.Host               // URL
	request = "/api/v4/public/shares/downloads/" // Request for system information
	apicall = linkURL + request + accessKey      // build API call
	curlcmd()                                    // execute curl commando

	// decoding JSON string Download Share Information
	err = json.Unmarshal(bodyText, &Responses1)
	if err != nil {
		fmt.Println("Error decoding JSON string: ", err) // if error is not nil print error
	}

	// printing details of JSON strings
	fmt.Printf("Filename:        %s\n", Responses1.Filename)
	fmt.Printf("Filetype:        %s\n", Responses1.Filetype)
	fmt.Printf("Filesize:        %d\n", Responses1.Filesize)
	fmt.Printf("Von:             %s\n", Responses1.CreatorName)
	fmt.Printf("Linkname:        %s\n", Responses1.Freigabename)
	fmt.Printf("Notes:           %s\n", Responses1.Notes)
	fmt.Printf("Protected:       ")
	fmt.Printf("%v\n", Responses1.Protected) //"true or false"
	fmt.Printf("Encrypted:       ")
	fmt.Printf("%v\n", Responses1.Encrypted) //"true or false"
	fmt.Printf("Downloadlimit:   ")
	fmt.Printf("%v\n", Responses1.Downloadlimit) //"true or false"
	fmt.Printf("\n")
}

// download the file "dreams of the future"

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
