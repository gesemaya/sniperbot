package bot

import tele "github.com/gesemaya/tele/pkg/tele"

// onID is the handler for the id button. edits this message the id message with the menu.
func (b *Bot) onID(c tele.Context) error {
	var args tele.M = tele.M{
		"ID":     c.Sender().ID,
		"ChatID": c.Chat().ID,
	}
	return c.Edit(b.Text(c, "id", args), b.Markup(c, "menu"))
}
