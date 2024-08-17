package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null/v5"
	"github.com/seyLu/gofiftyville/internal/model"
)

type PersonBankAccount struct {
	Name           string      `json:"name"`
	PhoneNumber    null.String `json:"phoneNumber"`
	PassportNumber null.Int    `json:"passportNumber"`
	LicensePlate   null.String `json:"licensePlate"`
	AccountNumber  int         `json:"accountNumber"`
	CreationYear   int         `json:"creationYear"`
}

func GetPeople(c *gin.Context) {
	request := c.Request.URL.Query()

	f := model.PeopleFilter{
		LicensePlate:  strings.TrimSpace(request.Get("license-plate")),
		AccountNumber: -1,
	}

	accountNumber := strings.TrimSpace(request.Get("account-number"))
	if accountNumber != "" {
		aN, err := strconv.Atoi(accountNumber)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		f.AccountNumber = aN
	}

	people, err := model.People(f)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting People (licensePlate %s, accountNumber %s): %v", f.LicensePlate, accountNumber, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var accounts []PersonBankAccount
	for _, account := range people {
		accounts = append(accounts, PersonBankAccount{
			Name:           account.Name,
			PhoneNumber:    account.PhoneNumber,
			PassportNumber: account.PassportNumber,
			LicensePlate:   account.LicensePlate,
			AccountNumber:  account.AccountNumber,
			CreationYear:   account.CreationYear,
		})
	}

	c.JSON(http.StatusOK, accounts)
}
