package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/wallet-api/internal/models"
	"github.com/MalukiMuthusi/wallet-api/internal/storage"
	"github.com/MalukiMuthusi/wallet-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type BalanceHandler struct {
	Store storage.Store
}

func (b *BalanceHandler) Handle(c *gin.Context) {

	var walletID WalletIDParam

	if err := c.ShouldBindUri(&walletID); err != nil {
		e := models.BasicError{
			Code:    utils.InvalidAmount.String(),
			Message: "provide a valid wallet id parameter in the request",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	wallet, err := b.Store.GetWalletByID(c.Request.Context(), walletID.WalletID)

	if err != nil {

		switch err {

		case utils.ErrWalletNotFound:
			e := models.BasicError{
				Code:    utils.WalletNotFound.String(),
				Message: "wallet not found",
			}
			c.JSON(http.StatusNotFound, e)
			return

		case utils.ErrFailedToProcessRequest:
			e := models.BasicError{
				Code:    utils.InternalServerError.String(),
				Message: "failed to complete processing request",
			}

			c.JSON(http.StatusInternalServerError, e)
			return

		case utils.ErrOperationNotImplemented:
			e := models.BasicError{
				Code:    utils.NotImplemented.String(),
				Message: "operation not implemented on the server",
			}

			c.JSON(http.StatusNotImplemented, e)
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
