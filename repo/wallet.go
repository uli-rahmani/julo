package repo

import (
	"errors"
	"julo/constants"
	"julo/domain"
	"julo/utils"
	"log"

	uuid "github.com/nu7hatch/gouuid"
	"gopkg.in/guregu/null.v4"
)

// InitWallet - function for insert wallet by customer uuid on db.
func (or *WalletRepo) InitWallet(customerUUID string) (int64, error) {
	var lastInsertID int64
	var currWallet domain.WalletData

	query, args, err := or.DB.In(constants.QueryGetWalletsByCustomerUUID, customerUUID)
	if err != nil {
		log.Println("Repo | InitWallet | error build param get query, err: " + err.Error())
		return lastInsertID, err
	}

	query = or.DB.Rebind(query)
	err = or.DB.Get(&currWallet, query, args...)
	if err != nil {
		log.Println("Repo | InitWallet | error get query, err: " + err.Error())
		return lastInsertID, err
	}

	if currWallet.ID == 0 {
		walletUUID, err := uuid.NewV4()
		if err != nil {
			log.Println("Repo | InitWallet | error build wallet UUID, err: " + err.Error())
			return lastInsertID, err
		}

		query, args, err = or.DB.In(constants.QueryInsertWallet, walletUUID.String(), customerUUID)
		if err != nil {
			log.Println("Repo | InitWallet | error build param query, err: " + err.Error())
			return lastInsertID, err
		}

		query = or.DB.Rebind(query)
		rows, err := or.DB.Query(query, args...)
		if err != nil {
			log.Println("Repo | InitWallet | error exec query, err: " + err.Error())
			return lastInsertID, err
		}

		for rows.Next() {
			err = rows.Scan(&lastInsertID)
			if err != nil {
				log.Println("Repo | InitWallet | error get last insert id, err: " + err.Error())
				return lastInsertID, err
			}
		}
	} else {
		lastInsertID = currWallet.ID
	}

	return lastInsertID, nil
}

// EnableWallet - function for enable wallet by wallet id on db.
func (or *WalletRepo) EnableWallet(walletID int64) (domain.WalletResult, bool, error) {
	wallet, err := or.GetWallet(walletID, constants.WalletStatusActive)
	if err != nil {
		log.Println("Repo | EnableWallet | error get query, err: " + err.Error())
		return wallet, false, err
	}

	if wallet.Status == constants.WalletStatusActive {
		log.Println("Repo | EnableWallet | wallet already active")
		return wallet, true, err
	}

	timeData := utils.GetTimeString()

	query, args, err := or.DB.In(constants.QueryEnableWallet, constants.WalletStatusActive, timeData, timeData, walletID)
	if err != nil {
		log.Println("Repo | EnableWallet | error build param update query, err: " + err.Error())
		return wallet, false, err
	}

	query = or.DB.Rebind(query)
	res, err := or.DB.Exec(query, args...)
	if err != nil {
		log.Println("Repo | EnableWallet | error update query, err: " + err.Error())
		return wallet, false, err
	}

	row, err := res.RowsAffected()
	if err != nil {
		log.Println("Repo | EnableWallet | error get row affected, err: " + err.Error())
		return wallet, false, err
	}

	if row == 0 {
		log.Println("Repo | EnableWallet | error no row affected")
		return wallet, false, errors.New("no row affected")
	}

	wallet.Status = constants.WalletStatusActive
	wallet.EnableAt = null.NewString(timeData, true)

	return wallet, false, nil
}

// DisableWallet - function for disable wallet by wallet id on db.
func (or *WalletRepo) DisableWallet(walletID int64) (domain.WalletResult, bool, error) {
	wallet, err := or.GetWallet(walletID, constants.WalletStatusInactive)
	if err != nil {
		log.Println("Repo | DisableWallet | error get query, err: " + err.Error())
		return wallet, false, err
	}

	if wallet.Status == constants.WalletStatusInactive {
		log.Println("Repo | DisableWallet | wallet already disable")
		return wallet, true, err
	}

	timeData := utils.GetTimeString()

	query, args, err := or.DB.In(constants.QueryDisableWallet, constants.WalletStatusInactive, timeData, timeData, walletID)
	if err != nil {
		log.Println("Repo | DisableWallet | error build param update query, err: " + err.Error())
		return wallet, false, err
	}

	query = or.DB.Rebind(query)
	res, err := or.DB.Exec(query, args...)
	if err != nil {
		log.Println("Repo | DisableWallet | error update query, err: " + err.Error())
		return wallet, false, err
	}

	row, err := res.RowsAffected()
	if err != nil {
		log.Println("Repo | DisableWallet | error get row affected, err: " + err.Error())
		return wallet, false, err
	}

	if row == 0 {
		log.Println("Repo | DisableWallet | error no row affected")
		return wallet, false, errors.New("no row affected")
	}

	wallet.Status = constants.WalletStatusInactive
	wallet.DisableAt = null.NewString(timeData, true)

	return wallet, false, nil
}

// GetWallet - function for get wallet by wallet id on db.
func (or *WalletRepo) GetWallet(walletID int64, types string) (domain.WalletResult, error) {
	var wallet domain.WalletResult

	q := constants.QueryGetWalletsByWalletID
	if types == constants.WalletStatusInactive {
		q = constants.QueryGetDisableWalletsByWalletID
	}

	query, args, err := or.DB.In(q, walletID)
	if err != nil {
		log.Println("Repo | GetWallet | error build param get query, err: " + err.Error())
		return wallet, err
	}

	query = or.DB.Rebind(query)
	err = or.DB.Get(&wallet, query, args...)
	if err != nil {
		log.Println("Repo | GetWallet | error get query, err: " + err.Error())
		return wallet, err
	}

	return wallet, nil
}
