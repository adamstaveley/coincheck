package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// parse the command line for a user specified currency
	currency := flag.String("c", "USD", "Specify a currency by using its 3 character code")
	flag.Parse()
	*currency = strings.ToUpper(*currency)
	args := flag.Args()

	// format command line arguments
	var coins []string
	for _, arg := range args {
		if len(arg) == 3 {
			arg = strings.ToUpper(arg)
		} else if len(arg) > 3 {
			arg = strings.ToLower(arg)
		}
		coins = append(coins, arg)
	}

	// find the exchange rate for the given currency
	var rate float64
	if *currency != "USD" {
		ratesInterface, err := parseFixer()
		if err != nil {
			panic(err.Error())
		}
		ratesMap := ratesInterface.(map[string]interface{})
		rates := ratesMap["rates"].(map[string]interface{})
		rateInterface := rates[*currency]
		// handle non-existent currency
		var ok bool
		rate, ok = rateInterface.(float64)
		if !ok {
			fmt.Println("Currency not found:", *currency)
			os.Exit(0)
		}
	} else {
		rate = 1.0
	}

	// populate an array of relevant coins
	var coinArray []coinInfo
	allCoins, err := parseCoinMarketCap()
	if err != nil {
		panic(err.Error())
	} else {
		switch len(coins) {
		case 0:
			// use top 10 coins if no cryptocurrency specified
			coinArray = allCoins[:10]
		default:
			// find all matching requested cryptocurrencies
			for _, info := range allCoins {
				for i := 0; i < len(coins); i++ {
					if info.Symbol == coins[i] || info.ID == coins[i] {
						coinArray = append(coinArray, info)
					}
				}
			}
		}
	}

	// print in table format (max 80 char)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-5s %-7s %-17s %-11s %-13s %-12s %-s\n",
		"Rank", "Symbol", "Name", *currency+" Price", "BTC Price", "24h Change", "7d Change")
	fmt.Println(strings.Repeat("-", 80))

	for _, info := range coinArray {
		price, err := strconv.ParseFloat(info.PriceUSD, 64)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%-5s %-7s %-17s %-11.2f %-13s %-12s %-s\n",
			info.Rank, info.Symbol, info.Name, price*rate, info.PriceBTC, info.PercentChange24h+"%", info.PercentChange7d+"%")
	}
	fmt.Println(strings.Repeat("-", 80))
}

type coinInfo struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUSD         string `json:"price_usd"`
	PriceBTC         string `json:"price_btc"`
	Volume24hUSD     string `json:"24h_volume_usd"`
	MarketCapUSD     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

func requestBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func parseFixer() (interface{}, error) {
	body, err := requestBody("https://api.fixer.io/latest?base=USD")
	if err != nil {
		panic(err.Error())
	}
	var rates interface{}
	err = json.Unmarshal(body, &rates)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return rates, err
}

func parseCoinMarketCap() ([]coinInfo, error) {
	body, err := requestBody("https://api.coinmarketcap.com/v1/ticker")
	if err != nil {
		panic(err.Error())
	}
	var coins []coinInfo
	err = json.Unmarshal(body, &coins)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return coins, err
}
