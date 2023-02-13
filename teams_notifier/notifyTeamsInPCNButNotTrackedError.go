package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendInPCNButNotTrackedErrorToTeams(barcode string, productId int) {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "Item in PCN but not tracked error"
	card.Text = fmt.Sprintf("Script ran into an issue with a product.<BR/>"+
		"**Product Barcode**: %v<BR/>"+
		"**Link**: https://simply-chocolate-copenhagen.myshopify.com/admin/products/%v<BR/>"+
		"**Error**: Product exist in PCN system, but is not tracked in Shopify", barcode, productId)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.")
	}

}
