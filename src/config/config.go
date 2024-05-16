package config

import (
	"os"
)

var TwilioAccountSID string
var AuthToken string
var FromPhone string

func InitVariables() {
	TwilioAccountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	AuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	FromPhone = os.Getenv("TWILIO_FROM_PHONE")
}
