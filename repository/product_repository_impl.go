package repository

import (
	"context"
	"database/sql"
	"net/http"

	cError "github.com/vincen320/product-service-graphql/helper/error"
	"github.com/vincen320/product-service-graphql/model"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
func (p *productRepository) FindAllProducts(ctx context.Context) (response []model.Product, err error) {
	rows, err := p.db.Query(
		`SELECT
			id
			, name
			, description
			, price
			, created_by
			, created_at
			, "type"
		FROM products p
		ORDER BY id `,
	)
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		err = cError.New(http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var product model.Product
		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedBy,
			&product.CreatedAt,
			&product.Type,
		); err != nil {
			err = cError.New(http.StatusInternalServerError, "internal server error", err.Error())
			return
		}
		response = append(response, product)
	}
	return
}
