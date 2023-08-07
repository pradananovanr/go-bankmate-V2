package entity

import "time"

type Deposit struct {
	ID_Deposit          int       `json:"id_deposit"`
	ID_Customer         int       `json:"id_customer"`
	Deposit_Amount      float32   `json:"deposit_amount"`
	Deposit_Description string    `json:"deposit_description"`
	Date_Time           time.Time `json:"date_time"`
}

type DepositRequest struct {
	Deposit_Amount      float32 `json:"deposit_amount"`
	Deposit_Description string  `json:"deposit_description"`
}
