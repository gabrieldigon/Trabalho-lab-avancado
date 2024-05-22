//	Checagem de preco bitcoin e etherum de API publica de forma PARALELA
//
// Equipe Gabriel Dias & Luciano Uchoa
package main

// importacoes dos pacotes utilizados
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync" // Pra fazer de forma paralela precisamos dessa lib
)

// As unicas mudancas do codigo paralelo pro sequencial sao na funcao main
func main() {
	// Utiliza uma WaitGroup para sincronizar as goroutines
	var wg sync.WaitGroup
	var precoEth float64
	var precoBitCoin float64
	// Incrementa o contador da WaitGroup para cada goroutine lançada
	wg.Add(2)

	// Utiliza goroutines para chamar as funções de obtenção de preço de Bitcoin e Ethereum simultaneamente
	go func() {
		defer wg.Done() // Decrementa o contador da WaitGroup quando a goroutine é finalizada
		precoBitCoin = getBitcoinPrice()
		fmt.Println("Preço do Bitcoin:", precoBitCoin)
	}()

	go func() {
		defer wg.Done() // Decrementa o contador da WaitGroup quando a goroutine é finalizada
		precoEth = getCurrentEthereumPrice()
		fmt.Println("Preço do Ethereum:", precoEth)
	}()

	// Espera até que todas as goroutines tenham terminado
	wg.Wait()

	// Depois que todas as goroutines terminam, calcula quantos ETH valem um BTC
	calculaQuantosETHValeUmBTC(precoEth, precoBitCoin)
}

// funcao que pega os valores resultantes das requisicoes  e faz um calculo pra saber quantos BTC valem um ETH
func calculaQuantosETHValeUmBTC(eth float64, btc float64) {
	resultado := btc / eth
	if resultado > 0 {
		fmt.Printf("Um Bitcoin vale %.2f Ethereuns\n", resultado)
	}
}

func getBitcoinPrice() float64 {
	// Guarda da  url da API de consulta em uma variavel
	apiURL := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
	// Prepara dois Resultados  um erro ou uma resposta dependendo da leitura feita com o Get, se for erro printa erro de respsota
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Print("erro de resposta na API")
		return 0
	}
	// caso a resposta seja 200 que e positiva pega  o body dessa resposta
	defer response.Body.Close()
	// Prepara dois Resultados  um erro ou um body  se for erro printa erro no body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("erro no body do Json")
		return 0
	}
	// Faz o map do resultado do body, e como se fosse um decode e se for erro printa erro no map
	var data map[string]map[string]float64
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Print("erro no map")
		return 0
	}
	// retorna o preco do que foi mapeado com os dois parametros que e a resposta da api
	priceUSD := data["bitcoin"]["usd"]
	return priceUSD
}

// Essa func faz a mesma coisa da de cima, o que muda e a url que no caso se refere ao ethereum
func getCurrentEthereumPrice() float64 {
	// Guarda da  url da API de consulta em uma variavel
	apiURL := "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd"
	// Prepara dois Resultados  um erro ou uma resposta dependendo da leitura feita com o Get, se for erro printa erro de respsota
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Print("erro na API")
		return 0
	}
	// caso a resposta seja 200 que e positiva pega  o body dessa resposta
	defer response.Body.Close()
	// Prepara dois Resultados  um erro ou um body  se for erro printa erro no body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("erro no body")
		return 0
	}
	// Faz o map do resultado do body, e como se fosse um decode e se for erro printa erro no map
	var data map[string]map[string]float64
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Print("erro no map")
		return 0
	}
	// retorna o preco do que foi mapeado com os dois parametros que e a resposta da api
	priceUSD := data["ethereum"]["usd"]

	return priceUSD
}

