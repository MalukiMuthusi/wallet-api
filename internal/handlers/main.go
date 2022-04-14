package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/wallet-api/internal/models"
	"github.com/MalukiMuthusi/wallet-api/internal/utils"
	"github.com/gin-gonic/gin"
)

// WalletIDParam refers to an id specified in the path parameter
type WalletIDParam struct {
	WalletID uint32 `uri:"wallet_id" binding:"required"`
}

// GetWalletIDFromParam makes sure the wallet_id path parameter value provided is a valid one
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
