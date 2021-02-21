package transaction

type TransactionInput struct {
	UserId   int    `json:"user_id" binding:"required"`
	MethodId int    `json:"method_id" binding:"required"`
	Item     string `json:"item" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
}
