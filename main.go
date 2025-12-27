package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"io"
	"golang.org/x/net/proxy"
)

func main() {
	links := "targets.yaml"
	torAdres := "127.0.0.1:9150" 
	dialer, err := proxy.SOCKS5("tcp", torAdres, nil, proxy.Direct)
	if err != nil {
		fmt.Printf("proxy hatası: %v\n", err)
		return
	}
	
	tr := &http.Transport{Dial: dialer.Dial}
	client := &http.Client{Transport: tr}

	file, err := os.Open(links)
	if err != nil {
		fmt.Printf("dosya okunamıyor (%s): %v\n", links, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())

		if url == "" || strings.HasPrefix(url, "#") {
			continue
		}
		fmt.Printf("tarama yapılıyor %s", url)
		
		resp, err := client.Get(url)

		if err != nil {
			fmt.Println("hata handling") 
			continue 
		}

		fmt.Printf("çalışıyor (Kod: %d)\n", resp.StatusCode)

		dosyaIsmi := "sonuc_" + strings.NewReplacer("http://", "", "https://", "", "/", "_", ":", "").Replace(url) + ".html"
		out, _ := os.Create(dosyaIsmi)
		io.Copy(out, resp.Body)
		out.Close()

		resp.Body.Close()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("okunamıyor %v\n", err)
	}
}