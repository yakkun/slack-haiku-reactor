package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mattn/go-haiku"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"github.com/yakkun/slack-haiku-reactor/config"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Unable to load .env file. (Ignore me if you don't handle .env file)")
	}

	conf := config.New()
	if err := conf.Load(); err != nil {
		log.Fatalf("Unable to load config: %#v", err)
	}

	if conf.Debugging == true {
		log.Print("Debug mode is active.")
	}
	if conf.SlackBotToken == "" {
		log.Fatal("SlackBotToken is not set, must set it with Env-vars or .env")
	}
	if conf.SlackAppToken == "" {
		log.Fatal("SlackAppToken is not set, must set it with Env-vars or .env")
	}

	client := slack.New(
		conf.SlackBotToken,
		slack.OptionAppLevelToken(conf.SlackAppToken),
		slack.OptionDebug(conf.Debugging),
	)

	socketMode := socketmode.New(
		client,
		socketmode.OptionDebug(conf.Debugging),
	)

	authTest, err := client.AuthTest()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to authenticate: %v\n", err)
		os.Exit(1)
	}

	selfUserId := authTest.UserID

	go func() {
		for envelope := range socketMode.Events {
			if envelope.Type != socketmode.EventTypeEventsAPI {
				continue
			}

			socketMode.Ack(*envelope.Request)
			eventPayload, _ := envelope.Data.(slackevents.EventsAPIEvent)
			if eventPayload.Type != slackevents.CallbackEvent {
				continue
			}

			switch event := eventPayload.InnerEvent.Data.(type) {
			case *slackevents.MessageEvent:
				if event.User == selfUserId {
					continue
				}

				// TODO: 川柳への対応
				// TODO: 字余り/字足らずへの対応

				if !haiku.Match(event.Text, []int{5, 7, 5}) {
					continue
				}

				err := client.AddReaction(
					conf.ReactEmojiForHaiku, // TODO: ":emoji:" でもいいようにしたい (現在は "emoji")
					slack.ItemRef{
						Timestamp: event.TimeStamp,
						Channel:   event.Channel,
					},
				)
				if err != nil {
					log.Printf("Unable to react emoji: %v", err)
				}
			}
		}
	}()

	socketMode.Run()
}
