package product

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/LucasRosello/poc-dell/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Get(id int) (domain.Product, error)
	Save(w domain.Product) (int, error)
	Update(w domain.Product) error
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Hello() string {
	var greeting string
	err := r.db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return greeting
}

func (r *repository) GetAll() ([]domain.Product, error) {
	query := "SELECT * FROM product;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		p := domain.Product{}
		_ = rows.Scan(&p.ProductCode, &p.Description, &p.Avaliable)
		products = append(products, p)
	}

	return products, nil
}

func (r *repository) Get(id int) (domain.Product, error) {
	query := "SELECT * FROM products WHERE id=?;"
	row := r.db.QueryRow(query, id)
	p := domain.Product{}
	err := row.Scan(&p.ProductCode, &p.Description, &p.Avaliable)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) ExistsProduct(p domain.Product) bool {
	query := "SELECT id FROM products WHERE id=?;"
	row := r.db.QueryRow(query, p.ProductCode)
	err := row.Scan(&p.ProductCode, &p.Description, &p.Avaliable)
	return err == nil
}

func (r *repository) Save(p domain.Product) (int, error) {
	query := "INSERT INTO products(description,expiration_rate,freezing_rate,height,lenght,netweight,product_code,recommended_freezing_temperature,width,id_product_type,id_seller) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.ProductCode)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(p domain.Product) error {
	query := "UPDATE products SET name=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.ProductCode)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(id int) error {
	query := "DELETE FROM products WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return errors.New("Product not found")
	}

	return nil
}
