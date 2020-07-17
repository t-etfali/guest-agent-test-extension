package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/Azure/azure-extension-foundation/sequence"
	"github.com/Azure/azure-extension-foundation/settings"
	"github.com/Azure/azure-extension-foundation/status"
	"github.com/go-kit/kit/log"
)

func install() {

	operation := "install"
	msg := "Installed Successfully"

	// Configuring logger to print time and verison by default
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "time", log.DefaultTimestamp, "version", "1.0.0.0")

	logger.Log("event", "install", "message", "Installed Successfully.")

	// Get the MrSeq from environment and extension (this is not implemented on windows, just returns -1 -1)
	extensionMrseq, environmentMrseq, _ := sequence.GetMostRecentSequenceNumber()
	logger.Log("message", fmt.Sprintf("extensionMrSeq: %v environmentMrSeq: %v", extensionMrseq, environmentMrseq))

	// Just drops the status in the current file location
	_ = status.ReportSuccess(extensionMrseq, operation, msg)

	// Grab the handler environment struct
	he, err1 := settings.GetHandlerEnvironment()
	if err1 != nil {
		logger.Log("event", "Retrieve Handler Environment Failed")
	}

	// access the version field (testing to make sure handler environment can be accessed)
	logger.Log("handler version", he.Version)

	// TODO Figure out how to access the public settings
}

func enable() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`

	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := result["birds"].(map[string]interface{})

	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}

	fmt.Println("Enabled Successfully.")

	_ = status.ReportSuccess(extensionMrseq, operation, msg)
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
			/*	TODO : Not sure if there is a better method in regexp so don't need multiple vars
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
