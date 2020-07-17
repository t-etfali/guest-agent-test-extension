package main

import (
	"fmt"
	"os"
	"regexp"
)

func install() {
	fmt.Println("Installed Successfully.")
}

func enable() {
	fmt.Println("Enabled Successfully.")
}

func disable() {
	fmt.Println("Disabled Successfully.")
}

func uninstall() {
	fmt.Println("Uninstalled Successfully.")
}

func update() {
	fmt.Println("Updated Successfully.")
}

func main() {
	if len(os.Args[1:]) > 0 {
		for _, a := range os.Args[1:] {
			/* 	TODO : Not sure if there is a better method in regexp so don't need multiple vars

			TODO : Since there are only the 5 commands that should be called, this could be changed
			to just check for os.Args[1] and compare equality (ignore case)
			*/
			matchDisable, _ := regexp.MatchString("^([-/]*)(disable)", a)
			matchUninstall, _ := regexp.MatchString("^([-/]*)(uninstall)", a)
			matchInstall, _ := regexp.MatchString("^([-/]*)(install)", a)
			matchEnable, _ := regexp.MatchString("^([-/]*)(enable)", a)
			matchUpdate, _ := regexp.MatchString("^([-/]*)(update)", a)

			if matchDisable {
				disable()
			} else if matchUninstall {
				uninstall()
			} else if matchInstall {
				install()
			} else if matchEnable {
				enable()
			} else if matchUpdate {
				update()
			} else {
				fmt.Println("Command Not Recognized.")
			}
		}
	} else {
		fmt.Println("No command line arguments provided")
	}
	/* 	TODO : Error handling might be necessary for if there is no match, but this could
	just be a print statement else case if the regexp doesn't raise panics/errors
	*/
}
