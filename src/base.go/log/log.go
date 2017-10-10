package log

import (
	"fmt"
	"math/rand"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

var DoneOKEmoji = []string{":beer:", ":hamburger:", ":pizza:", ":fries:", ":doughnut:"}

func LogDebug(message string, flagDebug bool) {
	newMessage := emoji.Sprint(":beetle: ", message)
	if flagDebug {
		fmt.Println(newMessage)
	}
}

func LogError(message string) {
	newMessage := emoji.Sprint(":heavy_exclamation_mark: ", message)
	fmt.Println(newMessage)
}

func LogGit(message string) {
	newMessage := emoji.Sprint(":sleeping: ", message)
	fmt.Println(newMessage)
}

func LogSuccess(message string) {
	newMessage := emoji.Sprint(message, "\nEnjoy! "+DoneOKEmoji[rand.Intn(len(DoneOKEmoji))])
	fmt.Println(newMessage)
}

func Log(message string) {
	newMessage := emoji.Sprint(":point_right: ", message)
	fmt.Println(newMessage)
}

func LogClean(message string) {
	newMessage := emoji.Sprint(":fire: ", message)
	fmt.Println(newMessage)
}

func writeLog(logType string, message string) {

}
