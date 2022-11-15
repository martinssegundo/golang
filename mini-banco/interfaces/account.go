package interfaces

import "mini-banco/entities"

type Account interface {
	NoticeTranasaction(newTransaction entities.Transaction)
	Total()
}
