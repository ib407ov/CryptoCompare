package cryptocompare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CryptoCompare struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"USD"`
}

type CryptoCompareResponse map[string]CryptoCompare

func GetDataCurrencyRate(crypto []string) (CryptoCompareResponse, error) {
	countGoroutines := len(crypto)/25 + 1
	fmt.Println(countGoroutines)
	cryptoRow := strings.Join(crypto, ",")
	var result CryptoCompareResponse

	//Get API from URL
	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemulti?fsyms=%s&tsyms=USD", cryptoRow)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return result, err
	}

	//API in body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Помилка при читанні відповіді:", err)
		return result, err
	}

	//Unmarshal body
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Помилка при розборі JSON:", err)
		return result, err
	}

	return result, nil
}

//func main() {
//	crypto := []string{"THETA", "KCS", "XTZ", "BTT", "BGB",
//		"1000SATS", "CHZ", "MANA", "NEO", "PYTH",
//		"CFX", "EOS", "WEMIX", "BLUR", "OSMO", "ROSE",
//		"RON", "KLAY", "KAVA", "BONK", "PENDLE", "USDD",
//		"AKT", "WOO"}
//	res, _ := GetDataCurrencyRate(crypto)
//	fmt.Println(res)
//}
