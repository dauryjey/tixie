package main

import (
	"auth/config"
	"fmt"
)

func main() {
	config.InitGlobalEnv()

	fmt.Println("Hello auth!")
}
