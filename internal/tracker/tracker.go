package tracker

import (
	"encoding/json"
	"fmt"

	"github.com/GosMachine/WaybarCryptoTracker/internal/utils"
)

const binanceApi = "https://api.binance.com/api/v3/ticker/price?symbol=%s"

type BinanceResponse struct {
	Price string
}

func GetPrice(symbol string) (string, error) {
	respJson, err := utils.GetRequest(fmt.Sprintf(binanceApi, symbol))
	if err != nil {
		return "", err
	}
	var response BinanceResponse
	err = json.Unmarshal(respJson, &response)
	if err != nil {
		return "", err
	}
	return response.Price, nil
}
