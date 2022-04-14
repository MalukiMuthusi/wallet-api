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

	walletID, err := GetWalletIDFromParam(c)
	if err != nil {
		return
	}

	wallet, err := b.Store.GetWalletByID(c.Copy().Request.Context(), walletID.WalletID)

	if err != nil {

		var status int
		var basicError models.BasicError

		switch err {

		case utils.ErrWalletNotFound:
			basicError = models.BasicError{
				Code:    utils.WalletNotFound.String(),
				Message: "wallet not found",
			}

			status = http.StatusNotFound

		case utils.ErrFailedToProcessRequest:
			basicError = models.BasicError{
				Code:    utils.InternalServerError.String(),
				Message: "failed to complete processing request",
			}

			status = http.StatusInternalServerError

		case utils.ErrOperationNotImplemented:
			basicError = models.BasicError{
				Code:    utils.NotImplemented.String(),
				Message: "operation not implemented on the server",
			}

			status = http.StatusNotImplemented

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
