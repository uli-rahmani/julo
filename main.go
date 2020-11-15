package main

import (
	"fmt"
	"julo/domain"
	"julo/handlers"
	"julo/infra"
	"julo/interfaces"
	"julo/repo"
	"julo/utils"
	"log"
	"net/http"
	"os"

	"github.com/alexsasharegan/dotenv"
	"github.com/gorilla/mux"
)

func main() {
	conf, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := infra.DBHandler{}
	db.ConnectDB(&conf.DB)

	dbList := make(map[string]interfaces.Database)
	dbList["DB"] = &db

	walletRepo := &repo.WalletRepo{
		DB: dbList["DB"],
	}

	walletHandler := &handlers.WalletHandler{
		Repo:   walletRepo,
		Config: conf,
	}

	transactionRepo := &repo.TransactionRepo{
		DB: dbList["DB"],
	}

	transactionHandler := handlers.TransactionHandler{
		Repo:       transactionRepo,
		WalletRepo: walletRepo,
		Config:     conf,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/init", walletHandler.InitWalletHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/wallet", walletHandler.EnableWalletHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/wallet/deposits", transactionHandler.DepositHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/wallet/withdrawals", transactionHandler.WithdrawalHandler).Methods(http.MethodPost)

	router.HandleFunc("/api/v1/wallet", walletHandler.DisableWalletHandler).Methods(http.MethodPatch)

	router.HandleFunc("/api/v1/wallet", walletHandler.GetWalletHandler).Methods(http.MethodGet)

	port := fmt.Sprintf(":%s", conf.App.Port)
	log.Println("server listen to port ", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func getConfig() (*domain.SectionService, error) {
	var data domain.SectionService

	err := dotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	data.App.Name = os.Getenv("APPS_NAME")
	data.App.Environtment = os.Getenv("APPS_ENV")
	data.App.URL = os.Getenv("APPS_URL")
	data.App.Port = os.Getenv("APPS_PORT")
	data.App.SecretKey = os.Getenv("APPS_SECRET_KEY")

	data.DB.Username = os.Getenv("DATABASE_READ_USERNAME")
	data.DB.Password = os.Getenv("DATABASE_READ_PASSWORD")
	data.DB.URL = os.Getenv("DATABASE_READ_URL")
	data.DB.Port = os.Getenv("DATABASE_READ_PORT")
	data.DB.DBName = os.Getenv("DATABASE_READ_DB_NAME")
	data.DB.MaxIdleConns = utils.GetInt(os.Getenv("DATABASE_READ_MAXIDLECONNS"))
	data.DB.MaxOpenConns = utils.GetInt(os.Getenv("DATABASE_READ_MAXOPENCONNS"))
	data.DB.MaxLifeTime = utils.GetInt(os.Getenv("DATABASE_READ_MAXLIFETIME"))
	data.DB.Timeout = os.Getenv("DATABASE_READ_TIMEOUT")

	data.Redis.Username = os.Getenv("REDIS_USERNAME")
	data.Redis.Password = os.Getenv("REDIS_PASSWORD")
	data.Redis.URL = os.Getenv("REDIS_URL")
	data.Redis.Port = utils.GetInt(os.Getenv("REDIS_PORT"))
	data.Redis.MinIdleConns = utils.GetInt(os.Getenv("REDIS_MINIDLECONNS"))
	data.Redis.Timeout = os.Getenv("REDIS_TIMEOUT")

	return &data, nil
}
