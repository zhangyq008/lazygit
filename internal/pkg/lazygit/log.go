package lazygit

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/fatih/color"
)

func homeDirectory() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func devLog(objects ...interface{}) {
	localLog(color.FgWhite, homeDirectory()+"/go/src/github.com/jesseduffield/lazygit/development.log", objects...)
}

func colorLog(colour color.Attribute, objects ...interface{}) {
	localLog(colour, homeDirectory()+"/go/src/github.com/jesseduffield/lazygit/development.log", objects...)
}

func commandLog(objects ...interface{}) {
	localLog(color.FgWhite, homeDirectory()+"/go/src/github.com/jesseduffield/lazygit/commands.log", objects...)
}

func localLog(colour color.Attribute, path string, objects ...interface{}) {
	// if !debugging {
	// 	return
	// }
	return
	f, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	for _, object := range objects {
		colorFunction := color.New(colour).SprintFunc()
		f.WriteString(colorFunction(fmt.Sprint(object)) + "\n")
	}

}
