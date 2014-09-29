import "os"
import "exec"
import "fmt"

fName := "";
fName = os.Args[1]

if _, err := os.Stat(fName); err == nil {
	fmt.Println(file, "exist!")

	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()

}
