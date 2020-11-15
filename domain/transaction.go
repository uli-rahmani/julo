package domain

type TransactionBody struct {
	Amount      float64 `json:"amount"`
	ReferenceID string  `json:"reference_id"`
}

type TransactionResult struct {
	UUID        string  `json:"id" db:"transaction_uuid"`
	DepositBy   *string `json:"deposit_by,omitempty"`
	WithdrawnBy *string `json:"withdrawn_by,omitempty"`
	Status      string  `json:"status" db:"status"`
	DepositAt   *string `json:"deposit_at,omitempty" db:"created_at"`
	WithdrawnAt *string `json:"withdrawn_at,omitempty" db:"created_at"`
	Amount      float64 `json:"amount" db:"amount"`
	ReferenceID string  `json:"reference_id" db:"reference_id"`
}
