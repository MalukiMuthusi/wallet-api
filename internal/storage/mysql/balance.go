package mysql

import (
	"context"
	"errors"

	"github.com/MalukiMuthusi/wallet-api/internal/logger"
	"github.com/MalukiMuthusi/wallet-api/internal/models"
	"github.com/MalukiMuthusi/wallet-api/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (m *MysqlDB) GetWalletByID(ctx context.Context, id uint32) (*models.Wallet, error) {
	
	var wallet models.Wallet

	res := m.Db.First(&wallet, id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {

		logger.Log.WithFields(logrus.Fields{"id": id}).Info("wallet not found")

		return nil, utils.ErrWalletNotFound

	} else if res.Error != nil {

		logger.Log.WithFields(logrus.Fields{"id": id, "err": res.Error}).Info("failed to get wallet")

		return nil, utils.ErrFailedToProcessRequest
	}

	return &wallet, nil
}
