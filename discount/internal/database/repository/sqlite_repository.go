package repository

import (
	"database/sql"
	"fmt"
	"go-store/discount/internal/types"

	"github.com/sirupsen/logrus"
)

type SqliteRepository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewSqliteRepository(db *sql.DB, logger *logrus.Logger) *SqliteRepository {
	return &SqliteRepository{
		db:     db,
		logger: logger,
	}
}

func (r *SqliteRepository) GetDiscount(productName string) (*types.Coupon, error) {
	r.logger.Info("Getting discount")

	qry := `SELECT id, product_name, description, amount FROM Discount WHERE product_name = ?`
	row := r.db.QueryRow(qry, productName)

	var discount types.Coupon

	err := row.Scan(&discount.ID, &discount.ProductName, &discount.Description, &discount.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Warn("No discount found for product:", productName)
			return nil, nil
		}
		return nil, err
	}

	return &discount, nil
}

func (r *SqliteRepository) CreateDiscount(coupon *types.Coupon) (*types.Coupon, error) {
	r.logger.Infof("Creating discount for product: %s", coupon.ProductName)

	qry := `INSERT INTO Discount (product_name, description, amount) VALUES (?, ?, ?)`
	result, err := r.db.Exec(qry, coupon.ProductName, coupon.Description, coupon.Amount)
	if err != nil {
		r.logger.Errorf("Failed to create discount for product %s: %v", coupon.ProductName, err)
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		r.logger.Warnf("Failed to retrieve last insert ID for product %s: %v", coupon.ProductName, err)
	} else {
		coupon.ID = fmt.Sprintf("%d", lastInsertID)
		r.logger.Infof("Discount created successfully with ID: %d", lastInsertID)
	}

	return coupon, nil
}

func (r *SqliteRepository) UpdateDiscount(coupon *types.Coupon) error {
	r.logger.Infof("Updating discount for product: %s", coupon.ProductName)

	qry := `UPDATE Discount SET product_name = ?, description = ?, amount = ? WHERE id = ?`
	result, err := r.db.Exec(qry, coupon.ProductName, coupon.Description, coupon.Amount, coupon.ID)
	if err != nil {
		r.logger.Errorf("Failed to update discount for product %s: %v", coupon.ProductName, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Warnf("Failed to retrieve rows affected for product %s: %v", coupon.ProductName, err)
	} else {
		r.logger.Infof("Discount updated successfully with %d rows affected", rowsAffected)
	}

	return nil
}

func (r *SqliteRepository) DeleteDiscount(productName string) error {
	r.logger.Infof("Deleting discount for product: %s", productName)

	qry := `DELETE FROM Discount WHERE product_name = ?`
	result, err := r.db.Exec(qry, productName)
	if err != nil {
		r.logger.Errorf("Failed to delete discount for product %s: %v", productName, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Warnf("Failed to retrieve rows affected for product %s: %v", productName, err)
	} else {
		r.logger.Infof("Discount deleted successfully with %d rows affected", rowsAffected)
	}

	return nil
}
