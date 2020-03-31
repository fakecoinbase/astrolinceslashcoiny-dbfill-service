package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"sync"
)

func dbFill(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	fmt.Println("goroutine exit")
}

func main() {
	urls := []string{
		"https://blockstream.info/api/fee-estimates",
		"https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&&vs_currencies=usd",
		"https://blockstream.info/api/blocks/tip/height",
		"https://www.bitgo.com/api/v1/tx/fee",
		"https://api.smartbit.com.au/v1/blockchain/chart/block-size-total?from=2020-01-01",
		"https://ripio.com/api/v3/rates/?country=AR",
		"https://be.buenbit.com/api/market/tickers/",
		"https://api.bitso.com/v3/ticker",
		"https://argenbtc.com/public/cotizacion_js.php",
		"https://api.satoshitango.com/v3/ticker/ARS",
		"https://api.cryptomkt.com/v1/ticker",
		"https://bitex.la/api/tickers/btc_ars",
		"https://www.buda.com/api/v2/markets/btc-ars/order_book",
		"https://www.buda.com/api/v2/markets/eth-ars/order_book",
		"https://www.buda.com/api/v2/markets/ltc-ars/order_book",
		"https://www.qubit.com.ar/c_unvalue",
		"https://www.qubit.com.ar/c_value",
		"https://api.universalcoins.net/Criptomonedas/obtenerDatosHome",
		"https://api.pro.coinbase.com/products/btc-usdc/ticker",
		"https://api.pro.coinbase.com/products/dai-usdc/ticker",
		"https://api.pro.coinbase.com/products/eth-dai/ticker",
		"https://api.pro.coinbase.com/products/eth-btc/ticker",
		"https://api.pro.coinbase.com/products/ltc-btc/ticker",
		"https://api.pro.coinbase.com/products/xrp-btc/ticker",
		"https://api.pro.coinbase.com/products/xlm-btc/ticker"}

	urlsLength := len(urls)

	var wg sync.WaitGroup

	wg.Add(urlsLength)

	for i := 0; i < urlsLength; i++ {
		go dbFill(urls[i], &wg)
	}

	wg.Wait()

	fmt.Println("main goroutine exit")
}
