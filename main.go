package main

import (
	"log"

	"github.com/kanianursawitri/student_marks/config"
	"github.com/kanianursawitri/student_marks/db"
	"github.com/kanianursawitri/student_marks/server"
)

func main() {
	config.LoadConfig(".env")

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(80)

	s := server.NewServer(db)
	s.RegisterRoute()

	log.Fatal(s.Run())
}
