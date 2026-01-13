package repository

import (
	"database/sql"
	"errors"
	"pencatatan/internal/models"

	"github.com/google/uuid"
)

type SaleRepository interface {
	Create(sale *models.CreateSalesRequest) (*models.Sale, error)
	GetByID(id uuid.UUID) (*models.Sale, error)
	GetAll() ([]*models.Sale, error)
	Update(id uuid.UUID, sale *models.UpdateSaleRequest) (*models.Sale, error)
	Delete(id uuid.UUID) error
}

type saleRepository struct {
	db *sql.DB
}

func NewSaleRepository(db *sql.DB) SaleRepository {
	return &saleRepository{
		db: db,
	}
}

func (r *saleRepository) Create(saleReq *models.CreateSalesRequest) (*models.Sale, error) {
	query := `INSERT INTO sales (name, product, quantity, price, amount_received, is_debt) 
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id, name, product, quantity, price, total, amount_received, change_amount, transaction_date, is_debt,created_at, updated_at`

	var sale models.Sale
	err := r.db.QueryRow(
		query,
		saleReq.Name,
		saleReq.Product,
		saleReq.Quantity,
		saleReq.Price,
		saleReq.AmountReceived,
		saleReq.IsDebt,
	).Scan(
		&sale.ID,
		&sale.Name,
		&sale.Product,
		&sale.Quantity,
		&sale.Price,
		&sale.Total,
		&sale.AmountReceived,
		&sale.ChangeAmount,
		&sale.TransactionDate,
		&sale.IsDebt,
		&sale.CreatedAt,
		&sale.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &sale, nil
}

func (r *saleRepository) GetByID(id uuid.UUID) (*models.Sale, error) {
	query := `SELECT id, name, product, quantity, price, total, amount_received, change_amount, transaction_date, is_debt,created_at, updated_at
				FROM sales WHERE id = $1`

	var sale models.Sale
	err := r.db.QueryRow(query, id).Scan(
		&sale.ID,
		&sale.Name,
		&sale.Product,
		&sale.Quantity,
		&sale.Price,
		&sale.Total,
		&sale.AmountReceived,
		&sale.ChangeAmount,
		&sale.TransactionDate,
		&sale.IsDebt,
		&sale.CreatedAt,
		&sale.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &sale, nil
}

func (r *saleRepository) GetAll() ([]*models.Sale, error) {
	query := `SELECT id, name,product, quantity, price, total, amount_received, change_amount, transaction_date, is_debt,created_at, updated_at
				FROM sales ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []*models.Sale
	for rows.Next() {
		var sale models.Sale
		err = rows.Scan(
			&sale.ID,
			&sale.Name,
			&sale.Product,
			&sale.Quantity,
			&sale.Price,
			&sale.Total,
			&sale.AmountReceived,
			&sale.ChangeAmount,
			&sale.TransactionDate,
			&sale.IsDebt,
			&sale.CreatedAt,
			&sale.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		sales = append(sales, &sale)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sales, nil
}

func (r *saleRepository) Update(id uuid.UUID, saleReq *models.UpdateSaleRequest) (*models.Sale, error) {
	query := `
        UPDATE sales
        SET name = COALESCE(NULLIF($1, ''), name),
            product = COALESCE(NULLIF($2, ''), product),
            quantity = COALESCE(NULLIF($3, 0), quantity),
            price = COALESCE(NULLIF($4, 0), price),
            amount_received = COALESCE(NULLIF($5, 0), amount_received),
            is_debt = $6,
            updated_at = NOW()
        WHERE id = $7
        RETURNING id, name,product, quantity, price, total, amount_received,
                  change_amount, transaction_date, is_debt,created_at, updated_at
    `

	var sale models.Sale
	err := r.db.QueryRow(
		query,
		saleReq.Name,
		saleReq.Product,
		saleReq.Quantity,
		saleReq.Price,
		saleReq.AmountReceived,
		saleReq.IsDebt,
		id,
	).Scan(
		&sale.ID,
		&sale.Name,
		&sale.Product,
		&sale.Quantity,
		&sale.Price,
		&sale.Total,
		&sale.AmountReceived,
		&sale.ChangeAmount,
		&sale.TransactionDate,
		&sale.IsDebt,
		&sale.CreatedAt,
		&sale.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &sale, nil
}

func (r *saleRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM sales WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
