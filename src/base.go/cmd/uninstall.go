package cmd

import (
	"os/exec"

	"base.go/config"
	"base.go/log"
)

/**
 *
 * Delete the ~/PROJECT_DIR
 *
 */
func Uninstall() {

	tempDir := config.ConfigDir

	if err := exec.Command("sh", "-c", "rm -rf "+tempDir).Run(); err != nil {
		log.LogError("Error removing project home directory")
	}

	log.LogSuccess("Uninstall complete!")

}
