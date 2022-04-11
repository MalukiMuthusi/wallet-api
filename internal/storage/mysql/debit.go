package mysql

import (
	"context"

	"github.com/MalukiMuthusi/wallet-api/internal/models"
	"github.com/shopspring/decimal"
)

func (m *MysqlDB) DebitWallet(ctx context.Context, amount decimal.Decimal) (*models.Wallet, error) {
	panic("Not implemented")
}
