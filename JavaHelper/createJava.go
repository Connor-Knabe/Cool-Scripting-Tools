package main

import "os"
import "os/exec"
import "fmt"
import "io"

func main() {

	if (len(os.Args)<2){
		fmt.Printf("You need one argument!\n");
	} else {
		_, err := os.Open(os.Args[1]+".java")
		if err != nil {
		    // no such file or dir
			fmt.Printf("Creating File!")
			newJavaFile, err := os.Create(os.Args[1]+".java")
			n, err := io.WriteString(newJavaFile, "public class " + os.Args[1] + " {\n    public static void main(String[] args) {\n        \n        \n    }\n}")
			    if err != nil {
			        fmt.Println(n, err)
			    }
		    return
		} else {
			fmt.Printf("Compiling and running file\n")

			cmd := exec.Command("echo hi")
			cmd.Start()



		}


	}





}
