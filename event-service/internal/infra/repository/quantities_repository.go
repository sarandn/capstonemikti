package repository

import (
	"database/sql"
	"event-service/internal/domain/model"
	"log"
)

type QuantitiesRepository struct {
	db *sql.DB
}

func NewQuantitiesRepository(db *sql.DB) *QuantitiesRepository {
	return &QuantitiesRepository{db: db}
}

func (r *QuantitiesRepository) Create(quantity *model.Quantities) error {
	query := `INSERT INTO quantities (purchase_quantity) VALUES ($1) RETURNING quantity_id`
	err := r.db.QueryRow(query, quantity.PurchaseQuantity).Scan(&quantity.QuantityID)
	if err != nil {
		log.Printf("Failed to create quantity: %v", err)
		return err
	}
	return nil
}

func (r *QuantitiesRepository) GetByID(quantityID int) (*model.Quantities, error) {
	query := `SELECT quantity_id, purchase_quantity FROM quantities WHERE quantity_id = $1`
	row := r.db.QueryRow(query, quantityID)
	var quantity model.Quantities
	err := row.Scan(&quantity.QuantityID, &quantity.PurchaseQuantity)
	if err != nil {
		log.Printf("Failed to get quantity: %v", err)
		return nil, err
	}
	return &quantity, nil
}

func (r *QuantitiesRepository) GetAll() ([]*model.Quantities, error) {
	query := `SELECT quantity_id, purchase_quantity FROM quantities`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Failed to get quantities: %v", err)
		return nil, err
	}
	defer rows.Close()

	var quantities []*model.Quantities
	for rows.Next() {
		var quantity model.Quantities
		if err := rows.Scan(&quantity.QuantityID, &quantity.PurchaseQuantity); err != nil {
			log.Printf("Failed to scan quantity: %v", err)
			return nil, err
		}
		quantities = append(quantities, &quantity)
	}
	return quantities, nil
}

func (r *QuantitiesRepository) Update(quantity *model.Quantities) error {
	query := `UPDATE quantities SET purchase_quantity = $1 WHERE quantity_id = $2`
	_, err := r.db.Exec(query, quantity.PurchaseQuantity, quantity.QuantityID)
	if err != nil {
		log.Printf("Failed to update quantity: %v", err)
		return err
	}
	return nil
}

func (r *QuantitiesRepository) Delete(quantityID int) error {
	query := `DELETE FROM quantities WHERE quantity_id = $1`
	_, err := r.db.Exec(query, quantityID)
	if err != nil {
		log.Printf("Failed to delete quantity: %v", err)
		return err
	}
	return nil
}
