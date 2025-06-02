package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
}

func worker(url string, requests int, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()
	client := &http.Client{}
	for i := 0; i < requests; i++ {
		resp, err := client.Get(url)
		if err != nil {
			results <- Result{StatusCode: 0}
			continue
		}
		results <- Result{StatusCode: resp.StatusCode}
		resp.Body.Close()
	}
}

func main() {
	// CLI flags
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 1, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" || *totalRequests < 1 || *concurrency < 1 {
		fmt.Println("Uso: --url=<url> --requests=<n> --concurrency=<n>")
		os.Exit(1)
	}

	start := time.Now()

	results := make(chan Result, *totalRequests)
	var wg sync.WaitGroup

	requestsPerWorker := *totalRequests / *concurrency
	extra := *totalRequests % *concurrency

	// Start workers
	for i := 0; i < *concurrency; i++ {
		reqs := requestsPerWorker
		if i < extra {
			reqs++
		}
		wg.Add(1)
		go worker(*url, reqs, &wg, results)
	}

	wg.Wait()
	close(results)
	elapsed := time.Since(start)

	// Report
	statusCount := make(map[int]int)
	total := 0
	okCount := 0

	for res := range results {
		statusCount[res.StatusCode]++
		total++
		if res.StatusCode == 200 {
			okCount++
		}
	}

	fmt.Println("===== Relatório de Teste de Carga =====")
	fmt.Printf("Tempo total gasto: %v\n", elapsed)
	fmt.Printf("Total de requests realizados: %d\n", total)
	fmt.Printf("Requests com status 200: %d\n", okCount)
	fmt.Println("Distribuição dos códigos de status HTTP:")
	for code, count := range statusCount {
		if code == 0 {
			fmt.Printf("Falhas de requisição: %d\n", count)
		} else {
			fmt.Printf("Status %d: %d\n", code, count)
		}
	}
}
