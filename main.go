package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	var (
		url         string
		totalReqs   int
		concurrency int
	)

	flag.StringVar(&url, "url", "", "URL do servico a ser testado (obrigatório)")
	flag.IntVar(&totalReqs, "requests", 100, "Número total de requests (padrão:100)")
	flag.IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas (padrão 1)")
	flag.Parse()

	if url == "" {
		fmt.Println("É necessário informar a URL. Utilize --url=<URL>")
	}

	start := time.Now()

	requestChan := make(chan int)

	var wg sync.WaitGroup

	statusCount := make(map[int]int)
	var statusMutex sync.Mutex

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range requestChan {
				resp, err := http.Get(url)
				if err != nil {
					statusMutex.Lock()
					statusCount[-1]++
					statusMutex.Unlock()
					continue
				}
				defer resp.Body.Close()
				statusMutex.Lock()
				statusCount[resp.StatusCode]++
				statusMutex.Unlock()
			}

		}()

	}
	wg.Wait()

	for i := 0; i < totalReqs; i++ {
		requestChan <- i
	}
	close(requestChan)

	elapsed := time.Since(start)

	fmt.Println("===== RELATÓRIO DE TESTE DE CARGA =====")
	fmt.Printf("URL testada: %s\n", url)
	fmt.Printf("Tempo total gasto: %v\n", elapsed)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalReqs)

	var status200Count int
	for code, count := range statusCount {
		if code == 200 {
			status200Count += count
		}

	}
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", status200Count)

	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for code, count := range statusCount {
		if code == 200 {
			continue
		}
		fmt.Printf("  Status %d: %d\n", code, count)
	}

}
