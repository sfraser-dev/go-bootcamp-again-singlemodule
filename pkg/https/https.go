package pkg_http

import (
	"io"
	"log"
	"net/http"
	"os"
)

type DataStruct struct {
	TheData []byte
}

func GettingGoogleHomepageHTML() {
	// get info from Google homepage
	res, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	//// Get and print the body (HTML) of the Google homepage
	//
	//// Read from res.Body to my own byte slice then printf
	// ds := DataStruct{
	// 	TheData: make([]byte, 32*1024), // a way to create a slice of specific size
	// }
	// if n, _:= res.Body.Read(ds.TheData); n >= (32*1024){
	// 	log.Fatalln(err)
	// }
	// res.Body.Close()
	// log.Printf("%s", ds.TheData)
	//
	//// Read from res.Body to inferred byte slice then println
	// if body, err := io.ReadAll(res.Body); err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	res.Body.Close()
	// 	log.Println(string(body))
	// }
	//
	//// Read from res.Body directly to stdout
	// if _, err := io.Copy(os.Stdout, res.Body); err != nil {
	// 	log.Fatalln(err)
	// }
	// res.Body.Close()
	//
	//// Read from res.Body to my own byte slice then write byte slice explicitly to stdout
	ds := DataStruct{
		TheData: make([]byte, 32*1024), // a way to create a slice of specific size
	}
	if n, _ := res.Body.Read(ds.TheData); n >= (32 * 1024) {
		log.Fatalln(err)
	}
	res.Body.Close()
	if _, err := io.WriteString(os.Stdout, string(ds.TheData)); err != nil {
		log.Fatalln(err)
	}
}
