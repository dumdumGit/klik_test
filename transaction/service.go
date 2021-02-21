package transaction

import (
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Service interface {
	CreateTransaction(input TransactionInput) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransaction(input TransactionInput) (Transaction, error) {
	transaction := Transaction{}
	methodId := input.MethodId
	unixTime := strconv.FormatInt(time.Now().UnixNano(), 11)
	transId := uuid.NewV4()
	ctrans := transId.String()
	transCode := ctrans[len(ctrans)-6:]

	paymentMethod, err := s.repository.FindPaymentMethod(methodId)
	if err != nil {
		return transaction, err
	}
	transactionCode := paymentMethod.Code + unixTime[0:11] + string(transCode)

	transaction.Id = transId
	transaction.UserId = input.UserId
	transaction.MethodId = methodId
	transaction.Item = input.Item
	transaction.Amount = input.Amount
	transaction.Code = transactionCode

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
