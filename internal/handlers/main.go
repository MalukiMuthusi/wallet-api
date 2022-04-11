package handlers

type WalletIDParam struct {
	WalletID int32 `uri:"wallet_id" binding:"required"`
}
