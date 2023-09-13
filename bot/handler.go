package bot

import (
	tele "github.com/gesemaya/tele/pkg/tele"
)

// onStart is the handler for the /start command. sends the start message with the menu.
func (b *Bot) onStart(c tele.Context) error {
	return c.Send(b.Text(c, "start"), b.Markup(c, "start"))
}

// onUtils is the handler for the utils button. edits the utils message with the utils menu.
func (b *Bot) onUtils(c tele.Context) error {
	return c.Edit(b.Text(c, "utils"), b.Markup(c, "utils"))
}

// onMenu is the handler for the menu button. edits the menu message with the menu.
func (b *Bot) onMenu(c tele.Context) error {
	return c.Edit(b.Text(c, "menu"), b.Markup(c, "start"))
}
