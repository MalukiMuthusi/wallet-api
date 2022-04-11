package handlers

import (
	"github.com/MalukiMuthusi/wallet-api/internal/storage"
	"github.com/gin-gonic/gin"
)

type BalanceHandler struct {
	Store storage.Store
}

func (b *BalanceHandler) Handle(c *gin.Context) {
	// TODO: implement logic

	panic("Not Implemented")
}
