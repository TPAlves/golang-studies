// Fetchall busca URLs em paralelo e informa os tempos gastos e os tamanhos.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main()  {
	start := time.Now()
	ch := make(chan string)
	PATH := "files/fetchAll.txt"
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // inicia uma gorrotina
	}
	for range os.Args[1:] {
		_, err := os.Stat(PATH)
		if err != nil {
			_, err := os.Create(PATH)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		}
		file, err := os.OpenFile(PATH, os.O_RDWR | os.O_APPEND, 0666)
		defer file.Close()

		_, err = file.WriteString(<-ch + "\n")

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// fmt.Println(<-ch) // Recebe o sinal de ch
	}
	
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // envia para o canal ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}