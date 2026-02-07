package model

import (
	"database/sql"
	"encoding/json"
)

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	Price         int    `json:"price"`
	UnitOfMeasure string `json:"unit_of_measure"`
	TotalQty      int    `json:"total_qty"`
	Description   string `json:"description"` // Stored as JSONB in DB
	Rating        int    `json:"rating"`
	RatingCount   int    `json:"rating_count"`
	Tag           string `json:"tag"`
}

func UpsertProduct(tx *sql.Tx, p *Product) error {
	// Verify description is valid JSON
	if p.Description == "" {
		p.Description = "{}"
	} else {
		if !json.Valid([]byte(p.Description)) {
			// fallback assuming valid json or empty object
		}
	}

	// 1. If ID is 0, just Insert (Auto-increment ID)
	if p.ID == 0 {
		insertQuery := `
			INSERT INTO product (name, image, price, unit_of_measure, total_qty, description, rating, rating_count, tag)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`
		_, err := tx.Exec(insertQuery,
			p.Name, p.Image, p.Price, p.UnitOfMeasure, p.TotalQty, p.Description, p.Rating, p.RatingCount, p.Tag,
		)
		return err
	}

	// 2. Try to Update using ID if provided
	updateQuery := `
		UPDATE product SET
			name = $2,
			image = $3,
			price = $4,
			unit_of_measure = $5,
			total_qty = $6,
			description = $7,
			rating = $8,
			rating_count = $9,
			tag = $10
		WHERE id = $1
	`
	res, err := tx.Exec(updateQuery,
		p.ID, p.Name, p.Image, p.Price, p.UnitOfMeasure, p.TotalQty, p.Description, p.Rating, p.RatingCount, p.Tag,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	}

	// 3. If Update affected 0 rows (Explicit ID not found), Insert with that ID
	// We use OVERRIDING SYSTEM VALUE because the table 'id' column is likely an IDENTITY column.
	insertQuery := `
		INSERT INTO product (id, name, image, price, unit_of_measure, total_qty, description, rating, rating_count, tag)
		OVERRIDING SYSTEM VALUE
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err = tx.Exec(insertQuery,
		p.ID, p.Name, p.Image, p.Price, p.UnitOfMeasure, p.TotalQty, p.Description, p.Rating, p.RatingCount, p.Tag,
	)
	return err
}
