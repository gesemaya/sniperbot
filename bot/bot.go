package bot

import (
	"fmt"
	"github.com/gesemaya/sniperbot/client"

	"strings"
	"time"

	tele "github.com/gesemaya/tele/pkg/tele"
	"github.com/gesemaya/tele/pkg/tele/layout"
)

type Bot struct {
	*tele.Bot

	*layout.Layout
}

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
func New(path string) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}
	setting := lt.Settings()
	setting.Client = proxy().Client()

	b, err := tele.NewBot(setting)
	if err != nil {
		return nil, err
	}

	return &Bot{b, lt}, nil
}

func (b *Bot) Run() {
	fmt.Println(111)
	b.Use(middleware.Logger())
	b.Use(b.Middleware("en"))

	b.registerHandlers()

	b.Bot.Start()
}

// onInline is the handler for inline queries. it is a kind of router for other inline handlers.
func (b *Bot) onInline(c tele.Context) error {
	fmt.Println("======")
	fmt.Println(c.Data())
	command, _ := b.parseQuery(c.Data())

	switch command {
	case "s", "short", "shrt":
		return b.onShortenerInline(c)
	}

	return b.onHelp(c)
}

// parseQuery parses the query text and returns the command and the data.
func (b Bot) parseQuery(text string) (command, data string) {
	parts := strings.Split(text, " ")
	if len(parts) >= 1 {
		command = strings.ToLower(parts[0])
		data = strings.Join(parts[1:], " ")
	}
	return
}

// registerHandlers registers all the handlers.
func (b *Bot) registerHandlers() {
	b.Handle("/start", b.onStart)
	b.Handle("/short", b.onShortenerInline)

	b.Handle(b.Callback("menu"), b.onMenu)
	b.Handle(b.Callback("id"), b.onID)
	b.Handle(b.Callback("utils"), b.onUtils)
	fmt.Println(222)
	b.Handle(tele.OnQuery, b.onInline)
	b.Handle(tele.OnInlineResult, b.onPostShortenerInline)
	fmt.Println(333)
}
