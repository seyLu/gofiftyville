package model

type PhoneCall struct {
	ID       int
	Caller   string
	Receiver string
	Year     int
	Month    int
	Day      int
	Duration int
}
