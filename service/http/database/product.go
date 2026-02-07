package database

import (
	"fmt"

	"github.com/mahima-c/fruito/model"
)

func (s *service) UpsertProducts(products []model.Product) error {
	tx, err := s.client.DB.Begin()
	if err != nil {
		fmt.Println("Failed to begin transaction", err)
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	for _, product := range products {
		err = model.UpsertProduct(tx, &product)
		if err != nil {
			fmt.Println("Failed to upsert product", err)
			return err
		}
	}

	return nil
}

func (s *service) GetProductById(id int) (*model.Product, error) {
	tx, err := s.client.DB.Begin()
	if err != nil {
		fmt.Println("Failed to begin transaction", err)
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	product, err := model.GetProductById(tx, id)
	if err != nil {
		fmt.Println("Failed to get product", err)
		return nil, err
	}

	return product, nil
}

func (s *service) GetAllProducts(limit, offset int) ([]model.Product, error) {
	tx, err := s.client.DB.Begin()
	if err != nil {
		fmt.Println("Failed to begin transaction", err)
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	products, err := model.GetAllProducts(tx, limit, offset)
	if err != nil {
		fmt.Println("Failed to get products", err)
		return nil, err
	}

	return products, nil
}
