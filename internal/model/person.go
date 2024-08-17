package model

import (
	"fmt"
	"strings"

	"github.com/guregu/null/v5"
	"github.com/seyLu/gofiftyville/internal/store"
)

type Person struct {
	ID             int
	Name           string
	PhoneNumber    null.String
	PassportNumber null.Int
	LicensePlate   null.String
}

type PersonBankAccount struct {
	Person
	BankAccount
}

type PeopleFilter struct {
	LicensePlate  string
	AccountNumber int
}

func People(f PeopleFilter) ([]PersonBankAccount, error) {
	var filters []string
	query := `
		SELECT
			id, name, phone_number, passport_number, license_plate,
			account_number, person_id, creation_year
		FROM people
		INNER JOIN bank_accounts
			ON people.id=bank_accounts.person_id
	`
	args := []any{}

	if f.LicensePlate != "" {
		filters = append(filters, "LOWER(license_plate)=?")
		args = append(args, strings.ToLower(f.LicensePlate))
	}
	if f.AccountNumber != -1 {
		filters = append(filters, "account_number=?")
		args = append(args, f.AccountNumber)
	}

	query = QueryWithFilters(query, filters)

	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("People %+v: %w", f, err)
	}
	defer rows.Close()

	var accounts []PersonBankAccount
	for rows.Next() {
		var account PersonBankAccount
		if err := rows.Scan(&account.ID, &account.Name, &account.PhoneNumber, &account.PassportNumber, &account.LicensePlate, &account.AccountNumber, &account.PersonId, &account.CreationYear); err != nil {
			return nil, fmt.Errorf("People %+v: %w", f, err)
		}
		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("People %+v: %w", f, err)
	}

	return accounts, nil
}
