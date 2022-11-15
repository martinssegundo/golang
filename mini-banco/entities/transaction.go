package entities

import "time"

type Transaction struct {
	CreateDate      time.Time
	Amount          float64
	ReceiverAccount string
	SenderAccount   string
}
