package handlers

import (
	"github.com/MalukiMuthusi/wallet-api/internal/storage"
	"github.com/gin-gonic/gin"
)

type CreditHandler struct {
	Store storage.Store
}

func (credit *CreditHandler) Handle(c *gin.Context) {
	// TODO: Add logic

	panic("Not implemented")
}
