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
	Date            string `json:"date"`
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
			errMsg := fmt.Sprintf("(1) controller.GetAtmTransactions (date %s, atm-location %s, transaction-type %s): %v", transactionDate, f.AtmLocation, f.TransactionType, err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}

		f.Year, f.Month, f.Day = parsedTransactionDate.Year, parsedTransactionDate.Month, parsedTransactionDate.Day
	}

	atmTransactions, err := model.AtmTransactions(f)
	if err != nil {
		errMsg := fmt.Sprintf("(2) controller.GetAtmTransactions (date %s, atm-location %s, transaction-type %s): %v", transactionDate, f.AtmLocation, f.TransactionType, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var transactions []AtmTransaction
	for _, transaction := range atmTransactions {
		date := fmt.Sprintf("%s %d, %d", time.Month(transaction.Month).String(), transaction.Day, transaction.Year)

		transactions = append(transactions, AtmTransaction{
			AccountNumber:   transaction.AccountNumber,
			Date:            date,
			AtmLocation:     transaction.AtmLocation,
			TransactionType: transaction.TransactionType,
			Amount:          transaction.Amount,
		})
	}

	c.JSON(http.StatusOK, transactions)
}
