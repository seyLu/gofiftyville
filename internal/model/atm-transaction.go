package model

import (
	"fmt"
	"strings"

	"github.com/seyLu/gofiftyville/internal/store"
)

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

type AtmTransactionsFilter struct {
	Year            int
	Month           int
	Day             int
	AtmLocation     string
	TransactionType string
}

func AtmTransactions(f AtmTransactionsFilter) ([]AtmTransaction, error) {
	var filters []string
	query := `
		SELECT
			id, account_number, year, month, day, atm_location, transaction_type, amount
		FROM atm_transactions
	`
	args := []any{}

	if f.Year != -1 && f.Month != -1 && f.Day != -1 {
		filters = append(filters, "year=? AND month=? AND day=?")
		args = append(args, f.Year, f.Month, f.Day)
	}
	if f.AtmLocation != "" {
		filters = append(filters, "LOWER(atm_location)=?")
		args = append(args, strings.ToLower(f.AtmLocation))
	}
	if f.TransactionType != "" {
		filters = append(filters, "LOWER(transaction_type)=?")
		args = append(args, strings.ToLower(f.TransactionType))
	}

	query = QueryWithFilters(query, filters)

	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("AtmTransaction [year %q month %q day %q atm_location %q transaction_type %q]: %w", f.Year, f.Month, f.Day, f.AtmLocation, f.TransactionType, err)
	}

	var transactions []AtmTransaction
	for rows.Next() {
		var transaction AtmTransaction
		if err := rows.Scan(&transaction.ID, &transaction.AccountNumber, &transaction.Year, &transaction.Month, &transaction.Day, &transaction.AtmLocation, &transaction.TransactionType, &transaction.Amount); err != nil {
			return nil, fmt.Errorf("AtmTransaction [year %q month %q day %q atm_location %q transaction_type %q]: %w", f.Year, f.Month, f.Day, f.AtmLocation, f.TransactionType, err)
		}

		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AtmTransaction [year %q month %q day %q atm_location %q transaction_type %q]: %w", f.Year, f.Month, f.Day, f.AtmLocation, f.TransactionType, err)
	}

	return transactions, nil
}
