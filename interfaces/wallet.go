package interfaces

import "julo/domain"

type WalletRepo interface {
	InitWallet(customerUUID string) (int64, error)
	EnableWallet(walletID int64) (domain.WalletResult, bool, error)
	DisableWallet(walletID int64) (domain.WalletResult, bool, error)
	GetWallet(walletID int64, types string) (domain.WalletResult, error)
}
