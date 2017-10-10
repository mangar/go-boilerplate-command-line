package cmd

import "fmt"

/**
 * @BOILERPLATE
 *
 * help parameter from command line.
 *
 */
func Help() {

	fmt.Println("")
	fmt.Println("usage: go run boilerplate.go PARAM1 PARAM2")
	fmt.Println("")
	fmt.Println("Ex.:")
	fmt.Println("go run boilerplate.go PARAM1 PARAM2")
	fmt.Println("")
	fmt.Println("options:")
	// fmt.Println(" clear    perform the cleaning on temporary dir. The temporary dir is located at ~/.gonfig/ ")
	fmt.Println(" help       show this content")
	fmt.Println(" install    create the basic structure for configruation")
	fmt.Println(" uninstall  remove the the basic structure for configruation")
	// fmt.Println(" config     display the configuration loaded from config file")
	// fmt.Println(" version    display the version and build time")
	fmt.Println("")

}
