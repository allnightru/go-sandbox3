package main

import (
	"fmt"
	"github.com/allnightru/go-sandbox3/internal/comment"
	"github.com/allnightru/go-sandbox3/internal/db"
	transportHttp "github.com/allnightru/go-sandbox3/internal/transport/http"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected and migrated database")

	cmtSevice := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtSevice)
	if err := httpHandler.Serve(); err != nil {

	}

	return nil
}

func main() {
	fmt.Println("Ge Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
