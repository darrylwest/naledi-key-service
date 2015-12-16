package main

import (
	"fmt"
	"keyservice"
)

func main() {

	ctx := keyservice.ParseArgs()
	err := ctx.StartService()

	if err != nil {
		fmt.Println("error starting servers: ", err)
		panic( err )
	}

	fmt.Printf("KeyService Started: %v\n", ctx.ToMap())
}
