package main

import "os"
import "os/exec"
import "fmt"
import "io"

func main() {

	if (len(os.Args)<2){
		fmt.Printf("You need one argument!\n");
	} else {
		fName := os.Args[1]+".java"
		_, err := os.Open(fName)
		if err != nil {
		    // no such file or dir
			fmt.Printf("Creating File!")
			newJavaFile, err := os.Create(fName)
			n, err := io.WriteString(newJavaFile, "public class " + os.Args[1] + " {\n    public static void main(String[] args) {\n        \n        \n    }\n}")
			    if err != nil {
			        fmt.Println(n, err)
			    }
		    return
		} else {
			fmt.Printf("Compiling and running file\n")

			cmd := exec.Command("javac", fName)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Start()
			fmt.Printf(os.Args[1])
			cmd = exec.Command("java", os.Args[1])
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Start()


		}


	}





}
