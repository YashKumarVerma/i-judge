package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	ui "github.com/YashKumarVerma/go-lib-ui"
)

func main() {
	inputString := "5\napple\n"
	ui.ContextPrint("balance_scale", "I Judge")
	ui.ContextPrint("moai", "Input :")
	fmt.Println(inputString)

	// some mojo here
	rescueStdout := os.Stdout
	read, write, _ := os.Pipe()
	os.Stdout = write

	//Just for testing, replace with your subProcess
	subProcess := exec.Command("/home/yash/Desktop/files/works/projects/i-judge/codes/question1/one.out")
	stdin, err := subProcess.StdinPipe()
	ui.CheckError(err, "Error creating stdin pipe", true)
	defer stdin.Close()

	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	err = subProcess.Start()
	ui.CheckError(err, "Error executing code", true)

	io.WriteString(stdin, inputString)

	// close pipes
	write.Close()
	out, _ := ioutil.ReadAll(read)
	os.Stdout = rescueStdout

	// final display
	ui.ContextPrint("moai", "Output")
	fmt.Println(string(out))
}
