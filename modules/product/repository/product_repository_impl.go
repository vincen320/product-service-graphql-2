package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	cError "github.com/vincen320/product-service-graphql-2/helper/error"
	productModel "github.com/vincen320/product-service-graphql-2/modules/product/model"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
func (p *productRepository) FindAllProducts(ctx context.Context) (response []productModel.Product, err error) {
	rows, err := p.db.Query(
		`SELECT
			p.id
			, p.name
			, p.description
			, p.price
			, p.created_by
			, p.created_at
			, p."type"
			, pa.attribute
			, pa.value
		FROM products p
		JOIN product_attributes pa ON pa.product_id = p.id
		ORDER BY id`,
	)
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		err = cError.New(http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	defer rows.Close()
	mapProduct := map[int64]productModel.Product{}
	for rows.Next() {
		product := productModel.Product{
			AdditionalAttr: map[string]interface{}{},
		}
		var attribute, value string
		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedBy,
			&product.CreatedAt,
			&product.Type,
			&attribute,
			&value,
		); err != nil {
			err = cError.New(http.StatusInternalServerError, "internal server error", err.Error())
			return
		}
		if existingProduct, ok := mapProduct[product.ID]; !ok {
			product.AdditionalAttr[attribute] = value
			response = append(response, product)
			mapProduct[product.ID] = product
		} else {
			existingProduct.AdditionalAttr[attribute] = value
			lastIndex := len(response) - 1
			response[lastIndex] = existingProduct
			mapProduct[product.ID] = existingProduct
		}
	}
	return
}

func (p *productRepository) CreateProduct(ctx context.Context, request productModel.Product) (response productModel.Product, err error) {
	request.CreatedAt = time.Now().UTC()
	tx, err := p.db.Begin()
	if err != nil {
		err = cError.New(http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	err = tx.QueryRow(
		`INSERT INTO products(
			 name
			, description
			, price
			, created_by
			, created_at
			, "type"
		) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		request.Name,
		request.Description,
		request.Price,
		request.CreatedBy,
		request.CreatedAt,
		request.Type,
	).Scan(&request.ID)
	if err != nil {
		tx.Rollback()
		err = cError.New(http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	var placeholders []string
	params := []interface{}{request.ID}
	for k, v := range request.AdditionalAttr {
		if reflect.ValueOf(v).Kind() == reflect.Ptr {
			v = reflect.Indirect(reflect.ValueOf(v)).Interface()
			request.AdditionalAttr[k] = v
		}
		params = append(params, k, v)
		placeholders = append(placeholders, fmt.Sprintf("($1, $%d, $%d)", len(params)-1, len(params)))
	}
	if len(placeholders) > 0 {
		if _, err = tx.Exec(
			fmt.Sprintf(
				`INSERT INTO product_attributes(
				product_id
				, attribute
				, value
			) VALUES %s`,
				strings.Join(placeholders, ","),
			),
			params...,
		); err != nil {
			tx.Rollback()
			err = cError.New(http.StatusInternalServerError, "internal server error", err.Error())
			return
		}
	}
	tx.Commit()
	return request, err
}
