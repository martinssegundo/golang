package main

import (
	"fmt"
	"time"

	entities "mini-banco/entities"
	interfaces "mini-banco/interfaces"
)

func main() {

	transaction := entities.Transaction{
		CreateDate:      time.Now(),
		Amount:          100,
		ReceiverAccount: "1",
		SenderAccount:   "2",
	}

	var transactionPointer *entities.Transaction = &transaction

	transactions := make([]entities.Transaction, 1)
	transactions = append(transactions, transaction)
	newAccount := &entities.CheckingAccount{
		Tranasctions: transactions,
	}

	transactionPointer = &entities.Transaction{
		CreateDate:      time.Now(),
		Amount:          200,
		ReceiverAccount: "1",
		SenderAccount:   "2",
	}

	adicionaTransacao(newAccount, *transactionPointer)

	transactionPointer = &entities.Transaction{
		CreateDate:      time.Now(),
		Amount:          -1000,
		ReceiverAccount: "1",
		SenderAccount:   "2",
	}

	adicionaTransacao(newAccount, *transactionPointer)
	total(newAccount)

}

func adicionaTransacao(account interfaces.Account, transaction entities.Transaction) {
	defer recoverBalanceLessZero()
	account.NoticeTranasaction(transaction)
}

func total(accoun interfaces.Account) {
	fmt.Printf("Total %f", accoun.Total())
}

func recoverBalanceLessZero() {
	if r := recover(); r != nil {
		println("Sua conta não pode ser negativa")
	}
}
