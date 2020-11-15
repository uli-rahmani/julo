package domain

type ResponseData struct {
	Error      interface{} `json:"error,omitempty"`
	Wallet     interface{} `json:"wallet,omitempty"`
	Withdrawal interface{} `json:"withdrawal,omitempty"`
	Deposit    interface{} `json:"deposit,omitempty"`
}
