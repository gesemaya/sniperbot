# A bot for Telegram

This is my bot for Telegram. Written in Go. I use it to test CICD pipelines and general working with Telegram API.

## How to run

You can get your token from [@BotFather](https://t.me/BotFather).

Just clone the repo and run through Makefile:

```bash
git clone https://github.com/gesemaya/sniperbot
cd bot
echo "TOKEN=your_token_without_quotes" >> .env
make run
```
