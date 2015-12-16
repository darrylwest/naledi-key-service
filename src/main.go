package main

import (
	"fmt"
	"keyservice"
    "time"
)

func main() {

	ctx := keyservice.ParseArgs()
	err := ctx.StartService()

	if err != nil {
		fmt.Println("error starting servers: ", err)
		panic(err)
	}

	fmt.Printf("KeyService Started: %v\n", ctx.ToMap())

    // TODO : remove after socket listeners are in place
    time.Sleep( 5 * time.Second )
}
