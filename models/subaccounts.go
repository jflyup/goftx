package models

import (
	"time"

	"github.com/jflyup/decimal"
)

type SubAccount struct {
	Nickname    string `json:"nickname"`
	Deletable   bool   `json:"deletable"`
	Editable    bool   `json:"editable"`
	Competition bool   `json:"competition,omitempty"`
}

type Balance struct {
	Coin                   string          `json:"coin"`
	Free                   decimal.Decimal `json:"free"`
	AvailableWithoutBorrow decimal.Decimal
	Total                  decimal.Decimal `json:"total"`
}

type TransferPayload struct {
	Coin        string          `json:"coin"`
	Size        decimal.Decimal `json:"size"`
	Source      *string         `json:"source"`
	Destination *string         `json:"destination"`
}

type TransferResponse struct {
	ID     int64           `json:"id"`
	Coin   string          `json:"coin"`
	Size   decimal.Decimal `json:"size"`
	Time   time.Time       `json:"time"`
	Notes  string          `json:"notes"`
	Status TransferStatus  `json:"status"`
}
