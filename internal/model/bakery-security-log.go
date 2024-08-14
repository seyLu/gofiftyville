package model

type BakerySecurityLog struct {
	ID           int
	Year         int
	Month        int
	Day          int
	Hour         int
	Minute       int
	Activity     string
	LicensePlate string
}
