package pcn_api_wrapper

import (
	"errors"
	"fmt"
	"os"
)

type PcnApiGetStockDataResult struct {
	Results []struct {
		Barcode   string `json:"barcode"`
		Available int    `json:"instock"`
	} `json:"results"`
}

type PcnApiGetStockDataReturn struct {
	Body *PcnApiGetStockDataResult
}

// Calls the PCN API at endpoint stocklist to return Barcodes and available units in stock.
func PcnApiGetStockData() (PcnApiGetStockDataReturn, error) {
	resp, err := GetPcnApiBaseClient().
		R().
		SetResult(PcnApiGetStockDataResult{}).
		SetBody(map[string]interface{}{
			"cid":        os.Getenv("PCN_CID"),
			"olsuser":    os.Getenv("PCN_OLSUSER"),
			"olspass":    os.Getenv("PCN_OLSPASS"),
			"filter":     "all",
			"maxresults": 1000,
		}).Post("stocklist")
	if err != nil {
		fmt.Println(err)
		return PcnApiGetStockDataReturn{}, errors.New("error getting the stock data fro Pcn API")
	}

	return PcnApiGetStockDataReturn{
		Body: resp.Result().(*PcnApiGetStockDataResult),
	}, nil

}
