package main

import (
	"context"
	"fmt"
	"github.com/allnightru/go-sandbox3/internal/comment"
	db2 "github.com/allnightru/go-sandbox3/internal/db"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db2.NewDatabase()
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
	fmt.Println(cmtSevice.GetComment(
		context.Background(),
		"ebe26a6d-cce4-41e2-a889-cc3fc637bfba",
	))

	return nil
}

func main() {
	fmt.Println("Ge Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
