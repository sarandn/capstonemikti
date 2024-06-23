package repository

import (
	"database/sql"
	"event-service/internal/domain/model"
	"log"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	query := `INSERT INTO category (category_name) VALUES ($1) RETURNING category_id`
	err := r.db.QueryRow(query, category.CategoryName).Scan(&category.CategoryID)
	if err != nil {
		log.Printf("Failed to create category: %v", err)
		return err
	}
	return nil
}

func (r *CategoryRepository) GetByID(categoryID int) (*model.Category, error) {
	query := `SELECT category_id, category_name FROM category WHERE category_id = $1`
	row := r.db.QueryRow(query, categoryID)
	var category model.Category
	err := row.Scan(&category.CategoryID, &category.CategoryName)
	if err != nil {
		log.Printf("Failed to get category: %v", err)
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetAll() ([]*model.Category, error) {
	query := `SELECT category_id, category_name FROM category`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Failed to get categories: %v", err)
		return nil, err
	}
	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
			log.Printf("Failed to scan category: %v", err)
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (r *CategoryRepository) Update(category *model.Category) error {
	query := `UPDATE category SET category_name = $1 WHERE category_id = $2`
	_, err := r.db.Exec(query, category.CategoryName, category.CategoryID)
	if err != nil {
		log.Printf("Failed to update category: %v", err)
		return err
	}
	return nil
}

func (r *CategoryRepository) Delete(categoryID int) error {
	query := `DELETE FROM category WHERE category_id = $1`
	_, err := r.db.Exec(query, categoryID)
	if err != nil {
		log.Printf("Failed to delete category: %v", err)
		return err
	}
	return nil
}
