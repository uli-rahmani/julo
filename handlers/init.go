package handlers

import (
	"julo/domain"
	"julo/interfaces"
)

type WalletHandler struct {
	Repo   interfaces.WalletRepo
	Config *domain.SectionService
}

type TransactionHandler struct {
	Repo       interfaces.TransactionRepo
	WalletRepo interfaces.WalletRepo
	Config     *domain.SectionService
}
