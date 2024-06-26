package pcn_api_wrapper

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type PcnApiGetStockDataResult struct {
	Results []struct {
		Barcode   string `json:"barcode"`
		Available int    `json:"instock"`
		OnOrder   int    `json:"onorder"`
	} `json:"results"`
}

type PcnApiGetStockDataReturn struct {
	Body *PcnApiGetStockDataResult
}

// Calls the PCN API at endpoint stocklist to return Barcodes and available units in stock.
func PcnApiGetStockData() (PcnApiGetStockDataReturn, error) {
	resp, err := GetPcnApiBaseClient().
		R().
		SetSuccessResult(PcnApiGetStockDataResult{}).
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
	if resp.StatusCode == 403 {
		// wait 5 minutes and try again
		fmt.Println(resp)
		fmt.Printf("%v: 403 error waiting 5 minutes\n", time.Now())
		time.Sleep(5 * time.Minute)
		return PcnApiGetStockData()

	} else if resp.StatusCode != 200 {
		return PcnApiGetStockDataReturn{}, errors.New("error getting the stock data fro Pcn API")
	}

	return PcnApiGetStockDataReturn{
		Body: resp.SuccessResult().(*PcnApiGetStockDataResult),
	}, nil

}
