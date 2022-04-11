package handlers

import (
	"github.com/MalukiMuthusi/wallet-api/internal/storage"
	"github.com/gin-gonic/gin"
)

type DebitHandler struct {
	Store storage.Store
}

func (d *DebitHandler) Handle(c *gin.Context) {
	// TODO: Add logic

	panic("Not implemented")
}
