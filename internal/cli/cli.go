package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/GosMachine/WaybarCryptoTracker/internal/tracker"
	"github.com/spf13/cobra"
)

type Output struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
}

func Run() error {
	var symbolsStr string
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Waybar crypto price tracker",
		Long:  "Cryptocurrency price parsing (only binance api for now)",
	}
	var priceCmd = &cobra.Command{
		Use:   "price",
		Short: "Prints current price",
		Run: func(cmd *cobra.Command, args []string) {
			symbols := strings.Split(symbolsStr, ",")
			var output Output
			for _, symbol := range symbols {
				symbol = strings.TrimSpace(symbol)
				symbolSplit := strings.Split(symbol, "/")
				if len(symbolSplit) != 2 {
					continue
				}

				price, err := getPrice(symbolSplit[0] + symbolSplit[1])
				if err != nil {
					continue
				}
				if len(output.Text) != 0 {
					output.Text += " | "
				}
				output.Text += fmt.Sprintf("%s %s", symbolSplit[0], price)
			}
			outputJson, err := json.Marshal(&output)
			if err != nil {
				log.Fatalln("error marshalling output")
			} else {
				fmt.Println(string(outputJson))
			}

		},
	}
	priceCmd.Flags().StringVarP(&symbolsStr, "symbols", "s", "BTC/USDT", "Symbols to track (Example: BTC/USDT, ETH/USDT)")
	rootCmd.AddCommand(priceCmd)
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func getPrice(symbol string) (string, error) {
	priceStr, err := tracker.GetPrice(symbol)
	if err != nil {
		return "", err
	}
	priceFloat, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f", priceFloat), nil
}
