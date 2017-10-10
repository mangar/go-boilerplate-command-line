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
	log.LogStart("> > > > >")

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

	} else if os.Args[1] == "clear" {
		cmd.Clear()

	} else if os.Args[1] == "clearlog" {
		cmd.ClearLog()

		/**
		* @BOILERPLATE
		*
		* change the number of mandatory params needed
		 */
	} else if len(os.Args) <= 2 {

		/**
		 * @BOILERPLATE
		 *
		 * define your own error message and the basic usage message for the user
		 */
		fmt.Println("usage: go run go-boilerplate.go PARAM1 PARAM2")
		fmt.Println("Ex.:")
		fmt.Println("go run go-boilerplate.go PARAM1 PARAM2")
		fmt.Println("")

	} else {

		/**
		 * @BOILERPLATE
		 *
		 * save extra params in
		 */
		config.Param1 = os.Args[1]
		config.Param2 = os.Args[2]

		log.LogClean("Param1: " + config.Param1)
		log.LogClean("Param2: " + config.Param2)

		/**
		 * @BOILERPLATE
		 *
		 * run your program here
		 */

	}

	log.LogEnd("< < < < <")

}

func printVersion() {
	out := emoji.Sprint(":checkered_flag: Version: ", Version, "\n:clock2: Build Time: ", Build)
	fmt.Println(out)

}
