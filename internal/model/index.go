package model

import (
	"fmt"
)

func QueryWithFilters(query string, filters []string) string {
	for i, filter := range filters {
		switch i {
		case 0:
			query += fmt.Sprintf(" WHERE %s ", filter)
		default:
			query += fmt.Sprintf(" AND %s ", filter)
		}
	}
	return query
}
