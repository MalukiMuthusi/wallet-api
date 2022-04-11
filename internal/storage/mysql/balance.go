package mysql

import (
	"context"

	"github.com/MalukiMuthusi/wallet-api/internal/models"
)

func (m *MysqlDB) GetWalletByID(ctx context.Context, id string) (*models.Wallet, error) {
	panic("Not implemented")
}
