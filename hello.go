package main

import (
    "log"
    "strings"

    "github.com/sbstjn/hanu"
)

func main() {
    slack, err := hanu.New("xoxb-156379277474-qAuZAZo0K4Qk6ZScgysHLevW")

    if err != nil {
        log.Fatal(err)
    }

    Version := "0.0.1"

    slack.Command("shout <word>", func(conv hanu.ConversationInterface) {
        str, _ := conv.String("word")
        conv.Reply(strings.ToUpper(str))
    })

    slack.Command("whisper <word>", func(conv hanu.ConversationInterface) {
        str, _ := conv.String("word")
        conv.Reply(strings.ToLower(str))
    })

    slack.Command("version", func(conv hanu.ConversationInterface) {
        conv.Reply("Thanks for asking! I'm running `%s`", Version)
    })

    slack.Command("CS324", func(conv hanu.ConversationInterface) {
        conv.Reply("Skripting jezici")
    })

    slack.Command("hi", func(conv hanu.ConversationInterface) {
		conv.Reply("Hi yourself!")
    })

    slack.Command("woopwoop", func(conv hanu.ConversationInterface) {
        conv.Reply("http://i.imgur.com/pz5gnhq.gif")
    })

    slack.Listen()
}