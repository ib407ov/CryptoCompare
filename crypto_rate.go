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
	for i := 0; i < len(crypto)/25+1; i++ {
		if indexB > (len(crypto)) {
			indexB = len(crypto)
		}

		newCrypto := crypto[indexA:indexB]
		newCryptoStrig := strings.Join(newCrypto, ",")

		//Get API from URL
		url := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemulti?fsyms=%s&tsyms=USD", newCryptoStrig)
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

		//Unmarshal body
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println("Помилка при розборі JSON:", err)
			return CryptoCompareResponseStruct{}, err
		}

		for key, val := range result {
			dataStruct.Data = append(dataStruct.Data, CryptoCompare{
				Symbol: key,
				Price:  val.Price,
			})
		}

		indexA += 25
		indexB += 25
	}

	return dataStruct, nil
}
