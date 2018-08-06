package main

import (
	"flag"
	"os"
	"time"

	lazygit "github.com/jesseduffield/lazygit/internal/pkg/lazygit"
)

var (
	startTime time.Time
	debugging bool
)

func navigateToRepoRootDirectory() {
	_, err := os.Stat(".git")
	for os.IsNotExist(err) {
		devLog("going up a directory to find the root")
		os.Chdir("..")
		_, err = os.Stat(".git")
	}
}

func main() {
	debuggingPointer := flag.Bool("debug", false, "a boolean")
	flag.Parse()
	debugging = *debuggingPointer
	devLog("\n\n\n\n\n\n\n\n\n\n")
	startTime = time.Now()
	lazygit.verifyInGitRepo()
	navigateToRepoRootDirectory()
	lazygit.run()
}
