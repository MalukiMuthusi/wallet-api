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
	ParamWalletId = "wallet_id"
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

type CustomErrorCode int32

const (
	InsufficientFunds CustomErrorCode = iota
	NotImplemented
	InternalServerError
	InvalidAmount
	WalletNotFound
	RequestNotAllowed
	FailedToProcessRequest
)

func (e CustomErrorCode) String() string {
	switch e {
	case InsufficientFunds:
		return "INSUFFICIENT_FUNDS"
	case NotImplemented:
		return "NOT_IMPLIMENTED"
	case InternalServerError:
		return "INTERNAL_SERVER_ERROR"
	case InvalidAmount:
		return "INVALID_AMOUNT"
	case WalletNotFound:
		return "WALLET_NOT_FOUND"
	case RequestNotAllowed:
		return "REQUEST_NOT_ALLOWED"
	case FailedToProcessRequest:
		return "FAILED_TO_PROCESS_REQUEST"
	default:
		return "FAILED_TO_PROCESS_REQUEST"
	}
}
