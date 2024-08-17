package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

type AtmTransaction struct {
	AccountNumber   int    `json:"accountNumber"`
	DateFormatted   string `json:"dateFormatted"`
	AtmLocation     string `json:"atmLocation"`
	TransactionType string `json:"transactionType"`
	Amount          int    `json:"amount"`
}

func GetAtmTransactions(c *gin.Context) {
	request := c.Request.URL.Query()

	f := model.AtmTransactionsFilter{
		Year:            -1,
		Month:           -1,
		Day:             -1,
		AtmLocation:     strings.TrimSpace(request.Get("atm-location")),
		TransactionType: strings.TrimSpace(request.Get("transaction-type")),
	}

	transactionDate := strings.TrimSpace(request.Get("date"))
	if transactionDate != "" {
		parsedTransactionDate, err := ParseDate(transactionDate)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		f.Year, f.Month, f.Day = parsedTransactionDate.Year, parsedTransactionDate.Month, parsedTransactionDate.Day
	}

	atmTransactions, err := model.AtmTransactions(f)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting AtmTransactions (date %s, atm-location %s, transaction-type %s): %v", transactionDate, f.AtmLocation, f.TransactionType, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var transactions []AtmTransaction
	for _, transaction := range atmTransactions {
		dateFormatted := fmt.Sprintf("%s %d, %d", time.Month(transaction.Month).String(), transaction.Day, transaction.Year)

		transactions = append(transactions, AtmTransaction{
			AccountNumber:   transaction.AccountNumber,
			DateFormatted:   dateFormatted,
			AtmLocation:     transaction.AtmLocation,
			TransactionType: transaction.TransactionType,
			Amount:          transaction.Amount,
		})
	}

	c.JSON(http.StatusOK, transactions)
}
