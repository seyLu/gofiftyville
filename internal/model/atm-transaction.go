package model

type AtmTransaction struct {
	ID              int
	AccountNumber   int
	Year            int
	Month           int
	Day             int
	AtmLocation     string
	TransactionType string
	Amount          int
}
