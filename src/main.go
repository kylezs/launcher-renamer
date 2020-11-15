package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// for prod
var appPath, errPath = os.Executable()

// for dev
// var appPath = filepath.Join("/Users/kaz/go/src/github.com/kylezs/launcher-pkg-rename-util/src/launcher-pkgs")

// OutputLocation is the folder where the launcher copies go
var OutputLocation = filepath.Join(appPath, "./../../output")

// MacNum is the number input to select a build for MacOS
const MacNum = "1"

// WindowsNum is the number input to select a build for Windows
const WindowsNum = "2"

// FormInfo holds the data entered by the user into the script
type FormInfo struct {
	email   string
	buildOs string
}

func confirm() bool {
	reader := bufio.NewReader(os.Stdin)
	cont, _ := reader.ReadString('\n')
	contClean := strings.TrimRight(cont, "\r\n")
	lowerCont := strings.ToLower(contClean)
	if lowerCont == "y" {
		return true
	}
	return false
}

func numToOS(num string) string {
	if num == MacNum {
		return "MacOS"
	} else if num == WindowsNum {
		return "Windows"
	}
	log.Fatalf("Unknown operating system number %s", num)
	return "unknown"
}

func gatherInput() FormInfo {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the email of the Local Partner you wish to send a launcher package to below")
	var emailClean string
	for {
		fmt.Printf("Enter email: ")
		email, _ := reader.ReadString('\n')
		emailClean = strings.TrimRight(email, "\r\n")
		fmt.Printf("Is this email correct: %s", emailClean)
		fmt.Printf("\nPlease enter 'y' or 'n': ")
		if confirm() {
			break
		}
		continue
	}

	fmt.Println("Please enter the number of the OS that this Local Partner is using and then hit enter:")

	var lpOsClean string
	var osName string
	for {
		fmt.Printf("%s. MacOS\n%s. Windows\nOS: ", MacNum, WindowsNum)
		lpOs, _ := reader.ReadString('\n')
		lpOsClean = strings.TrimRight(lpOs, "\r\n")
		isOsValid := lpOsClean == MacNum || lpOsClean == WindowsNum
		// if lpOsClean == WindowsNum && runtime.GOOS == "darwin" {
		// 	fmt.Println("If you are running Mac, you cannot build windows packages.")
		// 	continue
		// }
		if !isOsValid {
			fmt.Println("Please enter a valid number")
			continue
		}
		osName = numToOS(lpOsClean)
		fmt.Printf("You have selected '%s', do you wish to proceed? 'y' or 'n': ", osName)
		if confirm() {
			break
		}
	}
	return FormInfo{emailClean, osName}
}

func createOutputFolder(email string) string {
	// Create the output folder if it does not exist
	emailFolder := filepath.Join(OutputLocation, email)
	if _, err := os.Stat(emailFolder); os.IsNotExist(err) {
		// rwx for the user
		os.MkdirAll(emailFolder, 0700)
	}
	return emailFolder
}

func buildRestrictionInfo() {
	fmt.Printf("You are running this program on a %s kernel", runtime.GOOS)
	if runtime.GOOS == "darwin" {
		fmt.Println("This means you can only build packages for Mac/darwin")
	} else {
		fmt.Println("This means you can build packages for Mac or Windows")
	}
}

func main() {
	fmt.Println("To exit this script press ctrl + C")
	formInfo := gatherInput()
	fmt.Printf("You entered email: '%s', and os: '%s'\n", formInfo.email, formInfo.buildOs)

	renamePackage(&formInfo)

	fmt.Printf("You will find a '%s' folder in the output folder containing the package\n", formInfo.email)
	fmt.Println("Hit Enter to end the program")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}
