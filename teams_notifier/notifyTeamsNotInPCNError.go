package teams_notifier

import (
	"encoding/json"
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendNotInPCNErrorToTeams(barcode string, productId json.Number) {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "Item not foind in PCN Error"
	card.Text = fmt.Sprintf("Script failed to find an item in PCN.<BR/>"+
		"**Product Barcode**: %v<BR/>"+
		"**Link**: https://simply-chocolate-copenhagen.myshopify.com/admin/products/%v<BR/>"+
		"**Error**: Product was not found in PCN system, but is tracked in Shopify", barcode, productId)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.", err)
	}

}
