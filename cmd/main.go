package main

import (
	"fmt"
	"github.com/gesemaya/sniperbot/bot"
	"github.com/gesemaya/sniperbot/client"
	"time"
)

func proxy() *client.HttpClient {

	// httpclient
	clientOpts := []client.HttpClientOption{
		client.WithTimeout(10 * time.Second),
	}

	clientOpts = append(clientOpts, client.WithProxyURL(fmt.Sprintf("socks5://%s", "127.0.0.1:15235")))

	clientOpts = append(clientOpts, client.WithUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.207.132.170 Safari/537.36 Edg/108.0.1462.54 Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"))

	httpClient := client.NewHttpClient(clientOpts...)
	return httpClient

}
func main() {
	bot, err := bot.New("bot.yml")
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}

	bot.Run()
}
