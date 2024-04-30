package main

import (
	"context"
	"fmt"
	"github.com/pcr9x/orders-api/application"
	"os/signal"
	"os"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println(err)
	}

}