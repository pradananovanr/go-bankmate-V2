package entity

import "time"

type Log struct {
	ID_Log      int       `json:"id_log"`
	ID_Customer int       `json:"id_customer"`
	Activity    string    `json:"activity"`
	Date_Time   time.Time `json:"date_time"`
}
