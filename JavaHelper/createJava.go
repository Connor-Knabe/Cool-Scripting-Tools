package main

import "os"
//import "os/exec"
import "fmt"


func main() {

	if (len(os.Args)<2){
		fmt.Printf("You need one argument!");
	} else {
		fmt.Printf("Ok")
	}





}
