package phonesender

import (
	"fmt"
	"gatepass/src/codetophone"
	"gatepass/src/config"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMessage(message string, toPhone string) (bool, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.TwilioAccountSID,
		Password: config.AuthToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(config.FromPhone)
	var messageToSend string

	switch config.Language {
	case "EN":
		messageToSend = "This is your login code from " + config.AppName + ":\n" + message + "\nDo not share this code with anyone"
	case "ES":
		messageToSend = "Este es tu código de identificación de " + config.AppName + ":\n" + message + "\nNo compartas este código con nadie"

	default:
		messageToSend = "This is your login code from " + config.AppName + ":\n" + message + "\nDo not share this code with anyone"
	}
	params.SetBody(messageToSend)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Printf("Mensaje enviado con SID: %s\n", *resp.Sid)
	codetophone.AddCodeToPhone(message, toPhone)
	return true, nil
}
