package cmd

import (
	"os"
	"os/exec"

	"base.go/config"
	"base.go/log"
)

/**
 *
 * Delete the ~/PROJECT_DIR/temp
 *
 */
func Clear() {

	// delete
	if err := exec.Command("sh", "-c", "rm -rf "+config.TempDir).Run(); err != nil {
		log.LogError("Error removing temp home directory")
	}

	// re-create
	_, err := os.Stat(config.TempDir)
	if os.IsNotExist(err) {
		log.LogDebug("Creating temp dir @"+config.TempDir, config.FlagDebug)
		os.MkdirAll(config.TempDir, os.ModePerm)
	}

	log.LogSuccess("Uninstall complete!")
}
