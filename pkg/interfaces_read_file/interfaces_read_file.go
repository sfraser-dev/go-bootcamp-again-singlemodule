package pkg_interfaces_read_file

import (
	"fmt"
	"log"
	"os"
)

func RunIt() {
	// os.Args hold the command-line arguments, starting with the program name
	if len(os.Args) == 2 {
		filename := os.Args[1]

		if fp, err := os.Open(filename); err != nil {
			log.Fatalln(err)
		} else {
			defer (*fp).Close() // close file via defer (out of scope or EOF)
			defer fmt.Println("closed file " + filename + " via defer")
			buffer := make([]byte, 32*1024)
			if _, err := fp.Read(buffer); err != nil {
				log.Fatalln(err)
			}
			log.Println("\n" + string(buffer))
		}
	} else {
		fmt.Printf("No input file on command line...\nUsage: go run %+v %+v\nSkipping...", os.Args[0], "sample_file_to_read.txt")
	}
}
