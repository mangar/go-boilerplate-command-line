package log

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"base.go/config"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

var DoneOKEmoji = []string{":beer:", ":hamburger:", ":pizza:", ":fries:", ":doughnut:"}

func LogDebug(message string, flagDebug bool) {
	newMessage := emoji.Sprint(":beetle: ", message)
	writeLog("debug", message)
	if flagDebug {
		fmt.Println(newMessage)
	}
}

func LogError(message string) {
	newMessage := emoji.Sprint(":heavy_exclamation_mark: ", message)
	writeLog("error", message)
	fmt.Println(newMessage)
}

func LogGit(message string) {
	newMessage := emoji.Sprint(":sleeping: ", message)
	writeLog("git", message)
	fmt.Println(newMessage)
}

func LogSuccess(message string) {
	newMessage := emoji.Sprint(message, "\nEnjoy! "+DoneOKEmoji[rand.Intn(len(DoneOKEmoji))])
	writeLog("success", message)
	fmt.Println(newMessage)
}

func Log(message string) {
	newMessage := emoji.Sprint(":point_right: ", message)
	writeLog("log", message)
	fmt.Println(newMessage)
}

func LogClean(message string) {
	newMessage := emoji.Sprint(":fire: ", message)
	writeLog("clean", message)
	fmt.Println(newMessage)
}

func LogStart(message string) {
	writeLog("start", message)
}

func LogEnd(message string) {
	writeLog("end", message)
}

/**
 * @BOILERPLATE
 *
 * change the log file pattern (go-boilerplate_060201.log) to one that fits your needs
 */
func writeLog(logType string, message string) {
	if config.WriteLog {

		now := time.Now()
		fileName := config.GetConfig().Log + "/" + now.Format("go-boilerplate_060201.log")
		// fmt.Println("FileName:", fileName)

		f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		defer f.Close()
		f.WriteString(now.Format("2006-02-01T03:04:05") + "|" + strings.ToUpper(logType) + "|" + config.ProcessID + "|" + message + "\n")
	}
}

/**
 * @BOILERPLATE
 *
 * change the log file pattern (go-boilerplate_060201.log) to one that fits your needs
 */
func LogFiles() []string {

	retFiles := []string{}

	files, _ := ioutil.ReadDir(config.GetConfig().Log)
	for _, f := range files {
		if strings.Index(f.Name(), "go-boilerplate_") > -1 {
			retFiles = append(retFiles, f.Name())
		}
	}

	return retFiles
}
