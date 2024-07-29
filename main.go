package main

import (
	"fmt"
	"runtime"
)

func getInformation() {
	fmt.Println("Operating System", runtime.GOOS)
	fmt.Println("Architecture", runtime.GOARCH)
}

func main() {
	getInformation()
	fmt.Println("Hello World!")
}
