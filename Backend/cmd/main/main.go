package main

import (
	"fmt"
	"log"

	"github.com/janharkonen/talousanalyysi/internal/database/dbconnect"
)

//func main() {
//	router := routes.NewRouter()
//
//	port := 8080
//	addr := fmt.Sprintf(":%d", port)
//	fmt.Printf("Server listening on http://localhost%s\n", addr)
//	err := http.ListenAndServe(addr, router)
//	if err != nil {
//		panic(err)
//	}
//}

func main() {
	db, err := dbconnect.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	//var str string = databaseinterface.RunQuery(db)
	//fmt.Println(str)
	fmt.Println("Successfully connected to the database!")
}
