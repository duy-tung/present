package database

import "database/sql"

type Store interface {
	// Get will fetch the value (which is an integer) for a given ID
	Get(ID int) (int, error)
}

type store struct {
	db *sql.DB
}

// Implement the "Get" method, in order to comply with the "Store" interface
func (d *store) Get(ID int) (int, error) {
	//we would perform some external database operation with d.db
	// for the sake of clarity, that code is not shown here
	return 0, nil
}

// Add a constructor function to return a new instance of a store
func NewStore(db *sql.DB) Store {
	return &store{db}
}
