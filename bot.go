package main

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var scheduler Scheduler

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go scheduler.Run()

	for update := range updates {
		if update.Message != nil {
			handleMessage(update.Message, bot)
		}
	}
}

func handleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	text := message.Text
	chatID := message.Chat.ID

	if text == "/start" {
		msg := tgbotapi.NewMessage(chatID, "Welcome to Task Scheduler Bot! Use /add <task_name> <delay_in_seconds> to add a task.")
		bot.Send(msg)
	} else if text == "/add" {
		// Simplified example: add a task that executes in 10 seconds
		scheduler.AddTask("Example Task", time.Now().Add(10*time.Second), 0, false, 1, true)
		msg := tgbotapi.NewMessage(chatID, "Task added! It will execute in 10 seconds.")
		bot.Send(msg)
	} else if len(text) > 5 && text[:4] == "/add" {
		args := text[5:]
		var name string
		var delay int
		fmt.Sscanf(args, "%s %d", &name, &delay)
		scheduler.AddTask(name, time.Now().Add(time.Duration(delay)*time.Second), 0, false, 1, true)
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Task '%s' added! It will execute in %d seconds.", name, delay))
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, "Unknown command.")
		bot.Send(msg)
	}
}
