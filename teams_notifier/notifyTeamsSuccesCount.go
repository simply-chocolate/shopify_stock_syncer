package teams_notifier

import (
	"fmt"
	"os"
	"strings"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

type ProductAmounts struct {
	Barcode         string
	ProductName     string
	QuantityShopify int
	QuantityPCN     int
}

func NotifyTeamsSuccesCount(products []ProductAmounts) {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")
	card := messagecard.NewMessageCard()

	if len(products) == 0 {
		card.Title = "No products were updated in Shopify"
		card.Text = "There was no products which inventory could be updated.<BR/>" +
			"If this was the only message during this run, then there was no products which needed to be updated.<BR/>" +
			"If not, check the other messages for detalied error messages."
	} else {
		card.Title = "Successfully updated inventory in Shopify"
		productsString := "Barcode - Name - Amount Shopify - Amount PCN <BR/>"
		for _, product := range products {
			productName := strings.Split(product.ProductName, "|")[0]
			if len(productName) > 15 {
				productName = productName[0:15]
			}

			productsString += fmt.Sprintf("%v - %v - **%v** - **%v**<BR/>", product.Barcode, productName, product.QuantityShopify, product.QuantityPCN)
		}
		card.Text = fmt.Sprintf("Script has finished and these %v barcodes were updated:.<BR/> "+productsString, len(products))
	}

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.: ", err)
	}

}
