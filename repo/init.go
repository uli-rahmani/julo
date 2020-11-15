package repo

import "julo/interfaces"

type WalletRepo struct {
	DB interfaces.Database
}

type TransactionRepo struct {
	DB interfaces.Database
}
