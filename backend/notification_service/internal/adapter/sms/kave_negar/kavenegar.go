package kavenegar

import (
	"fmt"

	"github.com/kavenegar/kavenegar-go"
)

func kavenegarConfig() {

}

func Send(message string, receptor []string) error {

	api := kavenegar.New(" your apikey ")
	sender := ""

	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			return err
		case *kavenegar.HTTPError:
			return err
		default:
			return err
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
			//...
		}
	}

}
