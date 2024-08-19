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
	Name           string
	PhoneNumber    null.String
	PassportNumber null.Int
	LicensePlate   null.String
	AccountNumber  int
	CreationYear   int
}

type PeopleFilter struct {
	LicensePlates  []string
	AccountNumbers []int
	PhoneNumbers   []string
}

func People(f PeopleFilter) ([]PersonBankAccount, error) {
	var filters []string
	query := `
		SELECT
			p.name, p.phone_number, p.passport_number, p.license_plate,
			b.account_number, b.creation_year
		FROM people AS p
		INNER JOIN bank_accounts AS b
			ON b.person_id=p.id
	`
	args := []any{}

	if len(f.LicensePlates) > 0 {
		placeholders := strings.Repeat("?, ", len(f.LicensePlates))
		placeholders = strings.TrimSuffix(placeholders, ", ")
		filters = append(filters, fmt.Sprintf("license_plate IN (%v)", placeholders))
		for _, licensePlate := range f.LicensePlates {
			args = append(args, licensePlate)
		}
	}
	if len(f.AccountNumbers) > 0 {
		placeholders := strings.Repeat("?, ", len(f.AccountNumbers))
		placeholders = strings.TrimSuffix(placeholders, ", ")
		filters = append(filters, fmt.Sprintf("account_number IN (%v)", placeholders))
		for _, accountNumber := range f.AccountNumbers {
			args = append(args, accountNumber)
		}
	}
	if len(f.PhoneNumbers) > 0 {
		placeholders := strings.Repeat("?, ", len(f.PhoneNumbers))
		placeholders = strings.TrimSuffix(placeholders, ", ")
		filters = append(filters, fmt.Sprintf("phone_number IN (%v)", placeholders))
		for _, phoneNumber := range f.PhoneNumbers {
			args = append(args, phoneNumber)
		}
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
		if err := rows.Scan(
			&account.Name, &account.PhoneNumber, &account.PassportNumber, &account.LicensePlate,
			&account.AccountNumber, &account.CreationYear,
		); err != nil {
			return nil, fmt.Errorf("People %+v: %w", f, err)
		}
		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("People %+v: %w", f, err)
	}

	return accounts, nil
}
