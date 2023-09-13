package bot

import (
	tele "github.com/gesemaya/tele/pkg/tele"
	"github.com/google/uuid"
)

// onHelp is the handler for the /help command. sends the help message with the menu.
func (b Bot) onHelp(c tele.Context) error {
	results := tele.Results{}

	cmds := b.Layout.Commands()

	for _, cmd := range cmds {
		result := &tele.ArticleResult{
			Title:       cmd.Text,
			Description: cmd.Description,
			Text:        cmd.Description,
		}
		result.SetResultID(uuid.New().String())
		results = append(results, result)
	}

	return c.Answer(&tele.QueryResponse{
		Results:   results,
		CacheTime: -1,
	})
}
