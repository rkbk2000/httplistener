package main

import (
	"context"
	"fmt"

	"github.com/rkbk2000/httplistener/rest"
)

const serverPort string = "40010"

func main() {
	fmt.Println("Starting server")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rest.StartServer(ctx, serverPort)
}
