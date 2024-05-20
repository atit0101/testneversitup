package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-shop/config"
	"github.com/go-shop/pkg/database"
	"github.com/go-shop/server"
)

func main() {
	ctx := context.Background()
	// _ = ctx

	fmt.Println("dasd")

	// Iiitialize ronfig
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error .env path required")
		}
		log.Println(os.Args)
		return os.Args[1]
	}())

	// Database connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, &cfg, db)

}
