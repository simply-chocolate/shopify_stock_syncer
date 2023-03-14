package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func NotifyTeamsSuccesCount(barcodes []string) {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")
	card := messagecard.NewMessageCard()

	if len(barcodes) == 0 {
		card.Title = "No products were updated in Shopify"
		card.Text = "There was no products which inventory could be updated.<BR/>" +
			"If this was the only message during this run, then there was no products which needed to be updated.<BR/>" +
			"If not, check the other messages for detalied error messages."
	} else {
		card.Title = "Successfully updated inventory in Shopify"
		barcodesString := ""
		for _, barcode := range barcodes {
			barcodesString += "**Barcode**: " + barcode + "<BR/>"
		}

		card.Text = fmt.Sprintf("Script has finished and these %v barcodes were updated:.<BR/>"+barcodesString, len(barcodes))
	}

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.: ", err)
	}

}
