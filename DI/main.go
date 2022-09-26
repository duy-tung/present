package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"present/DI/database"
	"present/DI/module"
	"strconv"
)

func main() {
	// Create a new DB connection
	connString := "dbname=<your main db name> sslmode=disable"
	db, _ := sql.Open("postgres", connString)

	// Create a store dependency with the db connection
	store := database.NewStore(db)
	// Create the module by injecting the store as a dependency
	m := &module.Module{Store: store}

	// The following code implements a simple command line app to read the ID as input
	// and output the validity of the result of the entry with that ID in the database
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ID, _ := strconv.Atoi(scanner.Text())
		err := m.GetNumber(ID)
		if err != nil {
			fmt.Printf("result invalid: %v", err)
			continue
		}
		fmt.Println("result valid")
	}
}
