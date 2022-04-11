package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/wallet-api/internal/models"
	"github.com/MalukiMuthusi/wallet-api/internal/storage"
	"github.com/MalukiMuthusi/wallet-api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type CreditHandler struct {
	Store storage.Store
}

func (cr *CreditHandler) Handle(c *gin.Context) {
	walletID, err := GetWalletIDFromParam(c)
	if err != nil {
		// Will never get here
		return
	}

	var amount models.Amount

	if c.ShouldBind(&amount) != nil {
		e := models.BasicError{
			Code:    utils.InvalidAmount.String(),
			Message: "provide a valid post request",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	a, err := decimal.NewFromString(amount.Value)
	if err != nil {
		e := models.BasicError{
			Code:    utils.InvalidAmount.String(),
			Message: "provide a valid amount value",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	wallet, err := cr.Store.CreditWallet(c.Copy().Request.Context(), walletID.WalletID, &a)

	if err != nil {
		switch err {

		case utils.ErrOperationNotImplemented:
			e := models.BasicError{
				Code:    utils.NotImplemented.String(),
				Message: "operation not implemented on the server",
			}

			c.JSON(http.StatusNotImplemented, e)
			return

		case utils.ErrFailedToProcessRequest:
			e := models.BasicError{
				Code:    utils.InternalServerError.String(),
				Message: "failed to complete processing request",
			}

			c.JSON(http.StatusInternalServerError, e)
			return

		default:
			e := models.BasicError{
				Code:    utils.InternalServerError.String(),
				Message: "failed to process request",
			}

			c.JSON(http.StatusInternalServerError, e)
			return
		}

	}

	c.JSON(http.StatusOK, wallet)
}
