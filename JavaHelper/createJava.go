package main

import "os"
//import "os/exec"
import "fmt"


func main() {

    fName := "";
    fName = os.Args[1]

    if _, err := os.Stat(fName); err == nil {
	 
		fmt.Printf("HIII")

}


}
