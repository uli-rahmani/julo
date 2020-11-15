package domain

import (
	"gopkg.in/guregu/null.v4"
)

type InitWalletBody struct {
	CustomerUUID string `json:"customer_xid" db:"customer_uuid"`
}

type InitWalletResult struct {
	Token string   `json:"token,omitempty"`
	Error []string `json:"customer_xid,omitempty"`
}

type WalletResult struct {
	UUID         string      `json:"id" db:"wallet_uuid"`
	CustomerUUID string      `json:"owned_by" db:"customer_uuid"`
	Status       string      `json:"status" db:"status"`
	EnableAt     null.String `json:"enable_at" db:"enable_at"`
	DisableAt    null.String `json:"disabled_at" db:"disabled_at"`
	Balance      float64     `json:"balance" db:"balance"`
}

type WalletResponse struct {
	UUID         string  `json:"id" db:"wallet_uuid"`
	CustomerUUID string  `json:"owned_by" db:"customer_uuid"`
	Status       string  `json:"status" db:"status"`
	EnableAt     *string `json:"enabled_at,omitempty" db:"enable_at"`
	DisableAt    *string `json:"disabled_at,omitempty" db:"disabled_at"`
	Balance      float64 `json:"balance" db:"balance"`
}

type WalletData struct {
	ID           int64   `json:"wallet_id" db:"wallet_id"`
	UUID         string  `json:"wallet_uuid" db:"wallet_uuid"`
	CustomerUUID string  `json:"customer_uuid" db:"customer_uuid"`
	Status       string  `json:"status" db:"status"`
	Balance      float64 `json:"balance" db:"balance"`
}
