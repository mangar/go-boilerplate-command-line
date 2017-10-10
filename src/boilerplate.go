package main

import (
	"fmt"
	"os"

	"base.go/cmd"
	"base.go/config"
	"base.go/log"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

var (
	Version string
	Build   string
	Hash    string
)

func main() {

	if os.Getenv("FlagDebug") == "true" {
		config.FlagDebug = true
		fmt.Println("Debug mode enabled")
	}

	if os.Args[1] == "install" {
		cmd.Install()

	} else if os.Args[1] == "uninstall" {
		cmd.Uninstall()

	} else if os.Args[1] == "help" {
		cmd.Help()

	} else if os.Args[1] == "config" {
		cmd.ShowConfig()

	} else if os.Args[1] == "version" {
		printVersion()

	} else if len(os.Args) < 4 {
		fmt.Println("usage: go run main.go SERVICE_PROVIDER INPUT_DIR OUTPUT_DIR")
		fmt.Println("Ex.:")
		fmt.Println("go run main.go [uber|99] ../data/input/uber ../.data/output")
		fmt.Println("")

	} else {

		log.LogDebug("running!", config.FlagDebug)

		// general.ServiceProvider = os.Args[1]
		// general.InputDir = os.Args[2]
		// general.OutputDir = os.Args[3]

		// log.LogDebug("Service Provider Selected: "+general.ServiceProvider, config.FlagDebug)
		// log.LogDebug("InputDir: "+general.InputDir, config.FlagDebug)

		// if general.ServiceProvider == general.ServiceProviderUber {
		// 	uberRecords := uber.Read()
		// 	output.WriteUber(uberRecords)

		// } else if general.ServiceProvider == general.ServiceProvider99 {
		// 	taxi99Records := taxi99.Read()
		// 	output.WriteTaxi99(taxi99Records)

		// } else {
		// 	log.LogError("hum.... looks that the service provider informed is not quite right. check if it's uber or 99")
		// }

	}

}

func printVersion() {
	out := emoji.Sprint(":checkered_flag: Version: ", Version, "\n:clock2: Build Time: ", Build)
	fmt.Println(out)

}
