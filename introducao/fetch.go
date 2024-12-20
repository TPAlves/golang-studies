// Fetch exibe o conteúdo encontrado em cada URL especificada.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http")) {
			url = "http://" + url 
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "Fetch: %v\n", url, err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		codeStatus := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		
		fmt.Printf("%s", b, "\n")
		fmt.Println(codeStatus)
	}
}