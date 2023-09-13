package bot

import (
	"bytes"
	"encoding/json"
	"net/http"

	tele "github.com/gesemaya/tele/pkg/tele"
	"github.com/google/uuid"
)

func (b *Bot) onShortenerInline(c tele.Context) error {
	results := tele.Results{}

	result := &tele.ArticleResult{
		Title: "Create Short URL",
		Text:  "Create Short URL",
	}
	result.ReplyMarkup = (b.Markup(c, "wait"))
	result.SetResultID(uuid.New().String())

	results = append(results, result)

	return c.Answer(&tele.QueryResponse{
		Results:   results,
		CacheTime: -1,
	})
}

func (b *Bot) onPostShortenerInline(c tele.Context) error {
	_, data := b.parseQuery(c.Update().InlineResult.Query)
	if data == "" {
		return nil
	}

	type URL struct {
		URL string `json:"url"`
	}

	url, err := json.Marshal(URL{data})
	if err != nil {
		return err
	}

	resp, err := http.Post("https://s.x16.me/short", "application/json", bytes.NewReader(url))
	if err != nil {
		return err
	}

	var short struct {
		Short string `json:"short"`
	}

	err = json.NewDecoder(resp.Body).Decode(&short)
	if err != nil {
		return err
	}

	_, err = b.Edit(
		c.InlineResult(),
		b.Text(c, "link", tele.M{"Link": "https://s.x16.me/" + short.Short}),
		b.Markup(c, "link", "https://s.x16.me/"+short.Short))

	if err != nil {
		return err
	}

	return nil
}
