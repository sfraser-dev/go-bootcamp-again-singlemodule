package pkg_interfaces_read_file

import (
	"fmt"
	"io"
	"log"
	"os"
)

// using interfaces in std lib to complete this assignment more simply
func RunIt() {
	// os.Args hold the command-line arguments, starting with the program name
	if len(os.Args) == 2 {
		filename := os.Args[1]

		if fp, err := os.Open(filename); err != nil {
			log.Fatalln(err)
		} else {
			defer (*fp).Close() // close file via defer (out of scope or EOF)
			defer fmt.Println("closed file " + filename + " via defer")

			// buffer := make([]byte, 32*1024)
			// if _, err := (*fp).Read(buffer); err != nil {
			// 	log.Fatalln(err)
			// }
			// log.Println("\n" + string(buffer))
			//
			// Can see that fp is a pointer to os.File type
			// Can see that os.File type has method Read(b)(n,err), ie: (f)Read(b)(n,err)
			// From interfaces_http_get, we remember that Read(b)(n,err) is the only signature of the **Reader interface**
			// From interfaces_http_get, can pass this Reader interface to io.Copy()
			// This simplifies our code as there is no need to create and then print a buffer
			//
			// Note: a **pointer** to os.File type does implement the io.Reader interface, **variable** (*os.File) does not
			// Note: so we cannot dereference fp, we must pass in fp itself.
			if _, err := io.Copy(os.Stdout, fp); err != nil {
				log.Fatalln(err)
			}
		}
	} else {
		fmt.Printf("No input file on command line...\nUsage: go run %+v %+v\nSkipping...\n", os.Args[0], "sample_file_to_read.txt")
	}
}
