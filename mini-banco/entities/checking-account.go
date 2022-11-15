package entities

import (
	"fmt"
)

type CheckingAccount struct {
	Tranasctions []Transaction
}

func (c *CheckingAccount) NoticeTranasaction(newTransaction Transaction) {
	fmt.Printf("%f\n", newTransaction.Amount)
	c.Tranasctions = append(c.Tranasctions, newTransaction)
}

func (c *CheckingAccount) Total() {
	var total float64 = 0
	for _, transaction := range c.Tranasctions {
		total += transaction.Amount
	}
	fmt.Printf("O total é: %f\n", total)
}
