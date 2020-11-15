package handlers

import (
	"encoding/json"
	"io/ioutil"
	"julo/constants"
	"julo/domain"
	"julo/utils"
	"log"
	"net/http"
	"strconv"
)

// DepositHandler - for handle add deposit wallet by reference.
func (th TransactionHandler) DepositHandler(res http.ResponseWriter, req *http.Request) {
	var param domain.TransactionBody
	var errorData domain.ResponseData

	token := req.Header.Get("Authorization")

	walletIDs, err := utils.GetDecrypt([]byte(th.Config.App.SecretKey), token)
	if err != nil || walletIDs == "0" {
		errorData.Error = "token invalid"
		log.Println("DepositHandler: error token invalid")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletID, err := strconv.ParseInt(walletIDs, 10, 64)
	if err != nil {
		errorData.Error = "fail convert data"
		log.Println("DepositHandler: error convert wallet ID")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletData, err := th.WalletRepo.GetWallet(walletID, constants.WalletStatusActive)
	if err != nil {
		errorData.Error = "fail get wallet. Please try again later"
		log.Println("GetWalletHandler: fail get wallet, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	if walletData.Status != constants.WalletStatusActive {
		errorData.Error = "wallet not enable yet"
		log.Println("GetWalletHandler: wallet not enable yet")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorData.Error = "fail request param empty"
		log.Println("DepositHandler: init wallet body param empty")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		errorData.Error = "fail unmarshall req body"
		log.Println("DepositHandler: can't unmarshall req body, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	if param.Amount == 0 || param.ReferenceID == "" {
		errorData.Error = domain.InitWalletResult{Error: []string{"Missing data for required field."}}
		log.Println("DepositHandler: error param(s) empty")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	transactionData, isExist, err := th.Repo.AddDeposit(walletID, param)
	if err != nil {
		errorData.Error = "fail add deposit. Please try again later"
		log.Println("DepositHandler: fail add deposit, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	if isExist {
		errorData.Error = "transaction already exist"
		log.Println("EnableWalletHandler: transaction already exist")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	transactionData.DepositBy = &walletData.CustomerUUID

	writeResponse(
		res,
		ResponseData{
			Status: constants.Success,
			Data: domain.ResponseData{
				Deposit: transactionData,
			},
		},
		http.StatusCreated,
		nil)
	return
}

// WithdrawalHandler - for handle withdraw deposit wallet by reference.
func (th TransactionHandler) WithdrawalHandler(res http.ResponseWriter, req *http.Request) {
	var param domain.TransactionBody
	var errorData domain.ResponseData

	token := req.Header.Get("Authorization")

	walletIDs, err := utils.GetDecrypt([]byte(th.Config.App.SecretKey), token)
	if err != nil || walletIDs == "0" {
		errorData.Error = "token invalid"
		log.Println("DepositHandler: error token invalid")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletID, err := strconv.ParseInt(walletIDs, 10, 64)
	if err != nil {
		errorData.Error = "fail convert data"
		log.Println("DepositHandler: error convert wallet ID")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletData, err := th.WalletRepo.GetWallet(walletID, constants.WalletStatusActive)
	if err != nil {
		errorData.Error = "fail get wallet. Please try again later"
		log.Println("GetWalletHandler: fail get wallet, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	if walletData.Status != constants.WalletStatusActive {
		errorData.Error = "wallet not enable yet"
		log.Println("GetWalletHandler: wallet not enable yet")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorData.Error = "fail request param empty"
		log.Println("DepositHandler: init wallet body param empty")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		errorData.Error = "fail unmarshall req body"
		log.Println("DepositHandler: can't unmarshall req body, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	if param.Amount == 0 || param.ReferenceID == "" {
		errorData.Error = domain.InitWalletResult{Error: []string{"Missing data for required field."}}
		log.Println("DepositHandler: error param(s) empty")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	if walletData.Balance-param.Amount <= 0 {
		errorData.Error = "wallet balance not enough"
		log.Println("GetWalletHandler: error wallet balance not enough")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	transactionData, isExist, err := th.Repo.AddWithdrawal(walletID, param)
	if err != nil {
		errorData.Error = "fail add deposit. Please try again later"
		log.Println("DepositHandler: fail add deposit, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	if isExist {
		errorData.Error = "transaction already exist"
		log.Println("EnableWalletHandler: transaction already exist")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	transactionData.WithdrawnBy = &walletData.CustomerUUID

	writeResponse(
		res,
		ResponseData{
			Status: constants.Success,
			Data: domain.ResponseData{
				Deposit: transactionData,
			},
		},
		http.StatusCreated,
		nil)
	return
}
