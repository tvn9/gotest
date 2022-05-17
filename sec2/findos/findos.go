package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	var runtime string
	switch os {
	case "darwin":
		runtime = "darwin"
	case "linux":
		runtime = "linux"
	case "windows":
		runtime = "windows"
	default:
		runtime = "I am not sure about your system runtime! :("
	}
	fmt.Printf("Your system's runtime OS is %q\n", runtime)
}
