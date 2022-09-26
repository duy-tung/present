package module

import (
	"fmt"
	"present/DI/database"
)

type Module struct {
	Store database.Store
}

func (m *Module) GetNumber(ID int) error {
	// Use the `Get` method of the dependency to retreive the value of the database entry
	result, err := m.Store.Get(ID)
	if err != nil {
		return err
	}
	// Perform some validation, and output an error if it is too high
	if result > 10 {
		return fmt.Errorf("result too high: %d", result)
	}
	// Return nil, if the result is valid
	return nil
}
