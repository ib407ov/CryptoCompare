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

type CryptoCompareResponseStruct struct {
	Data []CryptoCompare
}

func GetDataCurrencyRate(crypto []string) (CryptoCompareResponseStruct, error) {
	var result CryptoCompareResponse
	var dataStruct CryptoCompareResponseStruct

	indexA := 0
	indexB := 25
	countFors := len(crypto) / 25
	if len(crypto)%25 != 0 {
		countFors += 1
	}

	for i := 0; i < countFors; i++ {
		if indexB > len(crypto) {
			indexB = len(crypto) - 1
		}

		newCrypto := crypto[indexA:indexB]
		newCryptoString := strings.Join(newCrypto, ",")

		//Get API from URL
		key := "7622387be1531249e85802476ed5ebe7af9b18a1a2446da5d606c3c06d8a9fb9"
		url := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemulti?fsyms=%s&tsyms=USD&api_key=%s", newCryptoString, key)
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			return CryptoCompareResponseStruct{}, err
		}

		//API -> body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Помилка при читанні відповіді:", err)
			return CryptoCompareResponseStruct{}, err
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println("Помилка при розборі JSON:", err)
			indexA += 25
			indexB += 25
			continue
		}
		indexA += 25
		indexB += 25

	}

	for key, val := range result {
		dataStruct.Data = append(dataStruct.Data, CryptoCompare{
			Symbol: key,
			Price:  val.Price,
		})
	}

	return dataStruct, nil
}
