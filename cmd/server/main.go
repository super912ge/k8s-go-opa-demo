package main

import (
	"database/sql"
	"fmt"
	"k8s-go-opa/pkg/api"
	"k8s-go-opa/pkg/app"
	"k8s-go-opa/pkg/repository"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

// func run will be responsible for setting up db connections, routers etc
func run() error {
	connectionString := "postgres://postgres:1234@localhost/k8s_go_opa?sslmode=disable"

	// setup database connection
	db, err := setupDatabase(connectionString)

	if err != nil {
		return err
	}

	// create storage dependency
	storage := repository.NewStorage(db)

	// create router dependecy
	router := gin.Default()
	router.Use(cors.Default())

	// create user service
	userService := api.NewUserService(storage)

	// create team service
	teamService := api.NewTeamService(storage)

	server := app.NewServer(router, userService, teamService)

	err = storage.Migrate(connectionString)

	if err != nil {
		return err
	}

	// start the server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
