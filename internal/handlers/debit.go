package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/wallet-api/internal/models"
	"github.com/MalukiMuthusi/wallet-api/internal/storage"
	"github.com/MalukiMuthusi/wallet-api/internal/utils"
	"github.com/gin-gonic/gin"
)

// DebitHandler handles requests to the /debit endpoint

type DebitHandler struct {
	Store storage.Store
}

func (d *DebitHandler) Handle(c *gin.Context) {

	walletID, err := GetWalletIDFromParam(c)
	if err != nil {
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

	a, err := amount.ValueFromString()
	if err != nil {

		var status int
		var basicError models.BasicError

		switch err {

		case utils.ErrInvalidAmountValue, utils.ErrInvalidAmountValue:
			basicError = models.BasicError{
				Code:    utils.InvalidAmount.String(),
				Message: "provide a valid amount value",
			}
			status = http.StatusUnprocessableEntity

		default:
			basicError = models.BasicError{
				Code:    utils.InvalidAmount.String(),
				Message: "provide a valid amount value",
			}
			status = http.StatusUnprocessableEntity
		}

		c.JSON(status, basicError)
		return
	}

	wallet, err := d.Store.DebitWallet(c.Copy().Request.Context(), walletID.WalletID, a)

	if err != nil {

		var status int
		var basicError models.BasicError

		switch err {

		case utils.ErrInsufficientFunds:
			basicError = models.BasicError{
				Code:    utils.InsufficientFunds.String(),
				Message: "insufficient funds",
			}

			status = http.StatusOK

		case utils.ErrOperationNotImplemented:
			basicError = models.BasicError{
				Code:    utils.NotImplemented.String(),
				Message: "operation not implemented on the server",
			}

			status = http.StatusNotImplemented

		case utils.ErrFailedToProcessRequest:
			basicError = models.BasicError{
				Code:    utils.InternalServerError.String(),
				Message: "failed to complete processing request",
			}

			status = http.StatusInternalServerError

		default:
			basicError = models.BasicError{
				Code:    utils.InternalServerError.String(),
				Message: "failed to process request",
			}

			status = http.StatusInternalServerError
		}

		c.JSON(status, basicError)
		return
	}

	c.JSON(http.StatusOK, wallet)
}
