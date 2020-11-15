package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"julo/constants"
	"julo/domain"
	"julo/utils"
	"log"
	"net/http"
	"strconv"
)

// InitWalletHandler - for handle init wallet by customer reference.
func (oh WalletHandler) InitWalletHandler(res http.ResponseWriter, req *http.Request) {
	var param domain.InitWalletBody
	var errorData domain.ResponseData
	// var tokenGen utils.ED

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorData.Error = "fail request param empty"
		log.Println("InitWalletHandler: init wallet body param empty")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		errorData.Error = "fail unmarshall req body"
		log.Println("InitWalletHandler: can't unmarshall req body, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	if param.CustomerUUID == "" {
		errorData.Error = domain.InitWalletResult{Error: []string{"Missing data for required field."}}
		log.Println("InitWalletHandler: error customer id param empty")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletID, err := oh.Repo.InitWallet(param.CustomerUUID)
	if err != nil {
		errorData.Error = "fail init wallet. Please try again later"
		log.Println("InitWalletHandler: fail init wallet, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	token, err := utils.GetEncrypt([]byte(oh.Config.App.SecretKey), fmt.Sprintf("%v", walletID))
	if err != nil {
		errorData.Error = "token invalid"
		log.Println("InitWalletHandler: fail token invalid, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	writeResponse(
		res,
		ResponseData{
			Status: constants.Success,
			Data: domain.InitWalletResult{
				Token: token,
			},
		},
		http.StatusCreated,
		nil)
	return
}

// EnableWalletHandler - for handle enable wallet.
func (oh WalletHandler) EnableWalletHandler(res http.ResponseWriter, req *http.Request) {
	var errorData domain.ResponseData
	token := req.Header.Get("Authorization")

	walletIDs, err := utils.GetDecrypt([]byte(oh.Config.App.SecretKey), token)
	if err != nil || walletIDs == "0" {
		errorData.Error = "token invalid"
		log.Println("EnableWalletHandler: error token invalid")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletID, err := strconv.ParseInt(walletIDs, 10, 64)
	if err != nil {
		errorData.Error = "fail convert data"
		log.Println("EnableWalletHandler: error convert wallet ID")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletData, alreadyActive, err := oh.Repo.EnableWallet(walletID)
	if err != nil {
		errorData.Error = "fail enable wallet. Please try again later"
		log.Println("EnableWalletHandler: fail enable wallet, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	if alreadyActive {
		errorData.Error = "Already enabled"
		log.Println("EnableWalletHandler: wallet already active")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletData.EnableAt.String, err = utils.ToFormatTime(walletData.EnableAt.String)
	if err != nil {
		errorData.Error = "fail convert date time"
		log.Println("EnableWalletHandler: fail convert date time, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	writeResponse(
		res,
		ResponseData{
			Status: constants.Success,
			Data: domain.ResponseData{
				Wallet: domain.WalletResponse{
					UUID:         walletData.UUID,
					CustomerUUID: walletData.CustomerUUID,
					Status:       walletData.Status,
					EnableAt:     &walletData.EnableAt.String,
					Balance:      walletData.Balance,
				},
			},
		},
		http.StatusCreated,
		nil)
	return
}

// DisableWalletHandler - for handle disable wallet.
func (oh WalletHandler) DisableWalletHandler(res http.ResponseWriter, req *http.Request) {
	var errorData domain.ResponseData
	token := req.Header.Get("Authorization")

	walletIDs, err := utils.GetDecrypt([]byte(oh.Config.App.SecretKey), token)
	if err != nil || walletIDs == "0" {
		errorData.Error = "token invalid"
		log.Println("DisableWalletHandler: error token invalid")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletID, err := strconv.ParseInt(walletIDs, 10, 64)
	if err != nil {
		errorData.Error = "fail convert data"
		log.Println("DisableWalletHandler: error convert wallet ID")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletData, alreadyDisable, err := oh.Repo.DisableWallet(walletID)
	if err != nil {
		errorData.Error = "fail disable wallet. Please try again later"
		log.Println("DisableWalletHandler: fail disable wallet, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	if alreadyDisable {
		errorData.Error = "Already disable"
		log.Println("DisableWalletHandler: wallet already disable")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletData.DisableAt.String, err = utils.ToFormatTime(walletData.DisableAt.String)
	if err != nil {
		errorData.Error = "fail convert date time"
		log.Println("DisableWalletHandler: fail convert date time, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	writeResponse(
		res,
		ResponseData{
			Status: constants.Success,
			Data: domain.ResponseData{
				Wallet: domain.WalletResponse{
					UUID:         walletData.UUID,
					CustomerUUID: walletData.CustomerUUID,
					Status:       walletData.Status,
					DisableAt:    &walletData.DisableAt.String,
					Balance:      walletData.Balance,
				},
			},
		},
		http.StatusOK,
		nil)
	return
}

// GetWalletHandler - for handle get wallet.
func (oh WalletHandler) GetWalletHandler(res http.ResponseWriter, req *http.Request) {
	var errorData domain.ResponseData
	token := req.Header.Get("Authorization")

	walletIDs, err := utils.GetDecrypt([]byte(oh.Config.App.SecretKey), token)
	if err != nil || walletIDs == "0" {
		errorData.Error = "token invalid"
		log.Println("GetWalletHandler: error token invalid")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletID, err := strconv.ParseInt(walletIDs, 10, 64)
	if err != nil {
		errorData.Error = "fail convert data"
		log.Println("GetWalletHandler: error convert wallet ID")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	walletData, err := oh.Repo.GetWallet(walletID, constants.WalletStatusActive)
	if err != nil {
		errorData.Error = "fail get wallet. Please try again later"
		log.Println("GetWalletHandler: fail get wallet, err: " + err.Error())
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusInternalServerError, nil)
		return
	}

	if walletData.Status != constants.WalletStatusActive {
		errorData.Error = "Disabled"
		log.Println("GetWalletHandler: wallet still disable")
		writeResponse(res, ResponseData{Status: constants.Fail, Data: errorData}, http.StatusBadRequest, nil)
		return
	}

	writeResponse(
		res,
		ResponseData{
			Status: constants.Success,
			Data: domain.ResponseData{
				Wallet: domain.WalletResponse{
					UUID:         walletData.UUID,
					CustomerUUID: walletData.CustomerUUID,
					Status:       walletData.Status,
					EnableAt:     &walletData.EnableAt.String,
					Balance:      walletData.Balance,
				},
			},
		},
		http.StatusOK,
		nil)
	return
}
