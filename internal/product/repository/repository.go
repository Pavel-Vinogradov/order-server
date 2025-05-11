package repository

import (
	"database/sql"
	"github.com/lib/pq"
	"order-server/internal/product/entity"
)

type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) Create(product entity.Product) (entity.Product, error) {
	query := `INSERT INTO products (name, description, images) 
              VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, product.Name, product.Description, pq.StringArray(product.Images)).
		Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
	return product, err

}

func (r *ProductRepository) Update(p entity.Product) (entity.Product, error) {
	query := `UPDATE products SET name=$1, description=$2, images=$3, updated_at=NOW() 
              WHERE id=$4 RETURNING updated_at`
	err := r.db.QueryRow(query, p.Name, p.Description, pq.StringArray(p.Images), p.ID).
		Scan(&p.UpdatedAt)
	return p, err
}
func (r *ProductRepository) FindByID(id int64) (entity.Product, error) {
	var p entity.Product
	query := `SELECT id, name, description, images, created_at, updated_at FROM products WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, pq.Array(&p.Images), &p.CreatedAt, &p.UpdatedAt)
	return p, err
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
func (r *ProductRepository) Delete(p entity.Product) (bool, error) {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", p.ID)
	return err == nil, err
}

func (r *ProductRepository) FindAll() ([]entity.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, images, created_at, updated_at FROM products")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	products := make([]entity.Product, 0)

	for rows.Next() {
		var p entity.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Description, pq.Array(&p.Images), &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
