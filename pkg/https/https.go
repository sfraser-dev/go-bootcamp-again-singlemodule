package pkg_http

import (
	"fmt"
	"net/http"
	"os"
)

type DataStruct struct {
	TheData []byte
}

func GettingGoogleHomepageHTML() {
	// the data structure
	ds := DataStruct{
		TheData: make([]byte, 99999),
	}

	// get info from Google homepage
	res, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Get and print the body (HTML) of the Google homepage
	res.Body.Read(ds.TheData)
	fmt.Printf("%s", ds.TheData)
}
