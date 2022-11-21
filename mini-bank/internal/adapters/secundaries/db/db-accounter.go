package db

import (
	entities "mini-bank/internal/adapters/secundaries/db/db-impl"
	domain "mini-bank/internal/domain/entities"
)

type BDAccount interface {
	save(account entities.AccountEntity) int
	update(account entities.AccountEntity) int
	findLast30DayTransactions(accountID int) []domain.Transaction
}
