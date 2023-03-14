package teams_notifier

import (
	"encoding/json"
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendUpdateInventoryLevelErrorToTeams(barcode string, productId json.Number, quantityPCN int, err error) {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "Inventory Item Error"
	card.Text = fmt.Sprintf("Script failed to find an inventory item.<BR/>"+
		"**Product Barcode**: %v<BR/>"+
		"**Quantity in PCN**: %v<BR/>"+
		"**Link**: https://simply-chocolate-copenhagen.myshopify.com/admin/products/%v<BR/>"+
		"**Error**: %v", barcode, quantityPCN, productId, err)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.", err)
	}

}
