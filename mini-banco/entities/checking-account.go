package entities

import (
	"fmt"
)

type CheckingAccount struct {
	Tranasctions []Transaction
	TotalAccount float64
}

func (c *CheckingAccount) NoticeTranasaction(newTransaction Transaction) {
	fmt.Printf("%f\n", newTransaction.Amount)
	checkTotalLessZero(newTransaction, c.TotalAccount)
	c.Tranasctions = append(c.Tranasctions, newTransaction)
	c.TotalAccount += newTransaction.Amount
}

func (c *CheckingAccount) Total() float64 {
	var total float64 = 0
	for _, transaction := range c.Tranasctions {
		total += transaction.Amount
	}
	return total
}

func checkTotalLessZero(newTransaction Transaction, total float64) {
	checkValue := total + newTransaction.Amount
	if checkValue < 0 {
		panic("A conta esta com saldo negativo")
	}
}
