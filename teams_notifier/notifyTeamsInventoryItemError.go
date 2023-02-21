package teams_notifier

import (
	"encoding/json"
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendInventoryItemErrorToTeams(barcode string, productId json.Number) {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "Inventory Item Error"
	card.Text = fmt.Sprintf("Script failed to find an inventory item.<BR/>"+
		"**Product Barcode**: %v<BR/>"+
		"**Link**: https://simply-chocolate-copenhagen.myshopify.com/admin/products/%v<BR/>"+
		"**Error**: Could not find inventory item from the product variant", barcode, productId)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.")
	}

}
