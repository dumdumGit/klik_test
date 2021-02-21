package transaction

import uuid "github.com/satori/go.uuid"

type TransactionFormatter struct {
	Id   uuid.UUID `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Item string    `json:"item"`
	Code string    `json:"code"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		Id:   transaction.Id,
		Item: transaction.Item,
		Code: transaction.Code,
	}

	return formatter
}