// Obs : Fizemos esse codigo pq gostariamos de testar a forma de fazer requisicoes e GO, a forma paralela e sequencial nao tem tantas mudancas, com uma internet rapida a diferenca deles e imperceptivel
// mas com uma conexao 2g ou 3g, seria possivel notar essa diferença, uma vez que as duas requisicoes ocorreriam de forma paralela. GO foi uma linguagem bem tranquila pra fazer requisicoes
// OBS 2 : A api publica utilizada pra pegar o valor das moedas foi a coingecko presente neste link-> https://www.coingecko.com/pt/api
// OBS 3 IMPORTANTE : Nao consegui anexar dois arquivos no colab entao colei o segundo programa aq em baixo, pra usar um basta comentar o outro

// --------------------------------------------------------------------------------outro programa----------------------------------------------------------------------------------------------------------

//                                                        Checagem de preco bitcoin e etherum de API publica de forma sequencial
// Equipe Gabriel Dias & Luciano Uchoa

/// comece a descomentar a partir da linha abaixo--
// package main

// // importacoes dos pacotes utilizados
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// // funcao main que serve apenas pra chamar as outras funcs
// func main() {
// 	precoBitCoin := getBitcoinPrice()
// 	precoEth := getCurrentEthereumPrice()
// 	calculaQuantosETHValeUmBTC(precoEth, precoBitCoin)

// }

// // funcao que pega os valores resultantes das requisicoes  e faz um calculo pra saber quantos BTC valem um ETH
// func calculaQuantosETHValeUmBTC(eth float64, btc float64) {
// 	resultado := btc / eth
// 	if resultado > 0 {
// 		fmt.Printf("um bitcoin vale %.2f etheruns", resultado)
// 	}

// }

// func getBitcoinPrice() float64 {
// 	// Guarda da  url da API de consulta em uma variavel
// 	apiURL := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
// 	// Prepara dois Resultados  um erro ou uma resposta dependendo da leitura feita com o Get, se for erro printa erro de respsota
// 	response, err := http.Get(apiURL)
// 	if err != nil {
// 		fmt.Print("erro de resposta na API")
// 		return 0
// 	}
// 	// caso a resposta seja 200 que e positiva pega  o body dessa resposta
// 	defer response.Body.Close()
// 	// Prepara dois Resultados  um erro ou um body  se for erro printa erro no body
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Print("erro no body do Json")
// 		return 0
// 	}
// 	// Faz o map do resultado do body, e como se fosse um decode e se for erro printa erro no map
// 	var data map[string]map[string]float64
// 	if err := json.Unmarshal(body, &data); err != nil {
// 		fmt.Print("erro no map")
// 		return 0
// 	}
// 	// retorna o preco do que foi mapeado com os dois parametros que e a resposta da api
// 	priceUSD := data["bitcoin"]["usd"]
// 	return priceUSD
// }

// // Essa func faz a mesma coisa da de cima, o que muda e a url que no caso se refere ao ethereum
// func getCurrentEthereumPrice() float64 {
// 	// Guarda da  url da API de consulta em uma variavel
// 	apiURL := "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd"
// 	// Prepara dois Resultados  um erro ou uma resposta dependendo da leitura feita com o Get, se for erro printa erro de respsota
// 	response, err := http.Get(apiURL)
// 	if err != nil {
// 		fmt.Print("erro na API")
// 		return 0
// 	}
// 	// caso a resposta seja 200 que e positiva pega  o body dessa resposta
// 	defer response.Body.Close()
// 	// Prepara dois Resultados  um erro ou um body  se for erro printa erro no body
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Print("erro no body")
// 		return 0
// 	}
// 	// Faz o map do resultado do body, e como se fosse um decode e se for erro printa erro no map
// 	var data map[string]map[string]float64
// 	if err := json.Unmarshal(body, &data); err != nil {
// 		fmt.Print("erro no map")
// 		return 0
// 	}
// 	// retorna o preco do que foi mapeado com os dois parametros que e a resposta da api
// 	priceUSD := data["ethereum"]["usd"]

// 	return priceUSD
// }
