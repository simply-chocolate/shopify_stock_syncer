package shopify_api_wrapper

import (
	"fmt"
	"time"
)

type SetInventoryLevelBody struct {
	Location_id       int64
	Inventory_item_id int
	Available         int
}

func SetInventoryLevel(body *SetInventoryLevelBody) error {
	resp, err := GetShopifyApiBaseClient().
		R().
		SetBody(map[string]interface{}{
			"location_id":       body.Location_id,
			"inventory_item_id": body.Inventory_item_id,
			"available":         body.Available,
		}).
		Post("inventory_levels/set.json")
	if err != nil {
		return err
	}

	// Sleep 10 seconds so we don't call the api more than 40 times a minute.
	time.Sleep(10 * time.Second)

	if resp.StatusCode == 429 {
		time.Sleep(1 * time.Minute)
	}

	if resp.StatusCode == 422 {
		return fmt.Errorf("%s", resp.SuccessResult())
	}

	return nil
}
