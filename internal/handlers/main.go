package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/wallet-api/internal/models"
	"github.com/MalukiMuthusi/wallet-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type WalletIDParam struct {
	WalletID uint32 `uri:"wallet_id" binding:"required"`
}

func GetWalletIDFromParam(c *gin.Context) (*WalletIDParam, error) {
	var walletID WalletIDParam

	if err := c.ShouldBindUri(&walletID); err != nil {
		e := models.BasicError{
			Code:    utils.InvalidWalletIdParam.String(),
			Message: "provide a valid wallet id parameter in the request",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return nil, utils.ErrInvalidAmountValue
	}

	return &walletID, nil
}
