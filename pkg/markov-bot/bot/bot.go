package bot

import (
	"log"
	"strings"

	"github.com/cjonesy/markov-bot/pkg/markov-bot/markov"
	"github.com/slack-go/slack"
)

// Bot represents a markov chat bot
type Bot struct {
	slackClient *slack.Client
	slackRTM    *slack.RTM
	userID      string
	userName    string
	maxResponse int
	chain       *markov.Chain
}

// New creates a new bot
func New(token string, chain *markov.Chain, maxResponse int) (*Bot, error) {
	client := slack.New(token)
	rtm := client.NewRTM()
	go rtm.ManageConnection()

	bot := &Bot{
		slackClient: client,
		slackRTM:    rtm,
		chain:       chain,
		maxResponse: maxResponse,
	}

	return bot, nil
}

// Start the bot
func (bot *Bot) Start() {
	for msg := range bot.slackRTM.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			bot.handleConnectedEvent(ev)
		case *slack.MessageEvent:
			bot.handleMessageEvent(ev)
		case *slack.RTMError:
			log.Print(ev.Error())
		case *slack.InvalidAuthEvent:
			log.Panic("Invalid credentials")
		default:
			// Ignore all other events..
		}
	}
}

func (bot *Bot) handleConnectedEvent(ev *slack.ConnectedEvent) {
	if bot.userID != "" {
		log.Print("Received unexpected Connected event")
		return
	}
	log.Printf(
		"Connected as user %s (%s)",
		ev.Info.User.Name,
		ev.Info.User.ID,
	)
	bot.userID = ev.Info.User.ID
	bot.userName = ev.Info.User.Name
}

func (bot *Bot) handleMessageEvent(ev *slack.MessageEvent) {
	if bot.userID == "" {
		log.Print("Received message event before finishing initialization")
		return
	}

	// Only respond if our bot was @'d
	if strings.Contains(ev.Text, bot.userID) {
		resp := bot.chain.Generate(bot.maxResponse)
		bot.slackRTM.SendMessage(bot.slackRTM.NewOutgoingMessage(resp, ev.Channel))
	}
}
