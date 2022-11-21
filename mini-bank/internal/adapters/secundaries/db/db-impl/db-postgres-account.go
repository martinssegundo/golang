package dbimpl

import (
	domain "mini-bank/internal/domain/entities"
)

type AccountEntity struct {
	ID           int
	TotalAccount float64
}

func (ae AccountEntity) save(account AccountEntity) int {
	return 1
}

func (ae AccountEntity) update(account AccountEntity) int {
	return 1
}

func (ae AccountEntity) findLast30DayTransactions(accountID int) []domain.Transaction {
	return nil
}
