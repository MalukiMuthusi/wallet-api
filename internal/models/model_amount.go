package models

import (
	"github.com/MalukiMuthusi/wallet-api/internal/utils"
	"github.com/shopspring/decimal"
)

type Amount struct {
	Value string `json:"value,omitempty" form:"value"`
	// Only one currency type is currently supported, euro
	Currency string `json:"currency,omitempty" form:"currency"`
}

func (a Amount) ValueFromString() (*decimal.Decimal, error) {

	amount, err := decimal.NewFromString(a.Value)

	if err != nil {

		return nil, utils.ErrInvalidAmountValue
	}

	if amount.IsNegative() {

		return nil, utils.ErrInvalidAmountValue
	}

	return &amount, nil
}
