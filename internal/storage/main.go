package storage

import (
	"context"

	"github.com/MalukiMuthusi/wallet-api/internal/models"
	"github.com/shopspring/decimal"
)

type Store interface {
	GetWalletByID(ctx context.Context, id int32) (*models.Wallet, error)

	DebitWallet(ctx context.Context, id int32, amount *decimal.Decimal) (*models.Wallet, error)

	CreditWallet(ctx context.Context, amount decimal.Decimal) (*models.Wallet, error)
}
