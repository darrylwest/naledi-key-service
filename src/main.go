package main

import (
	"fmt"
	"keyservice"
)

func main() {

	ctx := keyservice.ParseArgs()
	// err := keyservice.StartService()

	fmt.Printf("KeyService Started: %v\n", ctx.ToMap())
}
