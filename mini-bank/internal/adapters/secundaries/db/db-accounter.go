package db

import entities "mini-bank/internal/domain/entities"

type BDAccount[T any] interface {
	Dboperater[T]
	findLast30DayTransactions(accountID int) []entities.Transaction
}
