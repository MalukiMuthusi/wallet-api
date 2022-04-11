package utils

import (
	"errors"

	"gorm.io/gorm"
)

const (
	AppName = "wallet"
	Port    = "port"
)

const (
	DbUser           = "DB_USER"
	DbPwd            = "DB_PWD"
	DbName           = "DB_NAME"
	DbPort           = "DB_PORT"
	DbHost           = "DB_HOST"
	DbHostedOnCloud  = "DB_CLOUD"
	DbConnectionName = "DB_INSTANCE_CONNECTION_NAME"
	DbTimeZone       = "DB_TIMEZONE"
)

var (
	ErrRecordNotFound          = gorm.ErrRecordNotFound
	ErrInsufficientFunds       = errors.New("not enough funds")
	ErrInvalidAmountValue      = errors.New("invalid amount value")
	ErrWalletNotFound          = errors.New("wallet not found")
	ErrFailedToProcessRequest  = errors.New("failed to process request")
	ErrRequestNotAllowed       = errors.New("request not allowed")
	ErrInternalServerError     = errors.New("internal server error")
	ErrOperationNotImplemented = errors.New("operation not implemented")
)
