package mysql

import (
	"context"
	"errors"

	"github.com/MalukiMuthusi/wallet-api/internal/logger"
	"github.com/MalukiMuthusi/wallet-api/internal/models"
	"github.com/MalukiMuthusi/wallet-api/internal/utils"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (m *MysqlDB) DebitWallet(ctx context.Context, id int32, amount *decimal.Decimal) (*models.Wallet, error) {
	var wallet models.Wallet

	res := m.Db.First(&wallet, id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {

		logger.Log.WithFields(logrus.Fields{"id": id}).Info("wallet not found")

		return nil, utils.ErrWalletNotFound

	} else if res.Error != nil {

		logger.Log.WithFields(logrus.Fields{"id": id, "err": res.Error}).Info("failed to get wallet")

		return nil, utils.ErrFailedToProcessRequest
	}

	newBalance := wallet.Balance.Sub(*amount)
	wallet.Balance = newBalance

	if newBalance.IsNegative() {
		logger.Log.WithField("id", id).Info("insufficient funds")
		return nil, utils.ErrInsufficientFunds
	}

	newRes := m.Db.Model(&wallet).Update("balance", newBalance)
	if errors.Is(newRes.Error, gorm.ErrRecordNotFound) {

		logger.Log.WithFields(logrus.Fields{"id": id}).Info("wallet not found")

		return nil, utils.ErrWalletNotFound

	} else if newRes.Error != nil {

		logger.Log.WithFields(logrus.Fields{"id": id, "err": newRes.Error}).Info("failed to get wallet")

		return nil, utils.ErrFailedToProcessRequest
	}

	return &wallet, nil
}
