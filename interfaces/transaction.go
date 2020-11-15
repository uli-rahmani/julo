package interfaces

import "julo/domain"

type TransactionRepo interface {
	AddDeposit(walletID int64, param domain.TransactionBody) (domain.TransactionResult, bool, error)
	AddWithdrawal(walletID int64, param domain.TransactionBody) (domain.TransactionResult, bool, error)
}
