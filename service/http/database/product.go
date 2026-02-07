package database

import (
	"fmt"

	"github.com/Mrhb787/hospital-ward-manager/model"
)

func (s *service) UpsertProduct(product model.Product) error {
	tx, err := s.client.DB.Begin()
	if err != nil {
		fmt.Println("Failed to begin transaction", err)
		return err
	}
	// add validation layer
	fmt.Println("UpsertProduct", product)

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = model.UpsertProduct(tx, &product)
	if err != nil {
		fmt.Println("Failed to upsert product", err)
		return err
	}

	return nil
}
