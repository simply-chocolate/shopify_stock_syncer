# shopify_stock_syncer
Created by Jedikrigeren for Simply Chocolate.

The purpose of this script is to always have the latest inventory amounts updated in shopify.

It is supposed to run every 2 hours between 07:00 and 17:00 and sends error messages to a dedicated channels in teams.

Right now the known issues are:
  - Script sometimes crashes with the error : `panic: interface conversion: interface {} is nil, not *shopify_api_wrapper.ShopifyApiInventoryItemResult`. happens at `C:/Projects/shopify_stock_syncer/shopify_api_wrapper/req_inventoryitem_GET.go:35`

Known usererrors: 
  - Product is not tracked but should be tracked

Most usererrors should no longer relevant, since the creates of the sync-products-shopify-sap script, as it makes sure the products are getting created correctly in both Shopify and PCN.

