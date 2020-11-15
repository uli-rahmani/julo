package repo

import (
	"errors"
	"julo/utils"
	"julo/constants"
	"julo/domain"
	"log"

	uuid "github.com/nu7hatch/gouuid"
)

func (tr *TransactionRepo) AddDeposit(walletID int64, param domain.TransactionBody) (domain.TransactionResult, bool, error) {
	data := domain.TransactionResult{
		Status:      "success",
		Amount:      param.Amount,
		ReferenceID: param.ReferenceID,
	}

	isExist, err := tr.IsExistTransactionByReferenceID(walletID, param.ReferenceID, constants.TransactionTypeDeposit)
	if err != nil {
		log.Println("Repo | AddDeposit | error get exist data query, err: " + err.Error())
		return data, false, err
	}

	if isExist {
		log.Println("Repo | AddDeposit | data is exist")
		return data, true, nil
	}

	transUUID, err := uuid.NewV4()
	if err != nil {
		log.Println("Repo | AddDeposit | error build wallet transaction UUID, err: " + err.Error())
		return data, false, err
	}

	err = tr.AddTransaction(transUUID.String(), param.ReferenceID, walletID, param.Amount, constants.TransactionTypeDeposit)
	if err != nil {
		log.Println("Repo | AddDeposit | error insert wallet transation data, err: " + err.Error())
		return data, false, err
	}

	err = tr.UpdateWalletBalance(walletID, param.Amount, constants.TransactionTypeDeposit)
	if err != nil {
		log.Println("Repo | AddDeposit | error update wallet balance data, err: " + err.Error())
		return data, false, err
	}

	timeData := utils.GetTimeString()

	data.UUID = transUUID.String()
	data.DepositAt = &timeData

	return data, false, nil
}

func (tr *TransactionRepo) AddWithdrawal(walletID int64, param domain.TransactionBody) (domain.TransactionResult, bool, error) {
	data := domain.TransactionResult{
		Status:      "success",
		Amount:      param.Amount,
		ReferenceID: param.ReferenceID,
	}

	isExist, err := tr.IsExistTransactionByReferenceID(walletID, param.ReferenceID, constants.TransactionTypeWithDrawn)
	if err != nil {
		log.Println("Repo | AddWithdrawal | error get exist data query, err: " + err.Error())
		return data, false, err
	}

	if isExist {
		log.Println("Repo | AddWithdrawal | data is exist")
		return data, true, nil
	}

	transUUID, err := uuid.NewV4()
	if err != nil {
		log.Println("Repo | AddWithdrawal | error build wallet transaction UUID, err: " + err.Error())
		return data, false, err
	}

	err = tr.AddTransaction(transUUID.String(), param.ReferenceID, walletID, param.Amount, constants.TransactionTypeWithDrawn)
	if err != nil {
		log.Println("Repo | AddWithdrawal | error insert wallet transation data, err: " + err.Error())
		return data, false, err
	}

	err = tr.UpdateWalletBalance(walletID, param.Amount, constants.TransactionTypeWithDrawn)
	if err != nil {
		log.Println("Repo | AddWithdrawal | error update wallet balance data, err: " + err.Error())
		return data, false, err
	}

	timeData := utils.GetTimeString()

	data.UUID = transUUID.String()
	data.WithdrawnAt = &timeData

	return data, false, nil
}

func (tr *TransactionRepo) IsExistTransactionByReferenceID(walletID int64, referenceID string, types int) (bool, error) {
	var isExist bool

	query, args, err := tr.DB.In(constants.QueryIsExistTransactionByReferenceID, walletID, referenceID, types)
	if err != nil {
		log.Println("Repo | IsExistTransactionByReferenceID | error build param get query, err: " + err.Error())
		return isExist, err
	}

	query = tr.DB.Rebind(query)
	err = tr.DB.Get(&isExist, query, args...)
	if err != nil {
		log.Println("Repo | IsExistTransactionByReferenceID | error get query, err: " + err.Error())
		return isExist, err
	}

	return isExist, nil
}

func (tr *TransactionRepo) AddTransaction(transUUID, refID string, walletID int64, amount float64, types int) error {
	query, args, err := tr.DB.In(constants.QueryInsertWalletTransaction, transUUID, walletID, amount, refID, types)
	if err != nil {
		log.Println("Repo | AddTransaction | error build param get query, err: " + err.Error())
		return err
	}

	query = tr.DB.Rebind(query)
	res, err := tr.DB.Exec(query, args...)
	if err != nil {
		log.Println("Repo | AddTransaction | error insert query, err: " + err.Error())
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		log.Println("Repo | AddTransaction | error get row affected, err: " + err.Error())
		return err
	}

	if row == 0 {
		log.Println("Repo | AddTransaction | error no row affected")
		return errors.New("no row affected")
	}

	return nil
}

func (tr *TransactionRepo) UpdateWalletBalance(walletID int64, amount float64, types int) error {
	q := constants.QueryAddWalletBalance
	if types == constants.TransactionTypeWithDrawn {
		q = constants.QueryDecreaseWalletBalance
	}

	query, args, err := tr.DB.In(q, amount, walletID)
	if err != nil {
		log.Println("Repo | UpdateWalletBalance | error build param get query, err: " + err.Error())
		return err
	}

	query = tr.DB.Rebind(query)
	res, err := tr.DB.Exec(query, args...)
	if err != nil {
		log.Println("Repo | UpdateWalletBalance | error update query, err: " + err.Error())
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		log.Println("Repo | UpdateWalletBalance | error get row affected, err: " + err.Error())
		return err
	}

	if row == 0 {
		log.Println("Repo | UpdateWalletBalance | error no row affected")
		return errors.New("no row affected")
	}

	return nil
}
