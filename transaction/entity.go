package transaction

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	Id        uuid.UUID `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserId    int
	MethodId  int
	Item      string
	Amount    int
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentMethod struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
