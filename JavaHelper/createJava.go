package main

import "os"
//import "os/exec"
import "fmt"


func main() {



	if (len(os.Args)<1){
		fmt.Printf("No arguments!");
	} else {
		fmt.Printf("Ok")
	}





}
