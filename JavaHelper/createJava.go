package main

import "os"
//import "os/exec"
import "fmt"
import "io"

func main() {

	if (len(os.Args)<2){
		fmt.Printf("You need one argument!");
	} else {
		_, err := os.Open(os.Args[1]+".java")
		if err != nil {
		    // no such file or dir
			fmt.Printf("File not found")
			w, err := os.Create(os.Args[1])
			n, err := io.WriteString(w, "blahblahblah")
			    if err != nil {
			        fmt.Println(n, err)
			    }
		    return
		} else {

			fmt.Printf("HI")
		}


	}





}
