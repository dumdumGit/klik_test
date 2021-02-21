package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(transaction Transaction) (Transaction, error)
	FindPaymentMethod(method int) (PaymentMethod, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindPaymentMethod(method int) (PaymentMethod, error) {
	var payment PaymentMethod

	err := r.db.Where("id = ?", method).Find(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}
