package interfaces

import "mini-bank/internal/domain/entities"

type Account interface {
	NoticeTranasaction(newTransaction entities.Transaction)
	Total() float64
}
