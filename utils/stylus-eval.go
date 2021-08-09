package utils

import (
	"log"
	"os"
	"os/exec"
	"rpiSite/config"
)

var (
	currentDir = os.Getenv("PWD")
)

// Go to the /home/gusted/Desktop/coding/stylus-eval.
// And execeute the command:
// ./setup.sh "/path/to/CSS" "site"
func TakeScreenshot(CSS, site, filename string) error {
	// save CSS to a tmp file
	err := os.WriteFile("./tmp.css", UnsafeByteConversion(CSS), 0644)
	if err != nil {
		return err
	}

	// Take the screenshot.
	cmd := exec.Command(config.StylusEvalDir+"/setup.sh", currentDir+"/tmp.css", site)
	err = cmd.Run()

	if err != nil {
		log.Println("failed to take screenshot:", err)
		return err
	}

	// Move the screenshot to the right place.
	err = os.Rename(config.StylusEvalDir+"/output.png", currentDir+"/static/"+filename+".png")
	if err != nil {
		log.Println("failed to move screenshot:", err)
		return err
	}

	return nil
}