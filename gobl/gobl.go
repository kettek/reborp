package main

import (
	"runtime"

	. "github.com/kettek/gobl" // Yes, I should use dot imports.
)

func main() {
	var exe string
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}

	runArgs := append([]interface{}{}, "./reborp"+exe)

	Task("build").
		Exec("go", "build", "../")
	Task("run").
		Exec(runArgs...)
	Task("watch").
		Watch("../**/*.go", "../data/**/*").
		Signaler(SigQuit).
		Run("build").
		Run("run")
	Go()
}
