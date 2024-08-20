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
		LicensePlates:  nil,
		AccountNumbers: nil,
		PhoneNumbers:   nil,
	}

	licensePlates := request["license-plate"]
	for i, licensePlate := range licensePlates {
		licensePlates[i] = strings.TrimSpace(licensePlate)
	}
	if len(licensePlates) > 0 {
		f.LicensePlates = licensePlates
	}

	accountNumbersReq := request["account-number"]
	var accountNumbers []int
	for _, accountNumber := range accountNumbersReq {
		aN, err := strconv.Atoi(strings.TrimSpace(accountNumber))
		if err != nil {
			errMsg := fmt.Sprintf("(1) controller.GetPeople (licensePlates %v, accountNumbers %v, phoneNumbers %v): %v", licensePlates, accountNumbersReq, "", err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}
		accountNumbers = append(accountNumbers, aN)
	}
	if len(accountNumbers) > 0 {
		f.AccountNumbers = accountNumbers
	}

	phoneNumbers := request["phone-number"]
	for i, phoneNumber := range phoneNumbers {
		phoneNumbers[i] = strings.TrimSpace(phoneNumber)
	}
	if len(phoneNumbers) > 0 {
		f.PhoneNumbers = phoneNumbers
	}

	people, err := model.People(f)
	if err != nil {
		errMsg := fmt.Sprintf("(2) controller.GetPeople (licensePlates %v, accountNumbers %v, phoneNumbers %v): %v", f.LicensePlates, f.AccountNumbers, f.PhoneNumbers, err)
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
