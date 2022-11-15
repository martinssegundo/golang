package entities

import (
	"fmt"
)

type CheckingAccount struct {
	Tranasctions []Transaction
}

func (c *CheckingAccount) NoticeTranasaction(newTransaction Transaction) {
	fmt.Printf("%f\n", newTransaction.Amount)
	checkTotalLessZero(newTransaction, c.Tranasctions)
	c.Tranasctions = append(c.Tranasctions, newTransaction)
}

func (c *CheckingAccount) Total() float64 {
	var total float64 = 0
	for _, transaction := range c.Tranasctions {
		total += transaction.Amount
	}
	return total
}

func checkTotalLessZero(newTransaction Transaction, transactions []Transaction) {
	var total float64 = 0
	for _, transaction := range transactions {
		total += transaction.Amount
	}

	total += newTransaction.Amount

	if total < 0 {
		panic("A conta esta com saldo negativo")
	}
}
