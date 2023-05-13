//!env/go1.20.3 (windows\amd64)
//MobCat (2023)

package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"os/exec"
	"strings"
)


func lessShitCopy(inputFile, outputFile string) {
	// less shit as I was doing this with os.Open, os.Create and io.Copy
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(outputFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", outputFile)
		fmt.Println(err)
		return
	}
}

func main() {
	fmt.Print("Enter project name without spaces: ")
	var projectName string
	var initCMD string
	runCMD := false
	// Scan excapes on spaces. lol ok.
	fmt.Scan(&projectName)

	fmt.Printf("Building new project \"%s\"\n", projectName)

	// Make new folder
	os.Mkdir(projectName, 0755)

	// Copy and rename template into new folder
	lessShitCopy("init.go", fmt.Sprintf("%s/main.go", projectName))
	

	fmt.Print("Run init cmds? [Y/N]: ")
	fmt.Scan(&initCMD)
	initCMD = strings.ToLower(initCMD)
	if initCMD == "y" {
		lessShitCopy("init.cmd", fmt.Sprintf("%s/init.cmd", projectName))
		runCMD = true
	}
	

	// cd into new folder
	os.Chdir(projectName)

	// run the init project thingy
	cmd := exec.Command("go", "mod", "init", projectName)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}

	// Build our build batch for compiling project
	f, _ := os.Create("build.bat")
	defer f.Close()
	
	f.WriteString(fmt.Sprintf("go build -trimpath -buildmode=exe -ldflags=\"-s -w\" -o %s.exe main.go\n", projectName))
	f.WriteString(fmt.Sprintf("upx %s.exe", projectName))

	// if runCMD == true run the init cmds and delete the cmd when done.
	if runCMD {
		fmt.Println("Running init cmds...")
		cmd = exec.Command("cmd.exe", "/C", "init.cmd")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Printf("cmd.Run() failed with %s\n", err)
		}
	
		os.Remove("init.cmd")
	}

	fmt.Printf("\nProject \"%s\" has been built and is ready to use \\^__^/\n", projectName)

}