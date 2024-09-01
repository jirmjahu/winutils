package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This program can only be used on Windows.")
		return
	}

	bannerText := `
 	 _      ___      __  ____  _ __  
 	| | /| / (_)__  / / / / /_(_) /__
 	| |/ |/ / / _ \/ /_/ / __/ / (_-<
 	|__/|__/_/_//_/\____/\__/_/_/___/
	`

	initializeApp(bannerText)
}
