package main

import (
    "log"
    "strings"

    "github.com/sbstjn/hanu"
)

func main() {
    slack, err := hanu.New("TOKEN")

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
	
    slack.Command("joke", func(conv hanu.ConversationInterface) {
        conv.Reply("So this programmer goes out on a date with a hot girl")
    })
	
    slack.Command("joke", func(conv hanu.ConversationInterface) {
        conv.Reply("Why did the Integer drown? ‘Coz he couldn’t Float!")
    })

    slack.Command("L<word>", func(conv hanu.ConversationInterface) {
        str, _ := conv.String("word")
        
        switch str{
            case "05":
                conv.Reply("https://goo.gl/LeFwf3")
            case "06":
                conv.Reply("https://goo.gl/xavjQO")
            case "07":
                conv.Reply("https://goo.gl/pyF4AN")
            case "08":
                conv.Reply("https://goo.gl/U3TYLv")
            default:
                panic("pls stop trying")
        }
    })

    slack.Listen()
}
