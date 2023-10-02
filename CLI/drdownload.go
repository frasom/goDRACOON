// DRACOON API
// retrieve the information of a DRACOON downlodlink and
// download the file if desired
// from a DRACOON System
// API ist documented in support.dracoon.com
// Autor: f.sommer@dracoon.com
// Version 0.1.1
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
var response int

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

type Download_URL struct {
	DownloadURL string `json:"downloadUrl"`
}

func main() {

	// defining a struct instance
	var Responses1 Download_Share_information
	var Responses2 Download_URL

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

	// Request public Download Share information
	// build API call to retrieve the public information of a Download Share

	linkURL := "https://" + u.Host               // URL
	request = "/api/v4/public/shares/downloads/" // Request for public Download Share information
	apicall = linkURL + request + accessKey      // build API call
	curlget()                                    // execute curl commando

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

	// ask for file download
	fmt.Printf("Download File [yes/No] : ")
	fmt.Scanf("%c", &response) // user input

	switch response {
	default:
		os.Exit(0)
	case 'y', 'Y':
		// get the download url
		curlpost() // execute curl commando

		// decoding JSON string Download URL
		err = json.Unmarshal(bodyText, &Responses2)
		if err != nil {
			fmt.Println("Error decoding JSON string: ", err) // if error is not nil print error
		}

		// download the file
		err = downloadFile(Responses1.Filename, Responses2.DownloadURL)
		if err != nil {
			fmt.Println("Error Downloading File: ", err) // if error is not nil print error
		}
	}

}

func curlget() { // Execute the Curl GET Command
	client := &http.Client{}
	req, err := http.NewRequest("GET", apicall, nil)
	if err != nil {
		fmt.Println("Error execut Curl Command: ", err) // if error is not nil print error
	}
	// set headers for get request
	req.Header.Set("accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error set headers for GET request: ", err) // if error is not nil print error
	}
	// closed response body
	defer resp.Body.Close()
	//Read and parse response body
	bodyText, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error closed response body: ", err) // if error is not nil print error
	}
}

func curlpost() { // Execute the Curl POST Command
	client := &http.Client{}
	req, err := http.NewRequest("POST", apicall, nil)
	if err != nil {
		fmt.Println("Error execut Curl POST Command: ", err) // if error is not nil print error
	}
	// set headers for get request
	req.Header.Set("accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error set headers for POST request: ", err) // if error is not nil print error
	}
	// closed response body
	defer resp.Body.Close()
	//Read and parse response body
	bodyText, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error closed response body: ", err) // if error is not nil print error
	}
}

func downloadFile(filename string, url string) (err error) {

	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
